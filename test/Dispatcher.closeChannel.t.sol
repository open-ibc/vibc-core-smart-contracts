// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ChannelOpenTestBaseSetup, DappRevertTests, PacketSenderTestBase} from "./Dispatcher.t.sol";
import {RevertingStringCloseChannelMars, Mars} from "../contracts/examples/Mars.sol";
import {DummyLightClient} from "../contracts/utils/DummyLightClient.sol";
import "./Dispatcher.base.t.sol";
import {
    Channel,
    CounterParty,
    ChannelOrder,
    IbcPacket,
    ChannelState,
    AckPacket,
    IBCErrors,
    IbcUtils,
    Ibc
} from "../contracts/libs/Ibc.sol";

contract DispatcherCloseChannelTest is PacketSenderTestBase {
    Channel defaultChannel; // Uninitialized struct to compare that structs are deleted

    function setUp() public override {
        super.setUp();
        sendPacket();
    }

    function test_closeChannelInit_success() public {
        assertNotEq0(abi.encode(dispatcherProxy.getChannel(address(mars), channelId)), abi.encode(defaultChannel));
        vm.expectEmit(true, true, true, true);
        emit ChannelCloseInit(address(mars), channelId);
        mars.triggerChannelClose(channelId);
        assertEq(abi.encode(dispatcherProxy.getChannel(address(mars), channelId)), abi.encode(defaultChannel));
    }

    function test_closeChannelInit_mustOwner() public {
        Mars earth = new Mars(dispatcherProxy);
        vm.expectRevert(abi.encodeWithSignature("channelNotOwnedBySender()"));
        earth.triggerChannelClose(channelId);
        assertNotEq0(abi.encode(dispatcherProxy.getChannel(address(mars), channelId)), abi.encode(defaultChannel));
    }

    function test_closeChannelConfirm_success() public {
        assertNotEq0(abi.encode(dispatcherProxy.getChannel(address(mars), channelId)), abi.encode(defaultChannel));
        vm.expectEmit(true, true, true, true);
        emit ChannelCloseConfirm(address(mars), channelId);
        dispatcherProxy.channelCloseConfirm(address(mars), channelId, validProof);
        assertEq(abi.encode(dispatcherProxy.getChannel(address(mars), channelId)), abi.encode(defaultChannel));
    }

    function test_closeChannelConfirm_mustOwner() public {
        vm.expectRevert(abi.encodeWithSignature("channelNotOwnedByPortAddress()"));
        dispatcherProxy.channelCloseConfirm(address(mars), "channel-999", validProof);
        assertNotEq0(abi.encode(dispatcherProxy.getChannel(address(mars), channelId)), abi.encode(defaultChannel));
    }

    function test_closeChannelConfirm_invalidProof() public {
        vm.expectRevert(DummyLightClient.InvalidDummyMembershipProof.selector);
        dispatcherProxy.channelCloseConfirm(address(mars), channelId, invalidProof);
        assertNotEq0(abi.encode(dispatcherProxy.getChannel(address(mars), channelId)), abi.encode(defaultChannel));
    }

    function test_sendPacket_afterChannelCloseInit() public {
        mars.triggerChannelClose(channelId);
        sentPacket = genPacket(nextSendSeq);
        ackPacket = genAckPacket(Ibc.toStr(nextSendSeq));
        vm.expectRevert(IBCErrors.channelNotOwnedBySender.selector);
        mars.greet(payloadStr, channelId, maxTimeout);
    }

    function test_sendPacket_afterChannelCloseConfirm() public {
        dispatcherProxy.channelCloseConfirm(address(mars), channelId, validProof);
        sentPacket = genPacket(nextSendSeq);
        ackPacket = genAckPacket(Ibc.toStr(nextSendSeq));
        vm.expectRevert(IBCErrors.channelNotOwnedBySender.selector);
        mars.greet(payloadStr, channelId, maxTimeout);
    }
}

contract DappRevertTestsCloseChannel is DappRevertTests {
    RevertingStringCloseChannelMars revertingStringCloseMars;
    string portId = "eth1.7E5F4552091A69125d5DfCb7b8C2659029395Bdf";
    LocalEnd _local;
    CounterParty _remote;

    function setUp() public override {
        super.setUp();
        revertingStringCloseMars = new RevertingStringCloseChannelMars(dispatcherProxy);
        ChannelHandshakeSetting memory setting = ChannelHandshakeSetting(ChannelOrder.ORDERED, false, true, validProof);
        _local = LocalEnd(revertingStringCloseMars, portId, ch0.channelId, connectionHops, "1.0", "1.0");
        _remote = CounterParty("eth2.7E5F4552091A69125d5DfCb7b8C2659029395Bdf", "channel-2", "1.0");

        channelOpenInit(_local, _remote, setting, true);
        channelOpenTry(_local, _remote, setting, true);
        channelOpenAck(_local, _remote, setting, true);
        channelOpenConfirm(_local, _remote, setting, true);
    }

    function test_ibc_channel_close_init_dapp_revert() public {
        // We need to actually setup channelhandshake since foundry doesn't support packed slots
        vm.expectEmit(true, true, true, true);
        emit ChannelCloseInitError(
            address(revertingStringCloseMars),
            abi.encodeWithSignature("Error(string)", "close ibc channel is reverting")
        );
        revertingStringCloseMars.triggerChannelClose(ch0.channelId);
    }

    function test_ibc_channel_close_confirm_dapp_revert() public {
        vm.expectEmit(true, true, true, true);
        emit ChannelCloseConfirmError(
            address(revertingStringCloseMars),
            abi.encodeWithSignature("Error(string)", "close ibc channel is reverting")
        );

        dispatcherProxy.channelCloseConfirm(address(revertingStringCloseMars), ch0.channelId, validProof);
    }
}
