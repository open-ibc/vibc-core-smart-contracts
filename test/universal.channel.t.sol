// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "../contracts/Ibc.sol";
import {Dispatcher, InitClientMsg, UpgradeClientMsg} from "../contracts/Dispatcher.sol";
import {IbcEventsEmitter} from "../contracts/IbcDispatcher.sol";
import {Escrow} from "../contracts/Escrow.sol";
import {IbcReceiver} from "../contracts/IbcReceiver.sol";
import "../contracts/IbcVerifier.sol";
import "../contracts/UniversalChannelHandler.sol";
import "../contracts/Verifier.sol";
import "../contracts/Mars.sol";
import "../contracts/OpConsensusStateManager.sol";
import "./Dispatcher.base.t.sol";
import "./VirtualChain.sol";

contract UniversalChannelTestBase is Base {
    // UniversalChannelHandler ucHandler;
    // ChannelHandshakeSetting setting;
    // LocalEnd localEnd;
    // RemoteEnd remoteEnd;

    VirtualChain eth1;
    VirtualChain eth2;

    function setUp() public virtual {
        eth1 = new VirtualChain(100, "polyibc.eth1.");
        eth2 = new VirtualChain(200, "polyibc.eth2.");
    }

    function test_channel_ok() public {
        ChannelSetting memory setting = ChannelSetting(ChannelOrder.UNORDERED, "1.0", true, validProof);
        eth1.finishHandshake(eth1.ucHandler(), eth2, eth2.ucHandler(), setting);

        bytes32 channelId1 = eth1.channelIds(address(eth1.ucHandler()), address(eth2.ucHandler()));
        bytes32 channelId2 = eth2.channelIds(address(eth2.ucHandler()), address(eth1.ucHandler()));
        assertEq(eth1.ucHandler().connectedChannels(0), channelId1);

        Channel memory channel = eth1.dispatcher().getChannel(address(eth1.ucHandler()), channelId1);
        Channel memory channelExpected =
            eth2.expectedChannel(address(eth2.ucHandler()), channelId2, eth1.getConnectionHops(), setting);
        assertEq(abi.encode(channel), abi.encode(channelExpected));

        // confirm only one channel is created
        UniversalChannelHandler ucHandler1 = UniversalChannelHandler(eth1.ucHandler());
        vm.expectRevert();
        ucHandler1.connectedChannels(1);
    }

    // function test_channel_fail_unauthorized() public {
    //     address unauthorized = deriveAddress("unauthorized");
    //     vm.deal(unauthorized, 100 ether);
    //     vm.prank(unauthorized);
    //     expectRevertNonOwner();
    //     openChannel(localEnd, remoteEnd, setting, false);

    //     openChannel(localEnd, remoteEnd, setting, true);
    //     vm.prank(unauthorized);
    //     expectRevertNonOwner();
    //     connectChannel(localEnd, remoteEnd, setting, false);
    // }

    // function test_sendPacket_ok() public {
    //     openChannel(localEnd, remoteEnd, setting, true);
    //     connectChannel(localEnd, remoteEnd, setting, true);
    //     Earth earth = new Earth(dispatcher);
    //     assertNotEq(address(earth), address(0));
    //     earth.greet(remoteEnd.portId, localEnd.channelId, "hello");
    // }
}

contract UniversalChannelPacketTest is Base {
    VirtualChain eth1;
    VirtualChain eth2;

    function setUp() public virtual {
        eth1 = new VirtualChain(100, "polyibc.eth1.");
        eth2 = new VirtualChain(200, "polyibc.eth2.");
    }

    function test_channel_handshake_ok() public {
        ChannelSetting memory setting = ChannelSetting(ChannelOrder.UNORDERED, "1.0", true, validProof);
        // msg.sender is VirtualChain contracts

        eth1.finishHandshake(eth1.ucHandler(), eth2, eth2.ucHandler(), setting);
        eth1.finishHandshake(eth1.mars(), eth2, eth2.mars(), setting);
    }
}
