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
    ChannelEnd,
    ChannelOrder,
    ChannelState,
    IbcEndpoint,
    IbcPacket,
    AckPacket,
    Ibc,
    Height
} from "../../contracts/libs/Ibc.sol";
import {IbcUtils} from "../../contracts/libs/IbcUtils.sol";
import {IbcReceiver, IbcChannelReceiver} from "../../contracts/interfaces/IbcReceiver.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
import {OptimisticLightClient} from "../../contracts/core/OptimisticLightClient.sol";
import {IProofVerifier} from "../../contracts/core/OptimisticProofVerifier.sol";
import {DummyLightClient} from "../../contracts/utils/DummyLightClient.sol";

import {IDispatcher} from "../../contracts/interfaces/IDispatcher.sol";
import {DispatcherRc4} from "./upgrades/DispatcherRc4.sol";
import {DispatcherV2Initializable} from "./upgrades/DispatcherV2Initializable.sol";
import {DispatcherV2} from "./upgrades/DispatcherV2.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

abstract contract UpgradeTestUtils {
    string[] connectionHops0 = ["dummy-connection-1", "dummy-connection-2"];
    string[] connectionHops1 = ["connection-1", "connection-2"];
    LocalEnd _localDummy;
    ChannelEnd _remoteDummy;

    function upgradeDispatcher(string memory portPrefix, address dispatcherProxy)
        public
        returns (DispatcherV2Initializable newDispatcherImplementation)
    {
        // Upgrade dispatcherProxy for tests
        newDispatcherImplementation = new DispatcherV2Initializable();
        bytes memory initData = abi.encodeWithSignature("initialize(string)", portPrefix);
        UUPSUpgradeable(dispatcherProxy).upgradeToAndCall(address(newDispatcherImplementation), initData);
    }

    function deployDispatcherRC4ProxyAndImpl(string memory initPortPrefix, ILightClient initLightClient)
        public
        returns (IDispatcher proxy)
    {
        DispatcherRc4 dispatcherImplementation = new DispatcherRc4();
        bytes memory initData =
            abi.encodeWithSelector(DispatcherRc4.initialize.selector, initPortPrefix, initLightClient);
        proxy = IDispatcher(address(new ERC1967Proxy(address(dispatcherImplementation), initData)));
    }
}

contract ChannelHandShakeUpgradeUtil is ChannelHandshakeUtils {
    uint32 nextSequenceSendSlot = 153;
    uint32 sendPacketCommitmentSlot = 156;
    uint32 ackPacketCommitmentSlot = 158;
    uint32 nextSequenceAckSlot = 155;
    uint32 nextSequenceRecvSlot = 154;
    IbcPacket[3] packets;
    string payload = "msgPayload";
    bytes packet = abi.encodePacked(payload);
    uint64 timeoutTimestamp = 1000;

    // Conduct 4-step channel handshake between localChannelEnd and remoteChannelEnd end
    // Have to pass in receivingDapp because it's hard to parse out the addresss from the portId
    function doChannelHandshake(LocalEnd memory localEnd, ChannelEnd memory remoteEnd) public {
        // same setup as that run in test_connectChannel_ok
        ChannelHandshakeSetting[8] memory settings = createSettings2(true);

        string[2] memory versions = ["1.0", "2.0"];
        for (uint256 i = 0; i < settings.length; i++) {
            for (uint256 j = 0; j < versions.length; j++) {
                localEnd.versionCall = versions[j];
                localEnd.versionExpected = versions[j];
                remoteEnd.version = versions[j];
                channelOpenInit(localEnd, remoteEnd, settings[i], true);
                channelOpenTry(localEnd, remoteEnd, settings[i], true);
                channelOpenAck(localEnd, remoteEnd, settings[i], true);
                channelOpenConfirm(localEnd, remoteEnd, settings[i], true);
            }
        }
    }

    function doProofChannelHandshake(LocalEnd memory localEnd, ChannelEnd memory remoteEnd) public {
        ChannelHandshakeSetting memory setting = ChannelHandshakeSetting(ChannelOrder.ORDERED, false, true, validProof);

        channelOpenInit(localEnd, remoteEnd, setting, true);
        channelOpenTry(localEnd, remoteEnd, setting, true);
        channelOpenAck(localEnd, remoteEnd, setting, true);
        channelOpenConfirm(localEnd, remoteEnd, setting, true);
    }

    function sendOnePacket(bytes32 channelId, uint64 packetSeq, Mars sender) public {
        vm.expectEmit(true, true, true, true);
        emit SendPacket(address(sender), channelId, packet, packetSeq, timeoutTimestamp);
        sender.greet(payload, channelId, timeoutTimestamp);
    }

    function sendPacket(bytes32 channelId) public {
        for (uint64 index = 0; index < 3; index++) {
            uint64 packetSeq = index + 1;
            sendOnePacket(channelId, packetSeq, mars);
            IbcEndpoint memory dest = IbcEndpoint("polyibc.bsc.9876543210", "channel-99");
            string memory marsPort = string(abi.encodePacked(portPrefix, getHexBytes(address(mars))));
            IbcEndpoint memory src = IbcEndpoint(marsPort, channelId);
            packets[index] = IbcPacket(src, dest, packetSeq, bytes(payload), ZERO_HEIGHT, maxTimeout);
        }
    }

    function acknowledgePacket() public {}

    function findNextSequenceSendSlot(address portAddress, bytes32 channelId) public view returns (bytes32 slot) {
        bytes32 slot1 = keccak256(abi.encode(portAddress, nextSequenceSendSlot));
        slot = keccak256(abi.encode(channelId, slot1));
    }

    function findSendPacketCommitmentSlot(address portAddress, bytes32 channelId, uint64 sequence)
        public
        view
        returns (bytes32 slot)
    {
        bytes32 slot1 = keccak256(abi.encode(portAddress, sendPacketCommitmentSlot));
        bytes32 slot2 = keccak256(abi.encode(channelId, slot1));
        slot = keccak256(abi.encode(sequence, slot2));
    }

    function findAckPacketCommitmentSlot(address portAddress, bytes32 channelId, uint64 sequence)
        public
        view
        returns (bytes32 slot)
    {
        bytes32 slot1 = keccak256(abi.encode(portAddress, ackPacketCommitmentSlot));
        bytes32 slot2 = keccak256(abi.encode(channelId, slot1));
        slot = keccak256(abi.encode(sequence, slot2));
    }

    function findNextSequenceAck(address portAddress, bytes32 channelId) public view returns (bytes32 slot) {
        bytes32 slot1 = keccak256(abi.encode(portAddress, nextSequenceAckSlot));
        slot = keccak256(abi.encode(channelId, slot1));
    }

    function findNextSequenceRecv(address portAddress, bytes32 channelId) public view returns (bytes32 slot) {
        bytes32 slot1 = keccak256(abi.encode(portAddress, nextSequenceRecvSlot));
        slot = keccak256(abi.encode(channelId, slot1));
    }
}

contract DispatcherUpgradeTest is ChannelHandShakeUpgradeUtil, UpgradeTestUtils {
    function setUp() public override {
        address targetMarsAddress = 0x71C95911E9a5D330f4D621842EC243EE1343292e;
        (dispatcherProxy, dispatcherImplementation) = deployDispatcherProxyAndImpl(portPrefix);
        deployCodeTo("contracts/examples/Mars.sol:Mars", abi.encode(address(dispatcherProxy)), targetMarsAddress);
        dispatcherProxy.setClientForConnection(connectionHops[0], dummyLightClient);
        mars = new Mars(dispatcherProxy);
        string memory sendingPortId = IbcUtils.addressToPortId(portPrefix, address(mars));
        string memory receivingPortId = IbcUtils.addressToPortId(portPrefix, targetMarsAddress);
        _local = LocalEnd(mars, sendingPortId, "channel-1", connectionHops, "1.0", "1.0");
        _remote = ChannelEnd(receivingPortId, "channel-2", "1.0");

        // Add state to test if impacted by upgrade
        doChannelHandshake(_local, _remote);
        sendPacket(_local.channelId);

        // Upgrade dispatcherProxy for tests
        upgradeDispatcher("adfsafsa", address(dispatcherProxy));
    }

    function test_SentPacketState_Conserved() public {
        uint64 nextSequenceSendValue = uint64(
            uint256(vm.load(address(dispatcherProxy), findNextSequenceSendSlot(address(mars), _local.channelId)))
        );

        assertEq(4, nextSequenceSendValue);

        // packets
        assert(vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(mars), _local.channelId, 1)) > 0);
        assert(vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(mars), _local.channelId, 2)) > 0);
        assert(vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(mars), _local.channelId, 3)) > 0);

        // Test sending packet with the updated contract
        sendOnePacket(_local.channelId, 4, mars);
        assert(vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(mars), _local.channelId, 4)) > 0);
        uint64 nextSequenceSendAfterSending = uint64(
            uint256(vm.load(address(dispatcherProxy), findNextSequenceSendSlot(address(mars), _local.channelId)))
        );
        assertEq(5, nextSequenceSendAfterSending);
    }

    function test_OpenChannelState_Conserved() public {
        // State should be conserved after upgrade
        uint64 nextSequenceRecvValue =
            uint64(uint256(vm.load(address(dispatcherProxy), findNextSequenceRecv(address(mars), _local.channelId))));
        uint64 nextSequenceAckValue =
            uint64(uint256(vm.load(address(dispatcherProxy), findNextSequenceAck(address(mars), _local.channelId))));

        assertEq(1, nextSequenceRecvValue);
        assertEq(1, nextSequenceAckValue);
    }
}
