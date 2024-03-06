// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {DispatcherUpdateClientTestSuite} from "../Dispatcher.client.t.sol";
import {DispatcherIbcWithRealProofsSuite} from "../Dispatcher.proof.t.sol";
import {TestUtils} from "../TestUtils.sol";
import {Mars} from "../../contracts/examples/Mars.sol";
import "../../contracts/core/OpLightClient.sol";
import {ChannelHandshakeTestSuite} from "../Dispatcher.t.sol";
import {LocalEnd} from "../Dispatcher.base.t.sol";
import {Base, ChannelHandshakeSetting} from "../Dispatcher.base.t.sol";
import {
    CounterParty, ChannelOrder, IbcEndpoint, IbcPacket, AckPacket, Ibc, Height
} from "../../contracts/libs/Ibc.sol";
import {IbcReceiver} from "../../contracts/interfaces/IbcReceiver.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
import {OptimisticLightClient} from "../../contracts/core/OpLightClient.sol";
import {L1Block} from "optimism/L2/L1Block.sol";
import {ProofVerifier} from "../../contracts/core/OpProofVerifier.sol";
import {DummyLightClient} from "../../contracts/utils/DummyLightClient.sol";

import {IDispatcher} from "../../contracts/interfaces/IDispatcher.sol";
import {DispatcherV2Initializable} from "./upgrades/DispatcherV2Initializable.sol";

library UpgradeTestUtils {
    function deployAndUpgradeDispatcher(
        LightClient consensusStateManager,
        ProofVerifier oldOpProofVerifier,
        L1Block l1BlockProvider,
        string memory portPrefix
    ) public returns (IDispatcher dispatcher, DispatcherV2Initializable impl) {
        // // Deploy old version
        string memory oldPortPrefix = "Oldpolyibc.eth.";
        OptimisticLightClient oldConsensusManager = new OptimisticLightClient(1, oldOpProofVerifier, l1BlockProvider);
        (dispatcher,) = TestUtils.deployDispatcherProxyAndImpl(oldPortPrefix, oldConsensusManager);

        // Upgrade dispatcher for tests
        impl = new DispatcherV2Initializable();
        // string memory portPrefix = "polyibc.eth1.";
        bytes memory initData = abi.encodeWithSignature("initialize(string,address)", portPrefix, consensusStateManager);
        UUPSUpgradeable(address(dispatcher)).upgradeToAndCall(address(impl), initData);
    }
}

contract DispatcherUpgradeTest is DispatcherUpdateClientTestSuite, DispatcherIbcWithRealProofsSuite {
    function setUp() public virtual override {
        consensusStateManager = new OptimisticLightClient(1, opProofVerifier, l1BlockProvider);
        (dispatcher,) = UpgradeTestUtils.deployAndUpgradeDispatcher(
            consensusStateManager, opProofVerifier, l1BlockProvider, "polyibc.eth1."
        );
        mars = new Mars(dispatcher);
    }
}

contract DispatcherUpgradeChannelHandshakeTest is ChannelHandshakeTestSuite {
    function setUp() public override {
        (dispatcher,) = UpgradeTestUtils.deployAndUpgradeDispatcher(
            dummyConsStateManager, opProofVerifier, l1BlockProvider, "polyibc.eth1."
        );
        mars = new Mars(dispatcher);
        _local = LocalEnd(mars, portId, "channel-1", connectionHops, "1.0", "1.0");
        _remote = CounterParty("eth2.7E5F4552091A69125d5DfCb7b8C2659029395Bdf", "channel-2", "1.0");
    }
}

// This Base contract provides an open channel for sub-contract tests
contract ChannelOpenTestBaseSetup is Base {
    string portId = "eth1.7E5F4552091A69125d5DfCb7b8C2659029395Bdf";
    bytes32 channelId = "channel-1";
    address relayer = deriveAddress("relayer");
    bool feeEnabled = false;

    LocalEnd _local;
    CounterParty _remote;
    Mars mars;

    function setUp() public virtual override {
        (dispatcher,) = UpgradeTestUtils.deployAndUpgradeDispatcher(
            dummyConsStateManager, opProofVerifier, l1BlockProvider, portPrefix
        );
        ChannelHandshakeSetting memory setting =
            ChannelHandshakeSetting(ChannelOrder.ORDERED, feeEnabled, true, validProof);

        // anyone can run Relayers
        vm.startPrank(relayer);
        vm.deal(relayer, 100_000 ether);
        mars = new Mars(dispatcher);

        _local = LocalEnd(mars, portId, channelId, connectionHops, "1.0", "1.0");
        _remote = CounterParty("eth2.7E5F4552091A69125d5DfCb7b8C2659029395Bdf", "channel-2", "1.0");

        channelOpenInit(_local, _remote, setting, true);
        channelOpenTry(_local, _remote, setting, true);
        channelOpenAck(_local, _remote, setting, true);
        channelOpenConfirm(_local, _remote, setting, true);
    }
}

contract DispatcherSendPacketTestSuite is ChannelOpenTestBaseSetup {
    // default params
    string payload = "msgPayload";
    uint64 timeoutTimestamp = 1000;

    function test_success() public {
        bytes memory packet = abi.encodePacked(payload);
        for (uint64 index = 0; index < 3; index++) {
            vm.expectEmit(true, true, true, true);
            uint64 packetSeq = index + 1;
            emit SendPacket(address(mars), channelId, packet, packetSeq, timeoutTimestamp);
            mars.greet(payload, channelId, timeoutTimestamp);
        }
    }

    // sendPacket fails if calling dApp doesn't own the channel
    function test_mustOwner() public {
        Mars earth = new Mars(dispatcher);
        vm.expectRevert(abi.encodeWithSignature("channelNotOwnedBySender()"));
        earth.greet(payload, channelId, timeoutTimestamp);
    }
}

contract PacketSenderTestBase is ChannelOpenTestBaseSetup {
    IbcEndpoint dest = IbcEndpoint("polyibc.bsc.9876543210", "channel-99");
    IbcEndpoint src;
    string payloadStr = "msgPayload";
    bytes payload = bytes(payloadStr);
    bytes appAck = abi.encodePacked('{ "account": "account", "reply": "got the message" }');

    uint64 nextSendSeq = 1;
    // cached packet that was sent in `sendPacket`
    IbcPacket sentPacket;
    // ackPacket is the acknowledgement packet that is expected to be written for the `sentPacket`
    bytes ackPacket;

    function setUp() public virtual override {
        super.setUp();
        string memory marsPort = string(abi.encodePacked(portPrefix, getHexBytes(address(mars))));
        src = IbcEndpoint(marsPort, channelId);
    }

    // sendPacket writes a packet commitment, and updates cached `sentPacket` and `ackPacket`
    function sendPacket() internal {
        sentPacket = genPacket(nextSendSeq);
        ackPacket = genAckPacket(Ibc.toStr(nextSendSeq));
        mars.greet(payloadStr, channelId, maxTimeout);
        nextSendSeq += 1;
    }

    // genPacket generates a packet for the given packet sequence
    function genPacket(uint64 packetSeq) internal view returns (IbcPacket memory) {
        return IbcPacket(src, dest, packetSeq, payload, ZERO_HEIGHT, maxTimeout);
    }

    // genAckPacket generates an ack packet for the given packet sequence
    function genAckPacket(string memory packetSeq) internal pure returns (bytes memory) {
        return ackToBytes(AckPacket(true, bytes(packetSeq)));
    }
}

// Test Chains B receives a packet from Chain A
contract DispatcherRecvPacketTestSuite is ChannelOpenTestBaseSetup {
    IbcEndpoint src = IbcEndpoint("polyibc.bsc.9876543210", "channel-99");
    IbcEndpoint dest;
    bytes payload = bytes("msgPayload");
    bytes appAck = abi.encodePacked('{ "account": "account", "reply": "got the message" }');

    function setUp() public override {
        super.setUp();
        string memory marsPort = string(abi.encodePacked(portPrefix, getHexBytes(address(mars))));
        dest = IbcEndpoint(marsPort, channelId);
    }

    function test_success() public {
        for (uint64 index = 0; index < 3; index++) {
            uint64 packetSeq = index + 1;
            vm.expectEmit(true, true, true, true, address(dispatcher));
            emit RecvPacket(address(mars), channelId, packetSeq);
            vm.expectEmit(true, true, false, true, address(dispatcher));
            emit WriteAckPacket(address(mars), channelId, packetSeq, AckPacket(true, appAck));
            dispatcher.recvPacket(
                IbcReceiver(mars), IbcPacket(src, dest, packetSeq, payload, ZERO_HEIGHT, maxTimeout), validProof
            );
        }
    }

    // recvPacket emits a WriteTimeoutPacket if timestamp passes chain B's block time
    function test_timeout_timestamp() public {
        uint64 packetSeq = 1;
        IbcPacket memory pkt = IbcPacket(src, dest, packetSeq, payload, ZERO_HEIGHT, 1);
        vm.expectEmit(true, true, true, true, address(dispatcher));
        emit RecvPacket(address(mars), channelId, packetSeq);
        vm.expectEmit(true, true, false, true, address(dispatcher));
        emit WriteTimeoutPacket(address(mars), channelId, packetSeq, pkt.timeoutHeight, pkt.timeoutTimestamp);
        dispatcher.recvPacket(IbcReceiver(mars), pkt, validProof);
    }

    // recvPacket emits a WriteTimeoutPacket if block height passes chain B's block height
    function test_timeout_blockHeight() public {
        uint64 packetSeq = 1;
        IbcPacket memory pkt = IbcPacket(src, dest, packetSeq, payload, Height(0, 1), 0);
        vm.expectEmit(true, true, true, true, address(dispatcher));
        emit RecvPacket(address(mars), channelId, packetSeq);
        vm.expectEmit(true, true, false, true, address(dispatcher));
        emit WriteTimeoutPacket(address(mars), channelId, packetSeq, pkt.timeoutHeight, pkt.timeoutTimestamp);
        dispatcher.recvPacket(IbcReceiver(mars), pkt, validProof);
    }

    // cannot receive packets out of order for ordered channel
    function test_outOfOrder() public {
        dispatcher.recvPacket(IbcReceiver(mars), IbcPacket(src, dest, 1, payload, ZERO_HEIGHT, maxTimeout), validProof);
        vm.expectRevert(abi.encodeWithSignature("unexpectedPacketSequence()"));
        dispatcher.recvPacket(IbcReceiver(mars), IbcPacket(src, dest, 3, payload, ZERO_HEIGHT, maxTimeout), validProof);
    }

    // TODO: add tests for unordered channel, wrong port, and invalid proof
}

// Test Chain A receives an acknowledgement packet from Chain B
contract DispatcherAckPacketTestSuite is PacketSenderTestBase {
    function test_success() public {
        for (uint64 index = 0; index < 3; index++) {
            sendPacket();

            vm.expectEmit(true, true, false, true, address(dispatcher));
            emit Acknowledgement(address(mars), channelId, sentPacket.sequence);
            dispatcher.acknowledgement(IbcReceiver(mars), sentPacket, ackPacket, validProof);
            // confirm dapp recieved the ack
            (bool success, bytes memory data) = mars.ackPackets(sentPacket.sequence - 1);
            AckPacket memory parsed = Ibc.parseAckData(ackPacket);
            assertEq(success, parsed.success);
            assertEq(data, parsed.data);
        }
    }

    // cannot ack packets if packet commitment is missing
    function test_missingPacket() public {
        vm.expectRevert(abi.encodeWithSignature("packetCommitmentNotFound()"));
        dispatcher.acknowledgement(IbcReceiver(mars), genPacket(1), genAckPacket("1"), validProof);

        sendPacket();
        dispatcher.acknowledgement(IbcReceiver(mars), sentPacket, ackPacket, validProof);

        // packet commitment is removed after ack
        vm.expectRevert(abi.encodeWithSignature("packetCommitmentNotFound()"));
        dispatcher.acknowledgement(IbcReceiver(mars), sentPacket, ackPacket, validProof);
    }

    // cannot recieve ack packets out of order for ordered channel
    function test_outOfOrder() public {
        for (uint64 index = 0; index < 3; index++) {
            sendPacket();
        }
        // 1st ack is ok
        dispatcher.acknowledgement(IbcReceiver(mars), genPacket(1), genAckPacket("1"), validProof);

        // only 2nd ack is allowed; so the 3rd ack fails
        vm.expectRevert(abi.encodeWithSignature("unexpectedPacketSequence()"));

        dispatcher.acknowledgement(IbcReceiver(mars), genPacket(3), genAckPacket("3"), validProof);
    }

    function test_invalidPort() public {
        Mars earth = new Mars(dispatcher);
        string memory earthPort = string(abi.encodePacked(portPrefix, getHexBytes(address(earth))));
        IbcEndpoint memory earthEnd = IbcEndpoint(earthPort, channelId);

        sendPacket();

        // another valid packet but not the same port
        IbcPacket memory packetEarth = sentPacket;
        packetEarth.src = earthEnd;

        vm.expectRevert(abi.encodeWithSignature("receiverNotOriginPacketSender()"));
        dispatcher.acknowledgement(IbcReceiver(mars), packetEarth, ackPacket, validProof);
    }

    // ackPacket fails if channel doesn't match
    function test_invalidChannel() public {
        sendPacket();

        IbcEndpoint memory invalidSrc = IbcEndpoint(src.portId, "channel-invalid");
        IbcPacket memory packet = sentPacket;
        packet.src = invalidSrc;

        vm.expectRevert(abi.encodeWithSignature("packetCommitmentNotFound()"));
        dispatcher.acknowledgement(IbcReceiver(mars), packet, ackPacket, validProof);
    }
}

// Test Chain A receives a timeout packet from Chain B
contract DispatcherTimeoutPacketTestSuite is PacketSenderTestBase {
    // preconditions for timeout packet
    // - packet commitment exists
    // - packet timeout is verified by Polymer client
    function test_success() public {
        for (uint64 index = 0; index < 3; index++) {
            sendPacket();

            vm.expectEmit(true, true, true, true, address(dispatcher));
            emit Timeout(address(mars), channelId, sentPacket.sequence);
            dispatcher.timeout(IbcReceiver(mars), sentPacket, validProof);
        }
    }

    // cannot timeout packets if packet commitment is missing
    function test_missingPacket() public {
        vm.expectRevert(abi.encodeWithSignature("packetCommitmentNotFound()"));
        dispatcher.timeout(IbcReceiver(mars), genPacket(1), validProof);

        sendPacket();
        dispatcher.timeout(IbcReceiver(mars), sentPacket, validProof);

        // packet commitment is removed after timeout
        vm.expectRevert(abi.encodeWithSignature("packetCommitmentNotFound()"));
        dispatcher.timeout(IbcReceiver(mars), sentPacket, validProof);
    }

    // cannot timeout packets if original packet port doesn't match current port
    function test_invalidPort() public {
        Mars earth = new Mars(dispatcher);
        string memory earthPort = string(abi.encodePacked(portPrefix, getHexBytes(address(earth))));
        IbcEndpoint memory earthEnd = IbcEndpoint(earthPort, channelId);

        sendPacket();

        // another valid packet but not the same port
        IbcPacket memory packetEarth = sentPacket;
        packetEarth.src = earthEnd;

        vm.expectRevert(abi.encodeWithSignature("receiverNotIntendedPacketDestination()"));
        dispatcher.timeout(IbcReceiver(mars), packetEarth, validProof);
    }

    // cannot timeout packetsfails if channel doesn't match
    function test_invalidChannel() public {
        sendPacket();

        IbcEndpoint memory invalidSrc = IbcEndpoint(src.portId, "channel-invalid");
        IbcPacket memory packet = sentPacket;
        packet.src = invalidSrc;

        vm.expectRevert(abi.encodeWithSignature("packetCommitmentNotFound()"));
        /* vm.expectRevert('Packet commitment not found'); */
        dispatcher.timeout(IbcReceiver(mars), packet, validProof);
    }

    // cannot timeout packets if proof from Polymer is invalid
    function test_invalidProof() public {
        sendPacket();

        // vm.expectRevert("Invalid dummy non membership proof");
        vm.expectRevert(DummyLightClient.InvalidDummyNonMembershipProof.selector);
        dispatcher.timeout(IbcReceiver(mars), sentPacket, invalidProof);
    }
}
