// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Ibc, ChannelEnd, IbcEndpoint, Height} from "../../contracts/libs/Ibc.sol";
import {IbcUtils} from "../../contracts/libs/IbcUtils.sol";
import {IBCErrors} from "../../contracts/libs/IbcErrors.sol";
import {Dispatcher} from "../../contracts/core/Dispatcher.sol";
import {IbcEventsEmitter} from "../../contracts/interfaces/IbcDispatcher.sol";
import {IbcReceiver} from "../../contracts/interfaces/IbcReceiver.sol";
import {DummyLightClient} from "../../contracts/utils/DummyLightClient.sol";
import {Base} from "../utils/Dispatcher.base.t.sol";
import "../../contracts/examples/Mars.sol";
import "../../contracts/core/OptimisticLightClient.sol";
import {LocalEnd, ChannelHandshakeSetting, Base} from "../utils/Dispatcher.base.t.sol";
import {Earth} from "../../contracts/examples/Earth.sol";

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
}

// Test Chains B receives a packet from Chain A
contract DispatcherRecvPacketTestSuite is ChannelOpenTestBaseSetup {
    IbcEndpoint src = IbcEndpoint("polyibc.bsc.58b604DB8886656695442374D8E940D814F2eDd4", "channel-99");
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

    // Can't spoof a timeed out packet
    function test_writeTimeoutPacket_cannot_Spoof_Packet() public {
        IbcPacket memory spoofedPacket = IbcPacket(src, dest, 1, payload, ZERO_HEIGHT, 1);
        spoofedPacket.timeoutTimestamp = 1; // Set really low timeout timestamp
        vm.mockCallRevert(
            address(dummyLightClient),
            abi.encodeWithSelector(
                DummyLightClient.verifyMembership.selector,
                validProof,
                Ibc.packetCommitmentProofKey(spoofedPacket),
                abi.encode(Ibc.packetCommitmentProofValue(spoofedPacket))
            ),
            abi.encodeWithSelector(IProofVerifier.InvalidProofValue.selector)
        );
        vm.expectRevert(abi.encodeWithSelector(IProofVerifier.InvalidProofValue.selector));
        dispatcherProxy.writeTimeoutPacket(spoofedPacket, validProof);
    }

    // TODO: add tests for unordered channel, wrong port, and invalid proof
}
