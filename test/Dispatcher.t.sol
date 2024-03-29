// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "../contracts/libs/Ibc.sol";
import {Dispatcher} from "../contracts/core/Dispatcher.sol";
import {IbcEventsEmitter} from "../contracts/interfaces/IbcDispatcher.sol";
import {IbcReceiver, IbcReceiverBase} from "../contracts/interfaces/IbcReceiver.sol";
import {DummyLightClient} from "../contracts/utils/DummyLightClient.sol";
import {
    Mars,
    RevertingBytesMars,
    PanickingMars,
    RevertingEmptyMars,
    RevertingStringMars
} from "../contracts/examples/Mars.sol";
import "../contracts/core/OpLightClient.sol";
import "./Dispatcher.base.t.sol";
import {Earth} from "../contracts/examples/Earth.sol";

abstract contract ChannelHandshakeUtils is Base {
    string portId;
    LocalEnd _local;
    Mars mars;
    ChannelEnd _remote;

    function createSettings(bool localInitiate, bool isProofValid)
        internal
        view
        returns (ChannelHandshakeSetting[4] memory)
    {
        Ics23Proof memory proof = isProofValid ? validProof : invalidProof;
        ChannelHandshakeSetting[4] memory settings = [
            ChannelHandshakeSetting(ChannelOrder.ORDERED, false, localInitiate, proof),
            ChannelHandshakeSetting(ChannelOrder.UNORDERED, false, localInitiate, proof),
            ChannelHandshakeSetting(ChannelOrder.ORDERED, true, localInitiate, proof),
            ChannelHandshakeSetting(ChannelOrder.UNORDERED, true, localInitiate, proof)
        ];
        return settings;
    }

    function createSettings2(bool isProofValid) internal view returns (ChannelHandshakeSetting[8] memory) {
        // localEnd initiates
        ChannelHandshakeSetting[4] memory settings1 = createSettings(true, isProofValid);
        // remoteEnd initiates
        ChannelHandshakeSetting[4] memory settings2 = createSettings(false, isProofValid);
        ChannelHandshakeSetting[8] memory settings;
        for (uint256 i = 0; i < settings1.length; i++) {
            settings[i] = settings1[i];
            settings[i + settings1.length] = settings2[i];
        }
        return settings;
    }
}

abstract contract ChannelHandshakeTestSuite is ChannelHandshakeUtils {
    function test_openChannel_initiator_ok() public {
        ChannelHandshakeSetting[4] memory settings = createSettings(true, true);
        string[2] memory versions = ["1.0", "2.0"];
        for (uint256 i = 0; i < settings.length; i++) {
            for (uint256 j = 0; j < versions.length; j++) {
                LocalEnd memory le = _local;
                ChannelEnd memory re = _remote;
                le.versionCall = versions[j];
                le.versionExpected = versions[j];
                // remoteEnd has no channelId or version if localEnd is the initiator
                channelOpenInit(le, re, settings[i], true);
            }
        }
    }

    function test_openChannel_receiver_ok() public {
        ChannelHandshakeSetting[4] memory settings = createSettings(false, true);
        string[2] memory versions = ["1.0", "2.0"];
        for (uint256 i = 0; i < settings.length; i++) {
            for (uint256 j = 0; j < versions.length; j++) {
                LocalEnd memory le = _local;
                ChannelEnd memory re = _remote;
                re.version = versions[j];
                // explicit version
                le.versionCall = versions[j];
                le.versionExpected = versions[j];
                // remoteEnd version is used
                channelOpenInit(le, re, settings[i], true);

                // auto version selection
                le.versionCall = "";
                channelOpenTry(le, re, settings[i], true);
            }
        }
    }

    function test_connectChannel_ok() public {
        ChannelHandshakeSetting[8] memory settings = createSettings2(true);

        string[2] memory versions = ["1.0", "2.0"];
        for (uint256 i = 0; i < settings.length; i++) {
            for (uint256 j = 0; j < versions.length; j++) {
                LocalEnd memory le = _local;
                ChannelEnd memory re = _remote;
                le.versionCall = versions[j];
                le.versionExpected = versions[j];
                re.version = versions[j];
                channelOpenInit(le, re, settings[i], true);
                channelOpenTry(le, re, settings[i], true);
                channelOpenAck(le, re, settings[i], true);
                channelOpenConfirm(le, re, settings[i], true);
            }
        }
    }

    function test_openChannel_initiator_revert_unsupportedVersion() public {
        ChannelHandshakeSetting[4] memory settings = createSettings(true, true);
        string[2] memory versions = ["", "xxxxxxx"];
        for (uint256 i = 0; i < settings.length; i++) {
            for (uint256 j = 0; j < versions.length; j++) {
                LocalEnd memory le = _local;
                ChannelEnd memory re = _remote;
                le.versionCall = versions[j];
                le.versionExpected = versions[j];
                vm.expectEmit(true, true, true, true);
                emit IbcEventsEmitter.ChannelOpenInitError(
                    address(le.receiver), abi.encodeWithSelector(IbcReceiverBase.UnsupportedVersion.selector)
                );
                channelOpenInit(le, re, settings[i], false);
            }
        }
    }

    function test_openChannel_receiver_fail_invalidProof() public {
        // When localEnd initiates, no proof verification is done in channelOpenTry
        ChannelHandshakeSetting[4] memory settings = createSettings(false, false);
        string[1] memory versions = ["1.0"];
        for (uint256 i = 0; i < settings.length; i++) {
            for (uint256 j = 0; j < versions.length; j++) {
                LocalEnd memory le = _local;
                ChannelEnd memory re = _remote;
                le.versionCall = versions[j];
                le.versionExpected = versions[j];
                vm.expectRevert(DummyLightClient.InvalidDummyMembershipProof.selector);
                channelOpenTry(le, re, settings[i], false);
            }
        }
    }

    function test_connectChannel_fail_unsupportedVersion() public {
        // When localEnd initiates, counterparty version is only available in channelOpenAck
        ChannelHandshakeSetting[4] memory settings = createSettings(true, true);
        string[2] memory versions = ["", "xxxxxxx"];
        for (uint256 i = 0; i < settings.length; i++) {
            for (uint256 j = 0; j < versions.length; j++) {
                LocalEnd memory le = _local;
                ChannelEnd memory re = _remote;
                // no remote version applied in openChannel
                channelOpenInit(le, re, settings[i], true);
                channelOpenTry(le, re, settings[i], true);
                re.version = versions[j];
                vm.expectEmit(true, true, true, true);
                emit IbcEventsEmitter.ChannelOpenAckError(
                    address(le.receiver), abi.encodeWithSelector(IbcReceiverBase.UnsupportedVersion.selector)
                );

                channelOpenAck(le, re, settings[i], false);
            }
        }
    }

    function test_connectChannel_fail_invalidProof() public {
        // When localEnd initiates, counterparty version is only available in channelOpenAck
        ChannelHandshakeSetting[8] memory settings = createSettings2(true);
        string[1] memory versions = ["1.0"];
        for (uint256 i = 0; i < settings.length; i++) {
            for (uint256 j = 0; j < versions.length; j++) {
                LocalEnd memory le = _local;
                ChannelEnd memory re = _remote;
                // no remote version applied in openChannel
                channelOpenInit(le, re, settings[i], true);
                channelOpenTry(le, re, settings[i], true);
                re.version = versions[j];
                settings[i].proof = invalidProof;
                vm.expectRevert(DummyLightClient.InvalidDummyMembershipProof.selector);
                channelOpenAck(le, re, settings[i], false);
            }
        }
    }
}

contract ChannelHandshakeTest is ChannelHandshakeTestSuite {
    function setUp() public virtual override {
        (dispatcherProxy, dispatcherImplementation) = deployDispatcherProxyAndImpl(portPrefix);
        dispatcherProxy.setClientForConnection(connectionHops[0], dummyLightClient);
        mars = new Mars(dispatcherProxy);
        portId = IbcUtils.addressToPortId(portPrefix, address(mars));
        _local = LocalEnd(mars, portId, "channel-1", connectionHops, "1.0", "1.0");
        _remote = ChannelEnd("eth2.7E5F4552091A69125d5DfCb7b8C2659029395Bdf", "channel-2", "1.0");
    }
}

// This Base contract provides an open channel for sub-contract tests
contract ChannelOpenTestBaseSetup is Base {
    string portId;
    string revertingBytesMarsPortId;
    string invalidPortId = "eth1.0xd6292A04e605AFf917Bf05b2df5dDdbdc3E35e07";
    bytes32 channelId = "channel-1";
    address relayer = deriveAddress("relayer");
    bool feeEnabled = false;

    LocalEnd _local;
    LocalEnd _localRevertingMars;
    ChannelEnd _remote;
    Mars mars;
    RevertingBytesMars revertingBytesMars;

    function setUp() public virtual override {
        (dispatcherProxy, dispatcherImplementation) = deployDispatcherProxyAndImpl(portPrefix);
        dispatcherProxy.setClientForConnection(connectionHops[0], dummyLightClient);
        ChannelHandshakeSetting memory setting =
            ChannelHandshakeSetting(ChannelOrder.ORDERED, feeEnabled, true, validProof);

        // anyone can run Relayers
        vm.startPrank(relayer);
        vm.deal(relayer, 100_000 ether);
        mars = new Mars(dispatcherProxy);
        portId = IbcUtils.addressToPortId(portPrefix, address(mars));
        revertingBytesMars = new RevertingBytesMars(dispatcherProxy);
        revertingBytesMarsPortId = IbcUtils.addressToPortId(portPrefix, address(revertingBytesMars));

        _local = LocalEnd(mars, portId, channelId, connectionHops, "1.0", "1.0");
        _localRevertingMars =
            LocalEnd(revertingBytesMars, revertingBytesMarsPortId, channelId, connectionHops, "1.0", "1.0");
        _remote = ChannelEnd("eth2.7E5F4552091A69125d5DfCb7b8C2659029395Bdf", "channel-2", "1.0");

        vm.stopPrank();
        channelOpenInit(_local, _remote, setting, true);
        vm.startPrank(relayer);
        channelOpenTry(_localRevertingMars, _remote, setting, true);

        channelOpenAck(_local, _remote, setting, true);
        channelOpenConfirm(_localRevertingMars, _remote, setting, true);
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
        Mars earth = new Mars(dispatcherProxy);
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.channelNotOwnedBySender.selector));
        earth.greet(payload, channelId, timeoutTimestamp);
    }
}

contract PacketSenderTestBase is ChannelOpenTestBaseSetup {
    IbcEndpoint dest = IbcEndpoint("polyibc.bsc.9876543210", "channel-99");
    IbcEndpoint src;
    IbcEndpoint srcRevertingMars;

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
        string memory revertingMarsPort = string(abi.encodePacked(portPrefix, getHexBytes(address(revertingBytesMars))));

        src = IbcEndpoint(marsPort, channelId);
        srcRevertingMars = IbcEndpoint(revertingMarsPort, channelId);
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
            vm.expectEmit(true, true, true, true, address(dispatcherProxy));
            emit RecvPacket(address(mars), channelId, packetSeq);
            vm.expectEmit(true, true, false, true, address(dispatcherProxy));
            emit WriteAckPacket(address(mars), channelId, packetSeq, AckPacket(true, appAck));
            dispatcherProxy.recvPacket(IbcPacket(src, dest, packetSeq, payload, ZERO_HEIGHT, maxTimeout), validProof);
        }
    }

    // recvPacket emits a WriteTimeoutPacket if timestamp passes chain B's block time
    function test_timeout_timestamp() public {
        uint64 packetSeq = 1;
        IbcPacket memory pkt = IbcPacket(src, dest, packetSeq, payload, ZERO_HEIGHT, 1);
        vm.expectEmit(true, true, true, true, address(dispatcherProxy));
        emit RecvPacket(address(mars), channelId, packetSeq);
        vm.expectEmit(true, true, false, true, address(dispatcherProxy));
        emit WriteTimeoutPacket(address(mars), channelId, packetSeq, pkt.timeoutHeight, pkt.timeoutTimestamp);
        dispatcherProxy.recvPacket(pkt, validProof);
    }

    // recvPacket emits a WriteTimeoutPacket if block height passes chain B's block height
    function test_timeout_blockHeight() public {
        uint64 packetSeq = 1;
        IbcPacket memory pkt = IbcPacket(src, dest, packetSeq, payload, Height(0, 1), 0);
        vm.expectEmit(true, true, true, true, address(dispatcherProxy));
        emit RecvPacket(address(mars), channelId, packetSeq);
        vm.expectEmit(true, true, false, true, address(dispatcherProxy));
        emit WriteTimeoutPacket(address(mars), channelId, packetSeq, pkt.timeoutHeight, pkt.timeoutTimestamp);
        dispatcherProxy.recvPacket(pkt, validProof);
    }

    // cannot receive packets out of order for ordered channel
    function test_outOfOrder() public {
        dispatcherProxy.recvPacket(IbcPacket(src, dest, 1, payload, ZERO_HEIGHT, maxTimeout), validProof);
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.unexpectedPacketSequence.selector));
        dispatcherProxy.recvPacket(IbcPacket(src, dest, 3, payload, ZERO_HEIGHT, maxTimeout), validProof);
    }

    // TODO: add tests for unordered channel, wrong port, and invalid proof
}

// Test Chain A receives an acknowledgement packet from Chain B
contract DispatcherAckPacketTestSuite is PacketSenderTestBase {
    function test_success() public {
        for (uint64 index = 0; index < 3; index++) {
            sendPacket();

            vm.expectEmit(true, true, false, true, address(dispatcherProxy));
            emit Acknowledgement(address(mars), channelId, sentPacket.sequence);
            dispatcherProxy.acknowledgement(sentPacket, ackPacket, validProof);
            // confirm dapp recieved the ack
            (bool success, bytes memory data) = mars.ackPackets(sentPacket.sequence - 1);
            AckPacket memory parsed = Ibc.parseAckData(ackPacket);
            assertEq(success, parsed.success);
            assertEq(data, parsed.data);
        }
    }

    // cannot ack packets if packet commitment is missing
    function test_missingPacket() public {
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.packetCommitmentNotFound.selector));
        dispatcherProxy.acknowledgement(genPacket(1), genAckPacket("1"), validProof);

        sendPacket();
        dispatcherProxy.acknowledgement(sentPacket, ackPacket, validProof);

        // packet commitment is removed after ack
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.packetCommitmentNotFound.selector));
        dispatcherProxy.acknowledgement(sentPacket, ackPacket, validProof);
    }

    // cannot recieve ack packets out of order for ordered channel
    function test_outOfOrder() public {
        for (uint64 index = 0; index < 3; index++) {
            sendPacket();
        }
        // 1st ack is ok
        dispatcherProxy.acknowledgement(genPacket(1), genAckPacket("1"), validProof);

        // only 2nd ack is allowed; so the 3rd ack fails
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.unexpectedPacketSequence.selector));

        dispatcherProxy.acknowledgement(genPacket(3), genAckPacket("3"), validProof);
    }

    function test_invalidPort() public {
        Mars earth = new Mars(dispatcherProxy);
        string memory earthPort = string(abi.encodePacked(portPrefix, getHexBytes(address(earth))));
        IbcEndpoint memory earthEnd = IbcEndpoint(earthPort, channelId);

        sendPacket();

        // another valid packet but not the same port
        IbcPacket memory packetEarth = sentPacket;
        packetEarth.src = earthEnd;

        vm.expectRevert(abi.encodeWithSelector(IBCErrors.packetCommitmentNotFound.selector));
        dispatcherProxy.acknowledgement(packetEarth, ackPacket, validProof);
    }

    function test_misformattedPort() public {
        Mars earth = new Mars(dispatcherProxy);
        string memory earthPort = string(abi.encodePacked(portPrefix, getHexBytes(address(earth))));
        IbcEndpoint memory earthEnd = IbcEndpoint(earthPort, channelId);

        sendPacket();

        // another valid packet but not the same port
        IbcPacket memory packetEarth = sentPacket;
        packetEarth.src = earthEnd;
        packetEarth.src.portId = string(bytes.concat(bytes(portPrefix), bytes32(uint256(uint160(vm.addr(1))))));

        vm.expectRevert(abi.encodeWithSelector(IBCErrors.invalidHexStringLength.selector));
        dispatcherProxy.acknowledgement(packetEarth, ackPacket, validProof);
    }

    // ackPacket fails if channel doesn't match
    function test_invalidChannel() public {
        sendPacket();

        IbcEndpoint memory invalidSrc = IbcEndpoint(src.portId, "channel-invalid");
        IbcPacket memory packet = sentPacket;
        packet.src = invalidSrc;

        vm.expectRevert(abi.encodeWithSelector(IBCErrors.channelIdNotFound.selector, packet.src.channelId));
        dispatcherProxy.acknowledgement(packet, ackPacket, validProof);
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

            vm.expectEmit(true, true, true, true, address(dispatcherProxy));
            emit Timeout(address(mars), channelId, sentPacket.sequence);
            dispatcherProxy.timeout(sentPacket, validProof);
        }
    }

    function test_timeout_dapp_revert() public {
        sentPacket = IbcPacket(srcRevertingMars, dest, 1, payload, ZERO_HEIGHT, maxTimeout);
        revertingBytesMars.greet(payloadStr, channelId, maxTimeout);
        nextSendSeq += 1;
        vm.expectEmit(true, true, true, true, address(dispatcherProxy));
        emit TimeoutError(
            address(revertingBytesMars), abi.encodeWithSelector(RevertingBytesMars.OnTimeoutPacket.selector)
        );
        dispatcherProxy.timeout(sentPacket, validProof);
    }

    // cannot timeout packets if packet commitment is missing
    function test_missingPacket() public {
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.packetCommitmentNotFound.selector));
        dispatcherProxy.timeout(genPacket(1), validProof);

        sendPacket();
        dispatcherProxy.timeout(sentPacket, validProof);

        // packet commitment is removed after timeout
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.packetCommitmentNotFound.selector));
        dispatcherProxy.timeout(sentPacket, validProof);
    }

    // cannot timeout packets if original packet port doesn't match current port
    function test_invalidPort_123a() public {
        Mars earth = new Mars(dispatcherProxy);
        string memory earthPort = string(abi.encodePacked(portPrefix, getHexBytes(address(earth))));
        IbcEndpoint memory earthEnd = IbcEndpoint(earthPort, channelId);

        sendPacket();

        // another valid packet but not the same port
        IbcPacket memory packetEarth = sentPacket;
        packetEarth.src = earthEnd;

        vm.expectRevert(IBCErrors.packetCommitmentNotFound.selector);
        dispatcherProxy.timeout(packetEarth, validProof);
    }

    // cannot timeout packetsfails if channel doesn't match
    function test_invalidChannel() public {
        bytes32 connectionStr = bytes32(0x636f6e6e656374696f6e2d310000000000000000000000000000000000000018); //
        // Connection-1
        IbcPacket memory packet = sentPacket;
        _storeChannelidToConnectionMapping(packet.src.channelId, connectionStr);
        sendPacket();

        IbcEndpoint memory invalidSrc = IbcEndpoint(src.portId, "channel-invalid");
        packet.src = invalidSrc;
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.channelIdNotFound.selector, packet.src.channelId));
        dispatcherProxy.timeout(packet, validProof);
    }

    // cannot timeout packets if proof from Polymer is invalid

    function test_invalidProof() public {
        sendPacket();
        vm.expectRevert(DummyLightClient.InvalidDummyNonMembershipProof.selector);
        dispatcherProxy.timeout(sentPacket, invalidProof);
    }
}

contract DappRevertTests is Base {
    RevertingBytesMars revertingBytesMars;
    PanickingMars panickingMars;
    RevertingEmptyMars revertingEmptyMars;
    RevertingStringMars revertingStringMars;
    string[] connectionHops0 = ["connection-0", "connection-3"];
    string[] connectionHops1 = ["connection-2", "connection-1"];
    ChannelEnd ch0 =
        ChannelEnd("polyibc.eth.71C95911E9a5D330f4D621842EC243EE1343292e", IbcUtils.toBytes32("channel-0"), "1.0");
    ChannelEnd ch1 =
        ChannelEnd("polyibc.eth.71C95911E9a5D330f4D621842EC243EE1343292e", IbcUtils.toBytes32("channel-1"), "1.0");

    function setUp() public virtual override {
        (dispatcherProxy, dispatcherImplementation) = TestUtilsTest.deployDispatcherProxyAndImpl(portPrefix);
        dispatcherProxy.setClientForConnection(connectionHops0[0], dummyLightClient);
        dispatcherProxy.setClientForConnection(connectionHops1[0], dummyLightClient);
        dispatcherProxy.setClientForConnection(connectionHops[0], dummyLightClient);
        revertingBytesMars = new RevertingBytesMars(dispatcherProxy);
        panickingMars = new PanickingMars(dispatcherProxy);
        revertingEmptyMars = new RevertingEmptyMars(dispatcherProxy);
        revertingStringMars = new RevertingStringMars(dispatcherProxy);
    }

    function test_ibc_channel_open_non_dapp_call() public {
        address nonDappAddr = vm.addr(1);

        emit ChannelOpenInitError(nonDappAddr, bytes("call to non-contract"));
        dispatcherProxy.channelOpenInit(ch1.version, ChannelOrder.NONE, false, connectionHops1, ch0.portId);
    }

    function test_ibc_channel_open_dapp_without_handler() public {
        Earth earth = new Earth(vm.addr(1));
        emit ChannelOpenInitError(address(earth), "");
        dispatcherProxy.channelOpenInit(ch1.version, ChannelOrder.NONE, false, connectionHops1, ch0.portId);
    }

    function test_recv_packet_callback_revert_and_panic() public {
        // this data is taken from polymerase/tests/e2e/tests/evm.events.test.ts MarsDappPair.createSentPacket()
        IbcPacket memory packet;
        packet.data = bytes("packet-1");
        packet.timeoutTimestamp = 15_566_401_733_896_437_760;
        packet.dest.channelId = ch1.channelId;
        packet.dest.portId = string(abi.encodePacked(portPrefix, IbcUtils.toHexStr(address(revertingBytesMars))));
        packet.src.portId = ch0.portId;
        packet.src.channelId = ch0.channelId;
        packet.sequence = 1;

        // Ack packet will check for an existing channel
        bytes32 connectionStr = bytes32(0x636f6e6e656374696f6e2d300000000000000000000000000000000000000018); // Connection-0
        _storeChannelidToConnectionMapping(ch1.channelId, connectionStr);

        // Test Revert Memory
        vm.expectEmit(true, true, true, true);
        emit WriteAckPacket(
            address(revertingBytesMars),
            packet.dest.channelId,
            packet.sequence,
            AckPacket(false, abi.encodeWithSelector(RevertingBytesMars.OnRecvPacketRevert.selector))
        );
        dispatcherProxy.recvPacket(packet, validProof);

        // Test Revert String
        packet.dest.portId = string(abi.encodePacked(portPrefix, IbcUtils.toHexStr(address(revertingStringMars))));
        vm.expectEmit(true, true, true, true);
        emit WriteAckPacket(
            address(revertingStringMars),
            packet.dest.channelId,
            packet.sequence,
            AckPacket(false, abi.encodeWithSignature("Error(string)", "on recv packet is reverting"))
        );
        dispatcherProxy.recvPacket(packet, validProof);

        // Test Revert empty
        packet.dest.portId = string(abi.encodePacked(portPrefix, IbcUtils.toHexStr(address(revertingEmptyMars))));
        vm.expectEmit(true, true, true, true);
        emit WriteAckPacket(address(revertingEmptyMars), packet.dest.channelId, packet.sequence, AckPacket(false, ""));
        dispatcherProxy.recvPacket(packet, validProof);

        // Test Panic
        packet.dest.portId = string(abi.encodePacked(portPrefix, IbcUtils.toHexStr(address(panickingMars))));
        vm.expectEmit(true, true, true, true);
        emit WriteAckPacket(
            address(panickingMars),
            packet.dest.channelId,
            packet.sequence,
            AckPacket(false, abi.encodeWithSignature("Panic(uint256)", uint256(1)))
        );
        dispatcherProxy.recvPacket(packet, validProof);
    }

    function test_ack_packet_dapp_revert() public {
        // plant a fake packet commitment so the ack checks go through
        // use "forge inspect --storage" to find the slot
        bytes32 slot1 = keccak256(abi.encode(address(revertingStringMars), uint32(156))); // current nested mapping
            // slot:
        bytes32 slot2 = keccak256(abi.encode(ch0.channelId, slot1));
        bytes32 slot3 = keccak256(abi.encode(uint256(1), slot2));
        vm.store(address(dispatcherProxy), slot3, bytes32(uint256(1)));

        bytes32 connectionStr = bytes32(0x636f6e6e656374696f6e2d300000000000000000000000000000000000000018); // Connection-0
        _storeChannelidToConnectionMapping(ch0.channelId, connectionStr);

        IbcPacket memory packet;
        packet.data = bytes("packet-1");
        packet.timeoutTimestamp = 15_566_401_733_896_437_760;
        packet.src.channelId = ch0.channelId;
        packet.src.portId = string(abi.encodePacked(portPrefix, IbcUtils.toHexStr(address(revertingStringMars))));
        packet.dest.portId = ch1.portId;
        packet.dest.channelId = ch1.channelId;
        packet.sequence = 1;

        // this data is taken from the write_acknowledgement event emitted by polymer
        bytes memory ack =
            bytes('{"result":"eyAiYWNjb3VudCI6ICJhY2NvdW50IiwgInJlcGx5IjogImdvdCB0aGUgbWVzc2FnZSIgfQ=="}');

        vm.expectEmit(true, true, true, true);
        emit AcknowledgementError(
            address(revertingStringMars),
            abi.encodeWithSignature("Error(string)", "acknowledgement packet is reverting")
        );
        dispatcherProxy.acknowledgement(packet, ack, validProof);
    }

    function test_ibc_channel_open_dapp_revert() public {
        vm.expectEmit(true, true, true, true);
        vm.prank(address(revertingStringMars));
        emit ChannelOpenInitError(
            address(revertingStringMars), abi.encodeWithSignature("Error(string)", "open ibc channel is reverting")
        );
        dispatcherProxy.channelOpenInit(ch1.version, ChannelOrder.NONE, false, connectionHops1, ch0.portId);
    }

    function test_ibc_channel_ack_dapp_revert() public {
        vm.expectEmit(true, true, true, true);
        ch0.portId = IbcUtils.addressToPortId(portPrefix, address(revertingStringMars));
        emit ChannelOpenAckError(
            address(revertingStringMars), abi.encodeWithSignature("Error(string)", "connect ibc channel is reverting")
        );
        dispatcherProxy.channelOpenAck(ch0, connectionHops0, ChannelOrder.NONE, false, ch1, validProof);
    }
}
