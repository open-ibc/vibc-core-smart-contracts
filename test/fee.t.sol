// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import 'forge-std/Test.sol';
import '../contracts/Ibc.sol';
import {Dispatcher, InitClientMsg, UpgradeClientMsg} from '../contracts/Dispatcher.sol';
import {IbcReceiver} from '../contracts/IbcReceiver.sol';
import '../contracts/IbcVerifier.sol';
import '../contracts/Verifier.sol';
import '../contracts/Mars.sol';
import {ChannelOpenTestBase} from './Dispatcher.t.sol';

contract FeeTest is ChannelOpenTestBase {
    // default params
    string payload = 'msgPayload';
    uint64 timeoutTimestamp = 1000;
    PacketFee fee = PacketFee(1 ether, 2 ether, 3 ether);

    address payee = deriveAddress('payee');

    // pay extra packet fee after packet creation
    function test_payPacketFeeAsync() public {
        PacketFee memory accu = mergeFees(fee, fee);
        for (uint64 index = 0; index < 3; index++) {
            uint64 packetSeq = index + 1;

            mars.greet{value: calcFee(fee)}(IbcDispatcher(dispatcher), payload, channelId, timeoutTimestamp, fee);
            dispatcher.payPacketFeeAsync{value: calcFee(fee)}(address(mars), channelId, packetSeq, fee);

            PacketFee memory packetFee = dispatcher.getTotalPacketFees(address(mars), channelId, packetSeq);
            assertEq(keccak256(abi.encode(packetFee)), keccak256(abi.encode(accu)));
        }
    }

    // sendPacket fails if insufficient fee is paid.
    function test_insufficientFee() public {
        vm.expectRevert();
        mars.greet{value: calcFee(fee) - 1}(IbcDispatcher(dispatcher), payload, channelId, timeoutTimestamp, fee);
    }

    // claim ack fee on receving ack
    function test_ack_fee() public {
        mars.greet{value: calcFee(fee)}(IbcDispatcher(dispatcher), payload, channelId, timeoutTimestamp, fee);
    }
}
