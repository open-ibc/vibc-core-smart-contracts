// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ChannelOpenTestBaseSetup, PacketSenderTestBase} from "./Dispatcher.t.sol";
import {DappHandlerRevertTests} from "./Dispatcher.dappHandlerRevert.t.sol";
import {RevertingStringCloseChannelMars, Mars} from "../../contracts/examples/Mars.sol";
import {DummyLightClient} from "../../contracts/utils/DummyLightClient.sol";
import "../utils/Dispatcher.base.t.sol";
import {
    Channel, ChannelEnd, ChannelOrder, IbcPacket, ChannelState, AckPacket, Ibc
} from "../../contracts/libs/Ibc.sol";
import {IBCErrors} from "../../contracts/libs/IbcErrors.sol";
import {IbcUtils} from "../../contracts/libs/IbcUtils.sol";

contract DispatcherCloseChannelTest is PacketSenderTestBase {
    Channel defaultChannel; // Uninitialized struct to compare that structs are deleted

    function setUp() public override {
        super.setUp();
        sendPacket();
    }

    function test_closeChannelInit_success() public {
        vm.startPrank(mars.owner());
        assertNotEq0(abi.encode(dispatcherProxy.getChannel(address(mars), channelId)), abi.encode(defaultChannel));
        vm.expectEmit(true, true, true, true);
        emit ChannelCloseInit(address(mars), channelId);
        mars.triggerChannelClose(channelId);
        assertEq(abi.encode(dispatcherProxy.getChannel(address(mars), channelId)), abi.encode(defaultChannel));
        vm.stopPrank();
    }

    function test_closeChannelInit_mustOwner() public {
        Mars earth = new Mars(dispatcherProxy);
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.channelNotOwnedBySender.selector));
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
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.channelNotOwnedByPortAddress.selector));
        dispatcherProxy.channelCloseConfirm(address(mars), "channel-999", validProof);
        assertNotEq0(abi.encode(dispatcherProxy.getChannel(address(mars), channelId)), abi.encode(defaultChannel));
    }

    function test_closeChannelConfirm_invalidProof() public {
        vm.expectRevert(DummyLightClient.InvalidDummyMembershipProof.selector);
        dispatcherProxy.channelCloseConfirm(address(mars), channelId, invalidProof);
        assertNotEq0(abi.encode(dispatcherProxy.getChannel(address(mars), channelId)), abi.encode(defaultChannel));
    }

    function test_sendPacket_afterChannelCloseInit() public {
        vm.startPrank(mars.owner());
        mars.triggerChannelClose(channelId);
        sentPacket = genPacket(nextSendSeq);
        ackPacket = genAckPacket(Ibc.toStr(nextSendSeq));
        vm.expectRevert(IBCErrors.channelNotOwnedBySender.selector);
        mars.greet(payloadStr, channelId, maxTimeout);

        // Should also revert for fee enabled packets
        vm.deal(address(this), totalSendPacketFees);
        vm.expectRevert(IBCErrors.channelNotOwnedBySender.selector);
        mars.greetWithFee{value: totalSendPacketFees}(
            payloadStr, channelId, maxTimeout, sendPacketGasLimit, sendPacketGasPrice
        );
        vm.stopPrank();
    }

    function test_recvPacket_afterChannelCloseInit() public {
        // For the case where a channel close init is not relayed to polymer, we should still ensure the dispatcher
        // contract's
        // view of a closed channel is consistent both in send and recv packets after a chanCloseInit.

        // Dapp on local chain contract triggers channel close
        vm.startPrank(mars.owner());
        mars.triggerChannelClose(channelId);

        // If the above channelCloseInit wasn't relayed to peptide, the dapp on remote chain could still trigger a send
        // packet.
        // Even if the remote chain's dapp sends a packet a relayer's recv call on the local chain should revert. We
        // don't have to mock anything for the remote send packet here since this unit test is using sim client.

        uint64 packetSeq = 1;

        // Now relayer's call to relay the remote sendPacket should revert due to the local dapp having already called
        // chanCloseInit.
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.channelIdNotFound.selector, channelId));
        dispatcherProxy.recvPacket(IbcPacket(dest, src, packetSeq, payload, ZERO_HEIGHT, maxTimeout), validProof); // Note:
            // Src and dest are reversed since for this test case, packet is sent from dest to src.
    }

    function test_sendPacket_afterChannelCloseConfirm() public {
        dispatcherProxy.channelCloseConfirm(address(mars), channelId, validProof);
        sentPacket = genPacket(nextSendSeq);
        ackPacket = genAckPacket(Ibc.toStr(nextSendSeq));
        vm.expectRevert(IBCErrors.channelNotOwnedBySender.selector);
        mars.greet(payloadStr, channelId, maxTimeout);

        // Should also revert for fee enabled packets
        vm.deal(address(this), totalSendPacketFees);
        vm.expectRevert(IBCErrors.channelNotOwnedBySender.selector);
        mars.greetWithFee{value: totalSendPacketFees}(
            payloadStr, channelId, maxTimeout, sendPacketGasLimit, sendPacketGasPrice
        );
    }
}

contract DappRevertTestsCloseChannel is DappHandlerRevertTests {
    Channel defaultChannel; // Uninitialized struct to compare that structs are deleted

    RevertingStringCloseChannelMars revertingStringCloseMars;
    string portId;
    LocalEnd _local;
    ChannelEnd _remote;

    function setUp() public override {
        super.setUp();
        revertingStringCloseMars = new RevertingStringCloseChannelMars(dispatcherProxy);
        portId = IbcUtils.addressToPortId(portPrefix, address(revertingStringCloseMars));
        ChannelHandshakeSetting memory setting = ChannelHandshakeSetting(ChannelOrder.ORDERED, false, true, validProof);
        _local = LocalEnd(revertingStringCloseMars, portId, ch0.channelId, connectionHops, "1.0", "1.0");
        _remote = ChannelEnd("eth2.7E5F4552091A69125d5DfCb7b8C2659029395Bdf", "channel-2", "1.0");

        channelOpenInit(_local, _remote, setting, true);
        channelOpenTry(_local, _remote, setting, true);
        channelOpenAck(_local, _remote, setting, true);
        channelOpenConfirm(_local, _remote, setting, true);
    }

    function test_ibc_channel_close_confirm_dapp_revert() public {
        vm.expectEmit(true, true, true, true);
        emit ChannelCloseConfirmError(
            address(revertingStringCloseMars),
            abi.encodeWithSignature("Error(string)", "close ibc channel is reverting")
        );

        dispatcherProxy.channelCloseConfirm(address(revertingStringCloseMars), ch0.channelId, validProof);
        // Channels should still be deleted on dapp revert
        assertEq(
            abi.encode(defaultChannel),
            abi.encode(dispatcherProxy.getChannel(address(revertingStringCloseMars), ch0.channelId))
        );
    }
}
