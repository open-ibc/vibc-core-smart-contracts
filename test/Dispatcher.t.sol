// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "../contracts/Ibc.sol";
import {Dispatcher, InitClientMsg, UpgradeClientMsg} from "../contracts/Dispatcher.sol";
import {IbcEventsEmitter} from "../contracts/IbcDispatcher.sol";
import {Escrow} from "../contracts/Escrow.sol";
import {IbcReceiver} from "../contracts/IbcReceiver.sol";
import "../contracts/IbcVerifier.sol";
import "../contracts/Verifier.sol";
import "../contracts/Mars.sol";
import "../contracts/OpConsensusStateManager.sol";

contract Base is Test, IbcEventsEmitter {
    uint64 UINT64_MAX = 18446744073709551615;

    ConsensusState untrustedState = ConsensusState(
        80990, 590199110038530808131927832294665177527506259518072095333098842116767351199, 7000040, 1000
    );
    ConsensusState trustedState = ConsensusState(
        10934, 7064503680087416120706887577693908749828198688716609274705703517077803898371, 7002040, 1020
    );
    InitClientMsg initClientMsg = InitClientMsg(bytes("Polymer"), untrustedState);

    ZKMintVerifier verifier = new Verifier();
    ZkProof proof = ZkProof(
        [
            13449388914393258752883032560537386278857542833249142697090243871249761501123,
            5963894333042515966217276339656894890750758020607775733717462915787234629927
        ],
        [
            [
                4811559872397934450173412387101758297072581261546553338353504577141293696514,
                18400634037991283592418145553628778322894277765243742619561207628896194710939
            ],
            [
                17903685207039300384995795331083569887497623740119108595975170464164316221644,
                9246628133289276308945311797896077179503891414159382179119544904154776510385
            ]
        ],
        [
            17432552394458345841788798376121543520587716339044416231610790827968220675517,
            15082220514596158133191868403239442750535261032426474092101151620016661078026
        ]
    );
    Height ZERO_HEIGHT = Height(0, 0);
    uint64 maxTimeout = UINT64_MAX;
    Escrow escrow = new Escrow();

    OptimisticConsensusStateManager opConsensusStateManager = new OptimisticConsensusStateManager(1800, verifier);

    // Proofs from Polymer chain, to verify packet or channel state on Polymer
    Proof emptyProof;
    Proof invalidProof = Proof(Height(0, 42), bytes("")); // invalid proof with empty proof bytes
    Proof validProof = Proof(Height(0, 42), bytes("valid proof"));

    Dispatcher dispatcher;
    string portPrefix = "polyibc.eth.";

    // ⬇️ Utility functions for testing

    // getHexStr converts an address to a hex string without the leading 0x
    function getHexBytes(address addr) internal pure returns (bytes memory) {
        bytes memory addrWithPrefix = abi.encodePacked(vm.toString(addr));
        bytes memory addrWithoutPrefix = new bytes(addrWithPrefix.length - 2);
        for (uint256 i = 0; i < addrWithoutPrefix.length; i++) {
            addrWithoutPrefix[i] = addrWithPrefix[i + 2];
        }
        return addrWithoutPrefix;
    }

    // deriveAddress derives an address from a given string deterministically for testing
    function deriveAddress(string memory str) internal pure returns (address) {
        return vm.addr(uint256(keccak256(abi.encodePacked(str))));
    }

    function mergeFees(PacketFee memory a, PacketFee memory b) internal pure returns (PacketFee memory) {
        return PacketFee(a.recvFee + b.recvFee, a.ackFee + b.ackFee, a.timeoutFee + b.timeoutFee);
    }
}

contract DispatcherCreateClientTest is Test, Base {
    function setUp() public {
        dispatcher = new Dispatcher(verifier, escrow, portPrefix, opConsensusStateManager);
    }

    function test_success() public {
        dispatcher.createClient(initClientMsg);
    }

    function test_mustByOwner() public {
        vm.prank(deriveAddress("non-onwer"));
        vm.expectRevert("Ownable: caller is not the owner");
        dispatcher.createClient(initClientMsg);
    }

    function test_onlyOnce() public {
        dispatcher.createClient(initClientMsg);
        vm.expectRevert(abi.encodeWithSignature("clientAlreadyCreated()"));
        dispatcher.createClient(initClientMsg);
    }
}

contract DispatcherUpdateClientTest is Test, Base {
    function setUp() public {
        dispatcher = new Dispatcher(verifier, escrow, portPrefix, opConsensusStateManager);
        dispatcher.createClient(initClientMsg);
    }

    function test_updateConsensusState_success() public {
        dispatcher.updateClientWithConsensusState(trustedState, proof);
    }

    function test_updateConsensusState_invalid() public {
        vm.expectRevert(abi.encodeWithSignature("consensusStateVerificationFailed()"));
        dispatcher.updateClientWithConsensusState(untrustedState, proof);

        vm.expectRevert(abi.encodeWithSignature("consensusStateVerificationFailed()"));
        ConsensusState memory invalidConsensusState;
        dispatcher.updateClientWithConsensusState(invalidConsensusState, proof);
    }
}

contract DispatcherUpgradeClientTest is Test, Base {
    function setUp() public {
        dispatcher = new Dispatcher(verifier, escrow, portPrefix, opConsensusStateManager);
        dispatcher.createClient(initClientMsg);
        dispatcher.updateClientWithConsensusState(trustedState, proof);
    }

    function test_success() public {
        dispatcher.upgradeClient(UpgradeClientMsg(bytes("upgradeOptimisticConsensusState"), trustedState));
    }

    function test_ownerOnly() public {
        vm.prank(deriveAddress("non-onwer"));
        vm.expectRevert("Ownable: caller is not the owner");
        dispatcher.upgradeClient(UpgradeClientMsg(bytes("upgradeOptimisticConsensusState"), trustedState));
    }
}

struct VersionSet {
    string self;
    string counterparty;
    string expected;
}

contract ChannelTestBase is Test, Base {
    Mars mars;
    string[] connectionHops;
    VersionSet version = VersionSet("1.0", "1.0", "1.0");
    ChannelOrder ordering = ChannelOrder.ORDERED;
    bool feeEnabled = false;
    CounterParty cp;

    CounterParty cpBsc = CounterParty("polyibc.bsc.9876543210", bytes32("channel-99"), "1.0");

    function setUp() public virtual {
        connectionHops = new string[](2);
        connectionHops[0] = "connection-1";
        connectionHops[1] = "connection-2";

        dispatcher = new Dispatcher(verifier, escrow, portPrefix, opConsensusStateManager);
        mars = new Mars(dispatcher);

        dispatcher.createClient(initClientMsg);
        dispatcher.updateClientWithConsensusState(trustedState, proof);

        // channel handshake methods use optimistic consensus state
        // for membership verification. it is necessary to update the
        // optimistic consensus state and wait for it to pass the
        // fraud proof window before proceeeding.
        dispatcher.updateClientWithOptimisticConsensusState(validProof.proofHeight.revision_height, 1);

        // forward the time to make the consensus state transition to trusted state
        vm.warp(block.timestamp + 1801);

        polymerProof = validProof;
        vm.startPrank(deriveAddress("relayer"));
    }

    Proof polymerProof;

    function setTestParams(VersionSet memory _version, CounterParty memory _cp, ChannelOrder _ordering) internal {
        version = _version;
        _cp.version = version.counterparty;
        cp = _cp;
        ordering = _ordering;

        // 1st handshake call with unknown counterparty version or channelId
        if (cp.channelId == bytes32(0x0) && keccak256(abi.encodePacked(cp.version)) == keccak256(abi.encodePacked("")))
        {
            polymerProof = emptyProof;
        } else {
            // all other handshake calls requires a polymer proof
            polymerProof = validProof;
        }
    }

    modifier goodCases() {
        // NOTE: counterparty version is set in VersionSet
        ChannelOrder[2] memory orderings = [ChannelOrder.ORDERED, ChannelOrder.UNORDERED];
        // 1st handshake with unknown counterparty version or channelId
        cp = CounterParty("polyibc.bsc.9876543210", bytes32(0x0), "");
        VersionSet[2] memory versions = [VersionSet("1.0", "", "1.0"), VersionSet("2.0", "", "2.0")];
        for (uint256 i = 0; i < versions.length; i++) {
            for (uint256 j = 0; j < orderings.length; j++) {
                setTestParams(versions[i], cp, orderings[j]);
                _;
            }
        }
        // 2nd handshake with known counterparty version and channelId
        cp = CounterParty("polyibc.bsc.9876543210", bytes32("channel-99"), "");
        versions = [VersionSet("", "1.0", "1.0"), VersionSet("", "2.0", "2.0")];
        for (uint256 i = 0; i < versions.length; i++) {
            for (uint256 j = 0; j < orderings.length; j++) {
                setTestParams(versions[i], cp, orderings[j]);
                _;
            }
        }
    }
}

contract DispatcherOpenIbcChannelTest is ChannelTestBase {
    function test_success() public goodCases {
        vm.expectEmit(true, true, true, true);
        emit OpenIbcChannel(
            address(mars), version.expected, ordering, feeEnabled, connectionHops, cp.portId, cp.channelId
        );
        dispatcher.openIbcChannel(
            IbcReceiver(mars), version.self, ordering, feeEnabled, connectionHops, cp, polymerProof
        );
    }

    modifier unsupportedVersions() {
        // NOTE: counterparty version is set in VersionSet
        ChannelOrder[2] memory orderings = [ChannelOrder.ORDERED, ChannelOrder.UNORDERED];
        // 1st handshake with unknown counterparty version or channelId
        cp = CounterParty("polyibc.bsc.9876543210", bytes32(0x0), "");
        VersionSet[1] memory versions = [VersionSet("", "", "")];
        for (uint256 i = 0; i < versions.length; i++) {
            for (uint256 j = 0; j < orderings.length; j++) {
                setTestParams(versions[i], cp, orderings[j]);
                _;
            }
        }
        // 2nd handshake with known counterparty version and channelId
        cp = CounterParty("polyibc.bsc.9876543210", bytes32("channel-99"), "xxx");
        versions = [VersionSet("", "xxx", "")];
        for (uint256 i = 0; i < versions.length; i++) {
            for (uint256 j = 0; j < orderings.length; j++) {
                setTestParams(versions[i], cp, orderings[j]);
                _;
            }
        }
    }

    function test_unsupportedVersion() public unsupportedVersions {
        vm.expectRevert(bytes("Unsupported version"));
        dispatcher.openIbcChannel(
            IbcReceiver(mars), version.self, ordering, feeEnabled, connectionHops, cp, polymerProof
        );
    }

    function test_invalidCounterpartyPortId() public {
        CounterParty[4] memory cps = [
            CounterParty("", bytes32(0x0), ""),
            CounterParty("portX", bytes32(0x0), ""),
            CounterParty("", bytes32("channel-99"), "1.0"),
            CounterParty("portX", bytes32("channel-99"), "1.0")
        ];
        for (uint256 i = 0; i < cps.length; i++) {
            cp = cps[i];
            vm.expectRevert(abi.encodeWithSignature("invalidCounterPartyPortId()"));
            dispatcher.openIbcChannel(
                IbcReceiver(mars), version.self, ordering, feeEnabled, connectionHops, cp, polymerProof
            );
        }
    }

    function test_invalidProof() public {
        vm.expectRevert("Fail to prove channel state");
        dispatcher.openIbcChannel(
            IbcReceiver(mars), "", ChannelOrder.ORDERED, feeEnabled, connectionHops, cpBsc, invalidProof
        );
    }
}

contract DispatcherConnectIbcChannelTest is ChannelTestBase {
    function test_success() public goodCases {
        dispatcher.openIbcChannel(
            IbcReceiver(mars), version.self, ordering, feeEnabled, connectionHops, cp, polymerProof
        );
        string memory cpVersion =
            keccak256(abi.encode(cp.version)) == keccak256(abi.encode(bytes(""))) ? version.self : cp.version;

        bytes32 channelId = "channel-1"; // generated by IBC core on Polymer chain
        vm.expectEmit(true, true, true, true);
        emit ConnectIbcChannel(address(mars), channelId);
        dispatcher.connectIbcChannel(
            IbcReceiver(mars),
            channelId,
            connectionHops,
            ordering,
            feeEnabled,
            cpBsc.portId,
            bytes32("channel-99"),
            cpVersion,
            validProof
        );
    }

    function test_unsupportedVersions() public goodCases {
        dispatcher.openIbcChannel(
            IbcReceiver(mars), version.self, ordering, feeEnabled, connectionHops, cp, polymerProof
        );
        string memory cpVersion = "xxx";

        vm.expectRevert("Unsupported version");
        dispatcher.connectIbcChannel(
            IbcReceiver(mars),
            bytes32("channel-1"),
            connectionHops,
            ordering,
            feeEnabled,
            cp.portId,
            bytes32("channel-99"),
            cpVersion,
            validProof
        );
    }

    function test_invalidProof() public {
        vm.expectRevert("Fail to prove channel state");
        dispatcher.connectIbcChannel(
            IbcReceiver(mars),
            bytes32("channel-1"),
            connectionHops,
            ordering,
            feeEnabled,
            cpBsc.portId,
            cpBsc.channelId,
            cpBsc.version,
            invalidProof
        );
    }
}

// This Base contract provides an open channel for sub-contract tests
contract ChannelOpenTestBase is Test, Base {
    Mars mars;
    bytes32 channelId = "channel-1";
    address relayer = deriveAddress("relayer");
    bool feeEnabled = false;

    function setUp() public virtual {
        string[] memory connectionHops = new string[](2);
        connectionHops[0] = "connection-1";
        connectionHops[1] = "connection-2";

        dispatcher = new Dispatcher(verifier, escrow, portPrefix, opConsensusStateManager);
        dispatcher.createClient(initClientMsg);

        Escrow myEscrow = Escrow(escrow);
        myEscrow.setDispatcher(address(dispatcher));

        // anyone can run Relayers
        vm.startPrank(relayer);
        vm.deal(relayer, 100000 ether);
        mars = new Mars(dispatcher);

        dispatcher.updateClientWithConsensusState(trustedState, proof);

        // finish channel handshake as chain A
        CounterParty memory cp = CounterParty("polyibc.bsc.9876543210", bytes32(0x0), "");
        dispatcher.openIbcChannel(
            IbcReceiver(mars), "1.0", ChannelOrder.ORDERED, feeEnabled, connectionHops, cp, emptyProof
        );
        CounterParty memory cp2 = CounterParty("polyibc.bsc.9876543210", bytes32("channel-99"), "1.0");

        // channel handshake methods use optimistic consensus state
        // for membership verification. it is necessary to update the
        // optimistic consensus state and wait for it to pass the
        // fraud proof window before proceeeding.
        dispatcher.updateClientWithOptimisticConsensusState(validProof.proofHeight.revision_height, 1);

        // forward the time to make the consensus state transition to trusted state
        vm.warp(block.timestamp + 1801);

        dispatcher.connectIbcChannel(
            IbcReceiver(mars),
            channelId,
            connectionHops,
            ChannelOrder.ORDERED,
            feeEnabled,
            cp2.portId,
            cp2.channelId,
            cp2.version,
            validProof
        );
    }
}

contract DispatcherCloseChannelTest is ChannelOpenTestBase {
    function test_closeChannelInit_success() public {
        vm.expectEmit(true, true, true, true);
        emit CloseIbcChannel(address(mars), channelId);
        mars.triggerChannelClose(channelId);
    }

    function test_closeChannelInit_mustOwner() public {
        Mars earth = new Mars(dispatcher);
        vm.expectRevert(abi.encodeWithSignature("channelNotOwnedBySender()"));
        earth.triggerChannelClose(channelId);
    }

    function test_closeChannelConfirm_success() public {
        vm.expectEmit(true, true, true, true);
        emit CloseIbcChannel(address(mars), channelId);
        dispatcher.onCloseIbcChannel(address(mars), channelId, validProof);
    }

    function test_closeChannelConfirm_mustOwner() public {
        vm.expectRevert(abi.encodeWithSignature("channelNotOwnedByPortAddress()"));
        dispatcher.onCloseIbcChannel(address(mars), "channel-999", validProof);
    }

    function test_closeChannelConfirm_invalidProof() public {
        vm.expectRevert("Fail to prove channel state");
        dispatcher.onCloseIbcChannel(address(mars), channelId, invalidProof);
    }
}

contract DispatcherSendPacketTest is ChannelOpenTestBase {
    // default params
    string payload = "msgPayload";
    uint64 timeoutTimestamp = 1000;
    PacketFee fee = PacketFee(1 ether, 2 ether, 3 ether);

    function test_success() public {
        bytes memory packet = abi.encodePacked(payload);
        for (uint64 index = 0; index < 3; index++) {
            uint256 balance = address(escrow).balance;
            vm.expectEmit(true, true, true, true);
            uint64 packetSeq = index + 1;
            emit SendPacket(address(mars), channelId, packet, packetSeq, timeoutTimestamp, fee);
            mars.greet{value: Ibc.calcEscrowFee(fee)}(payload, channelId, timeoutTimestamp, fee);
            assertEq(address(escrow).balance - balance, Ibc.calcEscrowFee(fee));

            // query escrowed fee per packet
            PacketFee memory packetFee = dispatcher.getTotalPacketFees(address(mars), channelId, packetSeq);
            assertEq(keccak256(abi.encode(packetFee)), keccak256(abi.encode(fee)));
        }
    }

    // sendPacket fails if calling dApp doesn't own the channel
    function test_mustOwner() public {
        Mars earth = new Mars(dispatcher);
        vm.expectRevert(abi.encodeWithSignature("channelNotOwnedBySender()"));
        earth.greet{value: Ibc.calcEscrowFee(fee)}(payload, channelId, timeoutTimestamp, fee);
    }
}

contract PacketSenderTest is ChannelOpenTestBase {
    IbcEndpoint dest = IbcEndpoint("polyibc.bsc.9876543210", "channel-99");
    IbcEndpoint src;
    string payloadStr = "msgPayload";
    bytes payload = bytes(payloadStr);
    bytes appAck = abi.encodePacked('{ "account": "account", "reply": "got the message" }');
    PacketFee fee = PacketFee(1 ether, 2 ether, 3 ether);

    uint64 nextSendSeq = 1;
    // cached packet that was sent in `sendPacket`
    IbcPacket sentPacket;
    // ackPacket is the acknowledgement packet that is expected to be written for the `sentPacket`
    AckPacket ackPacket;

    function setUp() public virtual override {
        super.setUp();
        string memory marsPort = string(abi.encodePacked(portPrefix, getHexBytes(address(mars))));
        src = IbcEndpoint(marsPort, channelId);
    }

    // sendPacket writes a packet commitment, and updates cached `sentPacket` and `ackPacket`
    function sendPacket() internal {
        sentPacket = genPacket(nextSendSeq);
        ackPacket = genAckPacket(nextSendSeq);
        mars.greet{value: Ibc.calcEscrowFee(fee)}(payloadStr, channelId, maxTimeout, fee);
        nextSendSeq += 1;
    }

    // genPacket generates a packet for the given packet sequence
    function genPacket(uint64 packetSeq) internal view returns (IbcPacket memory) {
        return IbcPacket(src, dest, packetSeq, payload, ZERO_HEIGHT, maxTimeout);
    }

    // genAckPacket generates an ack packet for the given packet sequence
    function genAckPacket(uint64 packetSeq) internal pure returns (AckPacket memory) {
        return AckPacket(true, abi.encodePacked("ackPacket", packetSeq));
    }
}

// Test Chains B receives a packet from Chain A
contract DispatcherRecvPacketTest is ChannelOpenTestBase {
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
contract DispatcherAckPacketTest is PacketSenderTest {
    function test_success() public {
        for (uint64 index = 0; index < 3; index++) {
            sendPacket();

            vm.expectEmit(true, true, false, true, address(dispatcher));
            emit Acknowledgement(address(mars), channelId, sentPacket.sequence);
            dispatcher.acknowledgement(IbcReceiver(mars), sentPacket, ackPacket, validProof);
            // confirm dapp recieved the ack
            (bool success, bytes memory data) = mars.ackPackets(sentPacket.sequence - 1);
            assertEq(success, ackPacket.success);
            assertEq(data, ackPacket.data);
        }
    }

    // cannot ack packets if packet commitment is missing
    function test_missingPacket() public {
        vm.expectRevert(abi.encodeWithSignature("packetCommitmentNotFound()"));
        dispatcher.acknowledgement(IbcReceiver(mars), genPacket(1), genAckPacket(1), validProof);

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
        dispatcher.acknowledgement(IbcReceiver(mars), genPacket(1), genAckPacket(1), validProof);

        // only 2nd ack is allowed; so the 3rd ack fails
        vm.expectRevert(abi.encodeWithSignature("unexpectedPacketSequence()"));

        dispatcher.acknowledgement(IbcReceiver(mars), genPacket(3), genAckPacket(3), validProof);
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

    function test_no_incentivizedAck() public {
        sendPacket();
        IncentivizedAckPacket memory incAck =
            IncentivizedAckPacket(true, abi.encodePacked(deriveAddress("foward-relayer")), bytes("ack"));

        vm.expectRevert(abi.encodeWithSignature("invalidChannelType(string)", "non-incentivized"));
        dispatcher.incentivizedAcknowledgement(IbcReceiver(mars), sentPacket, incAck, validProof);
    }
}

// Test Chain A receives a timeout packet from Chain B
contract DispatcherTimeoutPacketTest is PacketSenderTest {
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

        vm.expectRevert(abi.encodeWithSignature("receiverNotIndtendedPacketDestination()"));
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

        vm.expectRevert("Fail to prove timeout");
        dispatcher.timeout(IbcReceiver(mars), sentPacket, invalidProof);
    }
}
