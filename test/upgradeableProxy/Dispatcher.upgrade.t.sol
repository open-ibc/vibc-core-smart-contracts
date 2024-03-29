// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/console2.sol";
import {DispatcherUpdateClientTestSuite} from "../Dispatcher.client.t.sol";
import {DispatcherIbcWithRealProofsSuite} from "../Dispatcher.proof.t.sol";
import {Mars} from "../../contracts/examples/Mars.sol";
import "../../contracts/core/OpLightClient.sol";
import {ChannelHandshakeTestSuite, ChannelHandshakeTest, ChannelHandshakeUtils} from "../Dispatcher.t.sol";
import {LocalEnd} from "../Dispatcher.base.t.sol";
import {Base, ChannelHandshakeSetting} from "../Dispatcher.base.t.sol";
import {
    ChannelEnd,
    ChannelOrder,
    IbcEndpoint,
    IbcPacket,
    AckPacket,
    Ibc,
    IbcUtils,
    Height
} from "../../contracts/libs/Ibc.sol";
import {IbcReceiver} from "../../contracts/interfaces/IbcReceiver.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
import {OptimisticLightClient} from "../../contracts/core/OpLightClient.sol";
import {ProofVerifier} from "../../contracts/core/OpProofVerifier.sol";
import {DummyLightClient} from "../../contracts/utils/DummyLightClient.sol";

import {IDispatcher} from "../../contracts/interfaces/IDispatcher.sol";
import {DispatcherV2Initializable} from "./upgrades/DispatcherV2Initializable.sol";
import {DispatcherV2} from "./upgrades/DispatcherV2.sol";

abstract contract UpgradeTestUtils {
    function upgradeDispatcher(string memory portPrefix, address dispatcherProxy)
        public
        returns (DispatcherV2Initializable newDispatcherImplementation)
    {
        // Upgrade dispatcherProxy for tests
        newDispatcherImplementation = new DispatcherV2Initializable();
        bytes memory initData = abi.encodeWithSignature("initialize(string)", portPrefix);
        UUPSUpgradeable(dispatcherProxy).upgradeToAndCall(address(newDispatcherImplementation), initData);
    }
}

contract ChannelHandShakeUpgradeUtil is ChannelHandshakeUtils {
    uint32 nextSequenceSendSlot = 153;
    uint32 sendPacketCommitmentSlot = 156;
    uint32 nextSequenceAckSlot = 155;
    uint32 nextSequenceRecvSlot = 154;
    IbcPacket[3] packets;
    string payload = "msgPayload";
    bytes packet = abi.encodePacked(payload);
    uint64 timeoutTimestamp = 1000;

    function doChannelHandshake() public {
        // same setup as that run in test_connectChannel_ok
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

    function sendOnePacket(bytes32 channelId, uint64 packetSeq) public {
        vm.expectEmit(true, true, true, true);
        emit SendPacket(address(mars), channelId, packet, packetSeq, timeoutTimestamp);
        mars.greet(payload, channelId, timeoutTimestamp);
    }

    function sendPacket(bytes32 channelId) public {
        for (uint64 index = 0; index < 3; index++) {
            uint64 packetSeq = index + 1;
            sendOnePacket(channelId, packetSeq);
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

    function findPacketCommitmentSlot(address portAddress, bytes32 channelId, uint64 sequence)
        public
        view
        returns (bytes32 slot)
    {
        bytes32 slot1 = keccak256(abi.encode(portAddress, sendPacketCommitmentSlot));
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
        (dispatcherProxy, dispatcherImplementation) = deployDispatcherProxyAndImpl(portPrefix);
        dispatcherProxy.setClientForConnection(connectionHops[0], dummyLightClient);
        mars = new Mars(dispatcherProxy);
        portId = IbcUtils.addressToPortId(portPrefix, address(mars));
        _local = LocalEnd(mars, portId, "channel-1", connectionHops, "1.0", "1.0");
        _remote = ChannelEnd("eth2.7E5F4552091A69125d5DfCb7b8C2659029395Bdf", "channel-2", "1.0");

        // Add state to test if impacted by upgrade
        doChannelHandshake();
        sendPacket(_local.channelId);

        // Upgrade dispatcherProxy for tests
        upgradeDispatcher("adfsafsa", address(dispatcherProxy));
    }

    function test_SentPacketState_Conserved() public {
        // Check  packet state  in sendPacketCommitment()[]
        uint64 nextSequenceSendValue = uint64(
            uint256(vm.load(address(dispatcherProxy), findNextSequenceSendSlot(address(mars), _local.channelId)))
        );

        assertEq(4, nextSequenceSendValue);

        // packets
        assert(vm.load(address(dispatcherProxy), findPacketCommitmentSlot(address(mars), _local.channelId, 1)) > 0);
        assert(vm.load(address(dispatcherProxy), findPacketCommitmentSlot(address(mars), _local.channelId, 2)) > 0);
        assert(vm.load(address(dispatcherProxy), findPacketCommitmentSlot(address(mars), _local.channelId, 3)) > 0);

        // Test sending packet with the updated contract
        sendOnePacket(_local.channelId, 4);
        assert(vm.load(address(dispatcherProxy), findPacketCommitmentSlot(address(mars), _local.channelId, 4)) > 0);
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
