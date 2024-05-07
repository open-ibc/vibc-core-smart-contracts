// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/console2.sol";
import {DispatcherUpdateClientTestSuite} from "../Dispatcher/Dispatcher.client.t.sol";
import {DispatcherIbcWithRealProofsSuite} from "../Dispatcher/Dispatcher.proof.t.sol";
import {Mars} from "../../contracts/examples/Mars.sol";
import {Earth} from "../../contracts/examples/Earth.sol";
import "../../contracts/core/OptimisticLightClient.sol";
import {ChannelHandshakeTestSuite, ChannelHandshakeTest, ChannelHandshakeUtils} from "../Dispatcher/Dispatcher.t.sol";
import {LocalEnd} from "../utils/Dispatcher.base.t.sol";
import {Base, ChannelHandshakeSetting} from "../utils/Dispatcher.base.t.sol";
import {
    ChannelEnd,
    ChannelOrder,
    ChannelState,
    IbcEndpoint,
    IbcPacket,
    AckPacket,
    Ibc,
    Height,
    UniversalPacket
} from "../../contracts/libs/Ibc.sol";
import {IbcUtils} from "../../contracts/libs/IbcUtils.sol";
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

// Tests to ensure that the upgrade between rc4 -> preaudit fixes doesn't break the state of the contract
contract DispatcherRC4UpgradeTest is ChannelHandShakeUpgradeUtil, UpgradeTestUtils {
    Mars receivingMars;
    Mars dummyMars;
    IDispatcher oldDummyDispatcherProxy;
    IUniversalChannelHandler uch;
    LocalEnd _localUch;
    Earth earth;

    function setUp() public override {
        // In Rc4 version, there can only be one dispatcher per light client so we deploy multiple clients
        // Deploy dummy old dispathcer
        oldDummyDispatcherProxy = deployDispatcherRC4ProxyAndImpl(portPrefix, dummyLightClient);

        // Deploy op old dispatcher
        DummyLightClient dummyLightClient2 = new DummyLightClient(); // dummyLightClient2 models the op light client in
            // prod - it will be the light client that is chosen for the upgrade (and the oldDummyDispatcherProxy will
            // be deprecated)
        dispatcherProxy = deployDispatcherRC4ProxyAndImpl(portPrefix, dummyLightClient2);
        uch = deployUCHV2ProxyAndImpl(address(dispatcherProxy));
        earth = new Earth(address(uch));

        // Set up dispatcher with non-trivial state
        mars = new Mars(dispatcherProxy);
        receivingMars = new Mars(dispatcherProxy);
        string memory sendingPortId = IbcUtils.addressToPortId(portPrefix, address(mars));
        string memory uchSendingPortId = IbcUtils.addressToPortId(portPrefix, address(uch));
        string memory receivingPortId = IbcUtils.addressToPortId(portPrefix, address(receivingMars));
        _local = LocalEnd(mars, sendingPortId, "channel-1", connectionHops1, "1.0", "1.0");
        _localUch = LocalEnd(uch, uchSendingPortId, "uch-channel", connectionHops1, "1.0", "1.0");
        _remote = ChannelEnd(receivingPortId, "channel-2", "1.0");

        // Should now be able to able to open a connection without proofs on the upgraded dispatcherproxy now
        _localDummy = LocalEnd(mars, sendingPortId, "dummy-channel-1", connectionHops0, "1.0", "1.0");
        _remoteDummy = ChannelEnd(receivingPortId, "dummy-channel-2", "1.0");

        // Add state to test if impacted by upgrade
        doProofChannelHandshake(_local, _remote);
        sendPacket(_local.channelId);

        // Do channel handshake via uch
        doProofChannelHandshake(_localUch, _remote);
        earth.greet(address(mars), _localUch.channelId, bytes("hello mars"), UINT64_MAX);

        // Upgrade dispatcherProxy and uch for tests
        upgradeDispatcher(portPrefix, address(dispatcherProxy));
        upgradeUch(address(uch));
        dispatcherProxy.setClientForConnection(connectionHops0[0], dummyLightClient2);
        dispatcherProxy.setClientForConnection(connectionHops1[0], dummyLightClient2);
    }

    function test_SentPacketState_Conserved_RC4_Upgrade() public {
        // Check packet state in sendPacketCommitment()[]
        uint64 nextSequenceSendValue = uint64(
            uint256(vm.load(address(dispatcherProxy), findNextSequenceSendSlot(address(mars), _local.channelId)))
        );

        assertEq(4, nextSequenceSendValue);

        // Validate packets from previous send
        assert(vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(mars), _local.channelId, 1)) > 0);
        assert(vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(mars), _local.channelId, 2)) > 0);
        assert(vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(mars), _local.channelId, 3)) > 0);
        assert(
            vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(uch), _localUch.channelId, 1)) > 0
        );

        // Test sending packet with the updated contract
        sendOnePacket(_local.channelId, 4, mars);
        assert(vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(mars), _local.channelId, 4)) > 0);
        uint64 nextSequenceSendAfterSending = uint64(
            uint256(vm.load(address(dispatcherProxy), findNextSequenceSendSlot(address(mars), _local.channelId)))
        );
        assertEq(5, nextSequenceSendAfterSending);
        earth.greet(address(receivingMars), _localUch.channelId, bytes("hello from upgrade mars!"), UINT64_MAX);
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
        uint64 nextSequenceRecvValue =
            uint64(uint256(vm.load(address(dispatcherProxy), findNextSequenceRecv(address(mars), _local.channelId))));
        uint64 nextSequenceAckValue =
            uint64(uint256(vm.load(address(dispatcherProxy), findNextSequenceAck(address(mars), _local.channelId))));

        assertEq(1, nextSequenceRecvValue);
        assertEq(1, nextSequenceAckValue);

        // Should be able to do the channel handshake and send a packet from the dummy client
        doChannelHandshake(_localDummy, _remoteDummy);
        sendOnePacket(_localDummy.channelId, 1, mars);
        assert(vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(mars), _local.channelId, 1)) > 0);
        uint64 nextSequenceSendAfterSendingDummy = uint64(
            uint256(vm.load(address(dispatcherProxy), findNextSequenceSendSlot(address(mars), _localDummy.channelId)))
        );
        assertEq(2, nextSequenceSendAfterSendingDummy);

        // Should be able to open a channel with another contract after upgrade through optimistic light client
        Mars mars2 = new Mars(dispatcherProxy);

        string memory portId2 = IbcUtils.addressToPortId(portPrefix, address(mars2));
        LocalEnd memory _local2 = LocalEnd(mars2, portId2, "channel-1", connectionHops1, "1.0", "1.0");
        ChannelHandshakeSetting memory setting = ChannelHandshakeSetting(ChannelOrder.ORDERED, false, false, validProof);

        channelOpenTry(_local2, _remote, setting, false);

        // Another dapp should be able to initialize another channel and send a packet from the new channel
        doProofChannelHandshake(_local2, _remote);
        sendOnePacket(_local2.channelId, 1, mars2);
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
contract DispatcherRC4MidwayUpgradeTest is ChannelHandShakeUpgradeUtil, UpgradeTestUtils {
    Mars receivingMars;
    LocalEnd _re;
    ChannelEnd _le;
    string sendingPortId;
    string receivingPortId;
    DummyLightClient dummyLightClient2;
    IUniversalChannelHandler uch;
    LocalEnd _localUch;
    Earth earth;

    function setUp() public override {
        ChannelHandshakeSetting memory setting = ChannelHandshakeSetting(ChannelOrder.ORDERED, false, true, validProof);

        DispatcherRc4 oldOpDispatcherImplementation = new DispatcherRc4();
        bytes memory initData = abi.encodeWithSelector(DispatcherRc4.initialize.selector, portPrefix, dummyLightClient);
        dispatcherProxy = IDispatcher(address(new ERC1967Proxy(address(oldOpDispatcherImplementation), initData)));

        dummyLightClient2 = new DummyLightClient();

        mars = new Mars(dispatcherProxy);
        uch = deployUCHV2ProxyAndImpl(address(dispatcherProxy));
        earth = new Earth(address(uch));
        receivingMars = new Mars(dispatcherProxy);

        sendingPortId = IbcUtils.addressToPortId(portPrefix, address(mars));
        string memory uchSendingPortId = IbcUtils.addressToPortId(portPrefix, address(uch));
        receivingPortId = IbcUtils.addressToPortId(portPrefix, address(receivingMars));
        // LocalEnd version of _remote
        _re = LocalEnd(receivingMars, receivingPortId, "channel-2", connectionHops1, "1.0", "1.0");
        // ChannelEnd version of _local
        _le = ChannelEnd(sendingPortId, "channel-1", "1.0");
        _local = LocalEnd(mars, sendingPortId, "channel-1", connectionHops1, "1.0", "1.0");
        _localUch = LocalEnd(uch, uchSendingPortId, "uch-channel", connectionHops1, "1.0", "1.0");
        _remote = ChannelEnd(receivingPortId, "channel-2", "1.0");

        // Do only half a handshake to test upgrading between handshakes
        channelOpenInit(_local, _remote, setting, true);

        channelOpenTry(_re, _le, setting, true);

        // Do channel handshake via uch
        doProofChannelHandshake(_localUch, _remote);
        earth.greet(address(mars), _localUch.channelId, bytes("hello mars"), UINT64_MAX);
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

        sendOnePacket(_local.channelId, 1, mars);

        // Ensure packet is sent correctly
        assert(vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(mars), _local.channelId, 1)) > 0);

        // Do upgrade before finishing packet handshake
        upgradeDispatcher(portPrefix, address(dispatcherProxy));
        upgradeUch(address(uch));
        dispatcherProxy.setClientForConnection(connectionHops1[0], dummyLightClient2);

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
        assert(vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(mars), _local.channelId, 1)) == 0);

        uint64 nextSequenceAckValue =
            uint64(uint256(vm.load(address(dispatcherProxy), findNextSequenceAck(address(mars), _local.channelId))));

        uint64 nextSequenceSendValue = uint64(
            uint256(vm.load(address(dispatcherProxy), findNextSequenceSendSlot(address(mars), _local.channelId)))
        );

        assertEq(2, nextSequenceAckValue);
        assertEq(2, nextSequenceSendValue);

        // Now recv uch packet
        bytes memory appData = abi.encodePacked("hello using mw stack");
        UniversalPacket memory ucPacket =
            UniversalPacket(IbcUtils.toBytes32(address(mars)), uch.MW_ID(), IbcUtils.toBytes32(address(earth)), appData);
        bytes memory packetData = IbcUtils.toUniversalPacketBytes(ucPacket);
        IbcPacket memory uchPacket = IbcPacket(
            IbcEndpoint(sendingPortId, _local.channelId),
            IbcEndpoint(_localUch.portId, _localUch.channelId),
            1,
            packetData,
            Height(0, 0),
            maxTimeout
        );

        AckPacket memory earthAck = earth.generateAckPacket(0x0, address(mars), appData);
        vm.expectEmit(true, true, true, true);
        emit WriteAckPacket(address(uch), _localUch.channelId, 1, earthAck);
        dispatcherProxy.recvPacket(uchPacket, validProof);

        (bytes32 actualChannelId, UniversalPacket memory storedUcPacket) = earth.recvedPackets(0);
        assertEq(actualChannelId, _localUch.channelId);
        assertEq(storedUcPacket.appData, ucPacket.appData);
    }
}
