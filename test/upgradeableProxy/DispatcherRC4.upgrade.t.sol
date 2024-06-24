// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/console2.sol";
import {DispatcherUpdateClientTestSuite} from "../Dispatcher/Dispatcher.client.t.sol";
import {DispatcherIbcWithRealProofsSuite} from "../Dispatcher/Dispatcher.proof.t.sol";
import {Mars} from "../../contracts/examples/Mars.sol";
import "../../contracts/core/OptimisticLightClient.sol";
import {ChannelHandshakeTestSuite, ChannelHandshakeTest, ChannelHandshakeUtils} from "../Dispatcher/Dispatcher.t.sol";
import {LocalEnd} from "../utils/Dispatcher.base.t.sol";
import {Base, ChannelHandshakeSetting} from "../utils/Dispatcher.base.t.sol";
import {
    ChannelEnd, ChannelOrder, ChannelState, IbcEndpoint, IbcPacket, Ibc, Height
} from "../../contracts/libs/Ibc.sol";
import {IbcReceiver} from "../../contracts/interfaces/IbcReceiver.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
import {OptimisticLightClient} from "../../contracts/core/OptimisticLightClient.sol";
import {IProofVerifier} from "../../contracts/core/OptimisticProofVerifier.sol";
import {DummyLightClient} from "../../contracts/utils/DummyLightClient.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {IDispatcher} from "../../contracts/interfaces/IDispatcher.sol";

import {ChannelHandShakeUpgradeUtil, UpgradeTestUtils} from "./Dispatcher.upgrade.t.sol";
import {IDispatcherRc4, DispatcherRc4} from "./upgrades/DispatcherRc4.sol";
import {IUniversalChannelHandler} from "../../contracts/interfaces/IUniversalChannelHandler.sol";
import {VirtualChain} from "../VirtualChain.sol";
import {GeneralMiddleware} from "../../contracts/base/GeneralMiddleware.sol";
import {Mars as MarsRc4, IbcDispatcher as IbcDispatcherRc4} from "./upgrades/MarsRc4.sol";
import {
    Earth as EarthRc4,
    UniversalPacket as UniversalPacketRc4,
    AckPacket as AckPacketRc4,
    IbcUtils
} from "./upgrades/EarthRc4.sol";

abstract contract DispatcherRC4TestUtils is ChannelHandShakeUpgradeUtil {
    function sendOnePacketRc4(bytes32 channelId, uint64 packetSeq, MarsRc4 sender) public {
        vm.expectEmit(true, true, true, true);
        emit SendPacket(address(sender), channelId, abi.encodePacked("msgPayload"), packetSeq, timeoutTimestamp);
        sender.greet(payload, channelId, timeoutTimestamp);
    }

    function sendPacketRc4(bytes32 channelId, MarsRc4 mars) public {
        for (uint64 index = 0; index < 3; index++) {
            uint64 packetSeq = index + 1;
            sendOnePacketRc4(channelId, packetSeq, mars);
            IbcEndpoint memory dest = IbcEndpoint("polyibc.bsc.9876543210", "channel-99");
            string memory marsPort = string(abi.encodePacked(portPrefix, getHexBytes(address(mars))));
            IbcEndpoint memory src = IbcEndpoint(marsPort, channelId);
            packets[index] = IbcPacket(src, dest, packetSeq, bytes(payload), ZERO_HEIGHT, maxTimeout);
        }
    }
}

// Tests to ensure that the upgrade between rc4 -> preaudit fixes doesn't break the state of the contract
contract DispatcherRC4UpgradeTest is DispatcherRC4TestUtils, UpgradeTestUtils {
    MarsRc4 sendingMars;
    MarsRc4 receivingMars;
    MarsRc4 dummyMars;
    IbcDispatcherRc4 oldDummyDispatcherProxy;
    IbcDispatcherRc4 oldDispatcherInterface;
    IUniversalChannelHandler uch;
    LocalEnd _localUch;
    EarthRc4 earth;

    function setUp() public override {
        // In Rc4 version, there can only be one dispatcher per light client so we deploy multiple clients
        // Deploy dummy old dispathcer
        oldDummyDispatcherProxy = IbcDispatcherRc4(address(deployDispatcherRC4ProxyAndImpl(portPrefix, dummyLightClient))); // we have to manually cast here because solidity is confused by having interfaces coming from seperate files

        // Deploy op old dispatcher
        DummyLightClient dummyLightClient2 = new DummyLightClient(); // dummyLightClient2 models the op light client in
            // prod - it will be the light client that is chosen for the upgrade (and the oldDummyDispatcherProxy will
            // be deprecated)
        oldDispatcherInterface = IbcDispatcherRc4(address(deployDispatcherRC4ProxyAndImpl(portPrefix, dummyLightClient2)));
        dispatcherProxy = IDispatcher(address(oldDispatcherInterface));
        uch = deployUCHV2ProxyAndImpl(address(dispatcherProxy));
        earth = new EarthRc4(address(uch));

        // Set up dispatcher with non-trivial state
        sendingMars = new MarsRc4(oldDispatcherInterface);
        receivingMars = new MarsRc4(oldDispatcherInterface);
        string memory sendingPortId = IbcUtils.addressToPortId(portPrefix, address(sendingMars));
        string memory uchSendingPortId = IbcUtils.addressToPortId(portPrefix, address(uch));
        string memory receivingPortId = IbcUtils.addressToPortId(portPrefix, address(receivingMars));
        _local = LocalEnd(IbcReceiver(address(sendingMars)), sendingPortId, "channel-1", connectionHops1, "1.0", "1.0");
        _localUch = LocalEnd(uch, uchSendingPortId, "uch-channel", connectionHops1, "1.0", "1.0");
        _remote = ChannelEnd(receivingPortId, "channel-2", "1.0");

        // Should now be able to able to open a connection without proofs on the upgraded dispatcherproxy now
        _localDummy =
            LocalEnd(IbcReceiver(address(sendingMars)), sendingPortId, "dummy-channel-1", connectionHops0, "1.0", "1.0");
        _remoteDummy = ChannelEnd(receivingPortId, "dummy-channel-2", "1.0");

        // Add state to test if impacted by upgrade
        doProofChannelHandshake(_local, _remote);
        sendPacketRc4(_local.channelId, sendingMars);

        // Do channel handshake via uch
        doProofChannelHandshake(_localUch, _remote);
        earth.greet(address(sendingMars), _localUch.channelId, bytes("hello sendingMars"), UINT64_MAX);

        // Upgrade dispatcherProxy and uch for tests
        upgradeDispatcher(portPrefix, address(dispatcherProxy));
        upgradeUch(address(uch));
        dispatcherProxy.setClientForConnection(connectionHops0[0], dummyLightClient2);
        dispatcherProxy.setClientForConnection(connectionHops1[0], dummyLightClient2);
    }

    function test_SentPacketState_Conserved_RC4_Upgrade() public {
        // Check packet state in sendPacketCommitment()[]
        uint64 nextSequenceSendValue = uint64(
            uint256(vm.load(address(dispatcherProxy), findNextSequenceSendSlot(address(sendingMars), _local.channelId)))
        );

        assertEq(4, nextSequenceSendValue);

        // Validate packets from previous send
        assert(
            vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(sendingMars), _local.channelId, 1))
                > 0
        );
        assert(
            vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(sendingMars), _local.channelId, 2))
                > 0
        );
        assert(
            vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(sendingMars), _local.channelId, 3))
                > 0
        );
        assert(
            vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(uch), _localUch.channelId, 1)) > 0
        );

        // Test sending packet with the updated contract
        sendOnePacketRc4(_local.channelId, 4, sendingMars);
        assert(
            vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(sendingMars), _local.channelId, 4))
                > 0
        );
        uint64 nextSequenceSendAfterSending = uint64(
            uint256(vm.load(address(dispatcherProxy), findNextSequenceSendSlot(address(sendingMars), _local.channelId)))
        );
        assertEq(5, nextSequenceSendAfterSending);
        earth.greet(address(receivingMars), _localUch.channelId, bytes("hello from upgrade sendingMars!"), UINT64_MAX);
        assert(
            vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(uch), _localUch.channelId, 2)) > 0
        );
        uint64 nextUCHSequenceSendAfterSending = uint64(
            uint256(vm.load(address(dispatcherProxy), findNextSequenceSendSlot(address(uch), _localUch.channelId)))
        );
        assertEq(3, nextUCHSequenceSendAfterSending);
    }

    function test_OpenChannelState_Conserved_RC4Upgrade() public {
        // State should be conserved after upgrade
        uint64 nextSequenceRecvValue = uint64(
            uint256(vm.load(address(dispatcherProxy), findNextSequenceRecv(address(sendingMars), _local.channelId)))
        );
        uint64 nextSequenceAckValue = uint64(
            uint256(vm.load(address(dispatcherProxy), findNextSequenceAck(address(sendingMars), _local.channelId)))
        );

        assertEq(1, nextSequenceRecvValue);
        assertEq(1, nextSequenceAckValue);

        // Should be able to do the channel handshake and send a packet from the dummy client
        doChannelHandshake(_localDummy, _remoteDummy);
        sendOnePacketRc4(_localDummy.channelId, 1, sendingMars);
        assert(
            vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(sendingMars), _local.channelId, 1))
                > 0
        );
        uint64 nextSequenceSendAfterSendingDummy = uint64(
            uint256(
                vm.load(address(dispatcherProxy), findNextSequenceSendSlot(address(sendingMars), _localDummy.channelId))
            )
        );
        assertEq(2, nextSequenceSendAfterSendingDummy);

        // Should be able to open a channel with another contract after upgrade through optimistic light client
        MarsRc4 mars2 = new MarsRc4(oldDispatcherInterface);

        string memory portId2 = IbcUtils.addressToPortId(portPrefix, address(mars2));
        LocalEnd memory _local2 =
            LocalEnd(IbcReceiver(address(mars2)), portId2, "channel-1", connectionHops1, "1.0", "1.0");
        ChannelHandshakeSetting memory setting = ChannelHandshakeSetting(ChannelOrder.ORDERED, false, false, validProof);

        channelOpenTry(_local2, _remote, setting, false);

        // Another dapp should be able to initialize another channel and send a packet from the new channel
        doProofChannelHandshake(_local2, _remote);
        sendOnePacketRc4(_local2.channelId, 1, mars2);
        assert(
            vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(mars2), _local2.channelId, 1)) > 0
        );
        uint64 nextSequenceSendAfterSending = uint64(
            uint256(vm.load(address(dispatcherProxy), findNextSequenceSendSlot(address(mars2), _local2.channelId)))
        );
        assertEq(2, nextSequenceSendAfterSending);
    }
}

// Contract to test that state upgrades within mappings are preserved .
contract DispatcherRC4MidwayUpgradeTest is DispatcherRC4TestUtils, UpgradeTestUtils {
    IbcDispatcherRc4 oldDispatcherInterface;
    MarsRc4 sendingMars;
    MarsRc4 receivingMars;
    LocalEnd _re;
    ChannelEnd _le;
    string sendingPortId;
    string receivingPortId;
    DummyLightClient dummyLightClient2;
    IUniversalChannelHandler uch;
    LocalEnd _localUch;
    EarthRc4 earth;

    event WriteAckPacket(
        address indexed writerPortAddress, bytes32 indexed writerChannelId, uint64 sequence, AckPacketRc4 ackPacket
    );

    function setUp() public override {
        ChannelHandshakeSetting memory setting = ChannelHandshakeSetting(ChannelOrder.ORDERED, false, true, validProof);

        DispatcherRc4 oldOpDispatcherImplementation = new DispatcherRc4();
        bytes memory initData = abi.encodeWithSelector(DispatcherRc4.initialize.selector, portPrefix, dummyLightClient);
        oldDispatcherInterface =
            IbcDispatcherRc4(address(new ERC1967Proxy(address(oldOpDispatcherImplementation), initData)));
        dispatcherProxy = IDispatcher(address(oldDispatcherInterface));

        dummyLightClient2 = new DummyLightClient();

        sendingMars = new MarsRc4(oldDispatcherInterface);
        uch = deployUCHV2ProxyAndImpl(address(dispatcherProxy));
        earth = new EarthRc4(address(uch));
        receivingMars = new MarsRc4(oldDispatcherInterface);

        sendingPortId = IbcUtils.addressToPortId(portPrefix, address(sendingMars));
        string memory uchSendingPortId = IbcUtils.addressToPortId(portPrefix, address(uch));
        receivingPortId = IbcUtils.addressToPortId(portPrefix, address(receivingMars));
        // LocalEnd version of _remote
        _re = LocalEnd(IbcReceiver(address(receivingMars)), receivingPortId, "channel-2", connectionHops1, "1.0", "1.0");
        // ChannelEnd version of _local
        _le = ChannelEnd(sendingPortId, "channel-1", "1.0");
        _local = LocalEnd(IbcReceiver(address(sendingMars)), sendingPortId, "channel-1", connectionHops1, "1.0", "1.0");
        _localUch = LocalEnd(uch, uchSendingPortId, "uch-channel", connectionHops1, "1.0", "1.0");
        _remote = ChannelEnd(receivingPortId, "channel-2", "1.0");

        // Do only half a handshake to test upgrading between handshakes
        channelOpenInit(_local, _remote, setting, true);

        channelOpenTry(_re, _le, setting, true);

        // Do channel handshake via uch
        doProofChannelHandshake(_localUch, _remote);
        earth.greet(address(sendingMars), _localUch.channelId, bytes("hello sendingMars"), UINT64_MAX);
    }

    // Test that channel handshake can be finished even if done during an upgrade
    function test_UpgradeBetween_ChannelOpen() public {
        upgradeDispatcher(portPrefix, address(dispatcherProxy));
        dispatcherProxy.setClientForConnection(connectionHops1[0], dummyLightClient2);
        ChannelHandshakeSetting memory setting = ChannelHandshakeSetting(ChannelOrder.ORDERED, false, true, validProof);
        channelOpenAck(_local, _remote, setting, true);
        channelOpenConfirm(_re, _le, setting, true);
    }

    // Test that packet sending can be finished even if done during an upgade
    function test_UpgradeBetween_SentPacketState() public {
        // Finish up channel handshake so that we can send packet
        ChannelHandshakeSetting memory setting = ChannelHandshakeSetting(ChannelOrder.ORDERED, false, true, validProof);
        channelOpenAck(_local, _remote, setting, true);
        channelOpenConfirm(_re, _le, setting, true);
        // Send packet before upgrade

        sendOnePacketRc4(_local.channelId, 1, sendingMars);

        // Ensure packet is sent correctly
        assert(
            vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(sendingMars), _local.channelId, 1))
                > 0
        );

        // Do upgrade before finishing packet handshake
        upgradeDispatcher(portPrefix, address(dispatcherProxy));
        upgradeUch(address(uch));
        dispatcherProxy.setClientForConnection(connectionHops1[0], dummyLightClient2);
        // earth.authorizeChannel(_localUch.channelId);

        // Now recv and ack packet
        IbcEndpoint memory src = IbcEndpoint(sendingPortId, _local.channelId);
        IbcEndpoint memory dest = IbcEndpoint(receivingPortId, _remote.channelId);
        IbcPacket memory pkt = IbcPacket(src, dest, 1, bytes("payload"), ZERO_HEIGHT, maxTimeout);
        dispatcherProxy.recvPacket(pkt, validProof);
        assert(
            vm.load(address(dispatcherProxy), findAckPacketCommitmentSlot(address(receivingMars), _remote.channelId, 1))
                > 0
        );

        dispatcherProxy.acknowledgement(pkt, genAckPacket("1"), validProof);
        // Send packet commitment should be deleted after sending
        assert(
            vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(sendingMars), _local.channelId, 1))
                == 0
        );

        uint64 nextSequenceAckValue = uint64(
            uint256(vm.load(address(dispatcherProxy), findNextSequenceAck(address(sendingMars), _local.channelId)))
        );

        uint64 nextSequenceSendValue = uint64(
            uint256(vm.load(address(dispatcherProxy), findNextSequenceSendSlot(address(sendingMars), _local.channelId)))
        );

        assertEq(2, nextSequenceAckValue);
        assertEq(2, nextSequenceSendValue);

        // Now recv uch packet
        bytes memory appData = abi.encodePacked("hello using mw stack");
        UniversalPacketRc4 memory ucPacket = UniversalPacketRc4(
            IbcUtils.toBytes32(address(sendingMars)), uch.MW_ID(), IbcUtils.toBytes32(address(earth)), appData
        );
        bytes memory packetData = IbcUtils.toUniversalPacketBytes(ucPacket);
        IbcPacket memory uchPacket = IbcPacket(
            IbcEndpoint(sendingPortId, _local.channelId),
            IbcEndpoint(_localUch.portId, _localUch.channelId),
            1,
            packetData,
            Height(0, 0),
            maxTimeout
        );

        AckPacketRc4 memory earthAck = earth.generateAckPacket(0x0, address(sendingMars), appData);

        vm.expectEmit(true, true, true, true);
        emit WriteAckPacket(address(uch), _localUch.channelId, 1, earthAck);
        dispatcherProxy.recvPacket(uchPacket, validProof);

        (bytes32 actualChannelId, UniversalPacketRc4 memory storedUcPacket) = earth.recvedPackets(0);
        assertEq(actualChannelId, _localUch.channelId);
        assertEq(storedUcPacket.appData, ucPacket.appData);
    }
}
