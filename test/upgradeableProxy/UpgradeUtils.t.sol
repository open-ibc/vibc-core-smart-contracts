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
import {IFeeVault} from "../../contracts/interfaces/IFeeVault.sol";
import {FeeVault} from "../../contracts/core/FeeVault.sol";
import {DummyLightClient} from "../../contracts/utils/DummyLightClient.sol";

import {IDispatcher} from "../../contracts/interfaces/IDispatcher.sol";
import {UniversalChannelHandler} from "../../contracts/core/UniversalChannelHandler.sol";
import {IUniversalChannelHandler} from "../../contracts/interfaces/IUniversalChannelHandler.sol";
import {DispatcherRc4} from "../upgradeableProxy/upgrades/DispatcherRc4.sol";
import {UniversalChannelHandlerV2} from "./upgrades/UCHV2.sol";
import {DispatcherV2Initializable} from "./upgrades/DispatcherV2Initializable.sol";
import {DispatcherV2} from "./upgrades/DispatcherV2.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

abstract contract UpgradeTestUtils {
    string[] connectionHops0 = ["dummy-connection-1", "dummy-connection-2"];
    string[] connectionHops1 = ["connection-1", "connection-2"];
    LocalEnd _localDummy;
    ChannelEnd _remoteDummy;

    function upgradeDispatcher(string memory portPrefix, IFeeVault feeVault, address dispatcherProxy)
        public
        returns (DispatcherV2Initializable newDispatcherImplementation)
    {
        // Upgrade dispatcherProxy for tests
        newDispatcherImplementation = new DispatcherV2Initializable();
        bytes memory initData = abi.encodeWithSignature("initialize(string,address)", portPrefix, feeVault);
        UUPSUpgradeable(dispatcherProxy).upgradeToAndCall(address(newDispatcherImplementation), initData);
    }

    function upgradeUch(address uchProxy) public returns (UniversalChannelHandler newUCHImplementation) {
        newUCHImplementation = new UniversalChannelHandler(); // Upgrade from v2 to v3
        UUPSUpgradeable(address(uchProxy)).upgradeTo(address(newUCHImplementation));
    }

    function deployDispatcherRC4ProxyAndImpl(string memory initPortPrefix, ILightClient initLightClient)
        public
        returns (address proxy)
    {
        DispatcherRc4 dispatcherImplementation = new DispatcherRc4();
        bytes memory initData =
            abi.encodeWithSelector(DispatcherRc4.initialize.selector, initPortPrefix, initLightClient);
        proxy = address(new ERC1967Proxy(address(dispatcherImplementation), initData));
    }

    function deployUCHV2ProxyAndImpl(address dispatcherProxy) public returns (IUniversalChannelHandler proxy) {
        UniversalChannelHandlerV2 uchImplementation = new UniversalChannelHandlerV2();
        bytes memory initData = abi.encodeWithSelector(UniversalChannelHandlerV2.initialize.selector, dispatcherProxy);
        proxy = IUniversalChannelHandler(address(new ERC1967Proxy(address(uchImplementation), initData)));
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
    uint64 timeoutTimestamp = maxTimeout;

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
