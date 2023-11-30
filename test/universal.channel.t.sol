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

contract UniversalChannelTestBase is Base {
    UniversalChannelHandler ucHandler;
    ChannelHandshakeSetting setting;
    LocalEnd localEnd;
    RemoteEnd remoteEnd;

    function setUp() public virtual {
        dispatcher = new Dispatcher(verifier, escrow, portPrefix, dummyConsStateManager);
        Escrow myEscrow = Escrow(escrow);
        myEscrow.setDispatcher(address(dispatcher));

        ucHandler = new UniversalChannelHandler(dispatcher);
        dispatcher.setUniversalChannelHandler(ucHandler);

        setting = ChannelHandshakeSetting(ChannelOrder.UNORDERED, true, true, validProof);
        remoteEnd = RemoteEnd("eth2.7E5F4552091A69125d5DfCb7b8C2659029395Bdf", "chanel-81", "1.0");
        localEnd = LocalEnd(ucHandler, "chanel-80", connectionHops, "1.0", "1.0");
    }

    function test_channel_ok() public {
        openChannel(localEnd, remoteEnd, setting, true);
        connectChannel(localEnd, remoteEnd, setting, true);
        assertEq(ucHandler.connectedChannels(0), localEnd.channelId);

        Channel memory channel = dispatcher.getChannel(address(localEnd.receiver), localEnd.channelId);
        Channel memory channelExpected = Channel(
            localEnd.versionExpected,
            setting.ordering,
            setting.feeEnabled,
            connectionHops,
            remoteEnd.portId,
            remoteEnd.channelId
        );
        assertEq(abi.encode(channel), abi.encode(channelExpected));
        // confirm only one channel is created
        vm.expectRevert();
        ucHandler.connectedChannels(1);
    }

    function test_channel_fail_unauthorized() public {
        address unauthorized = deriveAddress("unauthorized");
        vm.deal(unauthorized, 100 ether);
        vm.prank(unauthorized);
        expectRevertNonOwner();
        openChannel(localEnd, remoteEnd, setting, false);

        openChannel(localEnd, remoteEnd, setting, true);
        vm.prank(unauthorized);
        expectRevertNonOwner();
        connectChannel(localEnd, remoteEnd, setting, false);
    }

    function test_sendPacket_ok() public {
        openChannel(localEnd, remoteEnd, setting, true);
        connectChannel(localEnd, remoteEnd, setting, true);
        Earth earth = new Earth(dispatcher);
        assertNotEq(address(earth), address(0));
        earth.greet(remoteEnd.portId, localEnd.channelId, "hello");
    }
}

contract Earth is IbcReceiverBase {
    constructor(IbcDispatcher _dispatcher) IbcReceiverBase(_dispatcher) {}

    function greet(string calldata portId, bytes32 channelId, string calldata message) external payable {
        PacketFee memory fee;
        dispatcher.sendPacketOverUniversalChannel{value: Ibc.calcEscrowFee(fee)}(
            portId, channelId, bytes(message), 1 days, fee
        );
    }
}
