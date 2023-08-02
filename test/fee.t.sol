// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import 'forge-std/Test.sol';
import '../contracts/Ibc.sol';
import {Dispatcher, InitClientMsg, UpgradeClientMsg} from '../contracts/Dispatcher.sol';
import {IbcReceiver} from '../contracts/IbcReceiver.sol';
import '../contracts/IbcVerifier.sol';
import '../contracts/Verifier.sol';
import '../contracts/Mars.sol';
import {PacketSenderTest} from './Dispatcher.t.sol';

contract FeeTest is PacketSenderTest {
    address payee = deriveAddress('payee');

    // pay extra packet fee after packet creation
    function test_payPacketFeeAsync() public {
        PacketFee memory accu = mergeFees(fee, fee);
        for (uint64 index = 0; index < 3; index++) {
            uint64 packetSeq = index + 1;

            mars.greet{value: calcFee(fee)}(IbcDispatcher(dispatcher), payloadStr, channelId, maxTimeout, fee);
            dispatcher.payPacketFeeAsync{value: calcFee(fee)}(address(mars), channelId, packetSeq, fee);

            PacketFee memory packetFee = dispatcher.getTotalPacketFees(address(mars), channelId, packetSeq);
            assertEq(keccak256(abi.encode(packetFee)), keccak256(abi.encode(accu)));
        }
    }

    // sendPacket fails if insufficient fee is paid.
    function test_insufficientFee() public {
        vm.expectRevert();
        mars.greet{value: calcFee(fee) - 1}(IbcDispatcher(dispatcher), payloadStr, channelId, maxTimeout, fee);
    }

    // claim ack fee on receving ack
    function test_ack_fee() public {
        assertEq(address(payee).balance, 0);
        sendPacket();

        // No fee distributed for non-incentivized ack
        dispatcher.acknowledgement(IbcReceiver(mars), sentPacket, ackPacket, validProof);
        assertEq(address(payee).balance, 0);

        // Incentivized ack
    }
}
