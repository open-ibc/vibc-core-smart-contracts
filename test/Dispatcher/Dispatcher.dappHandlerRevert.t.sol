// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "../utils/Dispatcher.base.t.sol";
import "../../contracts/examples/Mars.sol";
import "../../contracts/examples/Earth.sol";
import "../../contracts/libs/Ibc.sol";
import {IbcUtils} from "../../contracts/libs/IbcUtils.sol";
import {IbcReceiver} from "../../contracts/interfaces/IbcReceiver.sol";

contract DappHandlerRevertTests is Base {
    RevertingBytesMars revertingBytesMars;
    PanickingMars panickingMars;
    RevertingEmptyMars revertingEmptyMars;
    RevertingStringMars revertingStringMars;
    string[] connectionHops0 = ["connection-0", "connection-3"];
    string[] connectionHops1 = ["connection-2", "connection-1"];
    ChannelEnd ch0 =
        ChannelEnd("polyibc.eth1.71C95911E9a5D330f4D621842EC243EE1343292e", IbcUtils.toBytes32("channel-0"), "1.0");
    ChannelEnd ch1 =
        ChannelEnd("polyibc.eth2.71C95911E9a5D330f4D621842EC243EE1343292e", IbcUtils.toBytes32("channel-1"), "1.0");

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
