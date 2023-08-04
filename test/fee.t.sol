// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import 'forge-std/Test.sol';
import 'forge-std/console.sol';
import '../contracts/Ibc.sol';
import {Dispatcher, InitClientMsg, UpgradeClientMsg} from '../contracts/Dispatcher.sol';
import {IbcReceiver} from '../contracts/IbcReceiver.sol';
import '../contracts/IbcVerifier.sol';
import '../contracts/Verifier.sol';
import '../contracts/Mars.sol';
import {PacketSenderTest} from './Dispatcher.t.sol';

contract FeeTest is PacketSenderTest {
    address forwardRelayerPayee = deriveAddress('forward-payee');
    address reverseRelayerPayee = deriveAddress('reverse-payee');

    PacketFee[] packetFees;

    // pay extra packet fee after packet creation
    function test_payPacketFeeAsync() public {
        PacketFee memory accu = mergeFees(fee, fee);
        for (uint64 index = 0; index < 3; index++) {
            uint64 packetSeq = index + 1;

            mars.greet{value: Ibc.calcEscrowFee(fee)}(
                IbcDispatcher(dispatcher),
                payloadStr,
                channelId,
                maxTimeout,
                fee
            );
            dispatcher.payPacketFeeAsync{value: Ibc.calcEscrowFee(fee)}(address(mars), channelId, packetSeq, fee);

            PacketFee memory packetFee = dispatcher.getTotalPacketFees(address(mars), channelId, packetSeq);
            assertEq(keccak256(abi.encode(packetFee)), keccak256(abi.encode(accu)));
        }
    }

    // sendPacket fails if insufficient fee is paid.
    function test_insufficientFee() public {
        vm.expectRevert();
        mars.greet{value: Ibc.calcEscrowFee(fee) - 1}(
            IbcDispatcher(dispatcher),
            payloadStr,
            channelId,
            maxTimeout,
            fee
        );
    }

    modifier useFeeTestCases() {
        // initialize an array of PacketFee with different values
        packetFees.push(PacketFee(0, 0, 0));
        packetFees.push(PacketFee(1, 0, 0));
        packetFees.push(PacketFee(0, 1, 0));
        packetFees.push(PacketFee(0, 0, 1));
        packetFees.push(PacketFee(1, 1, 1));
        packetFees.push(PacketFee(1, 2, 3));
        packetFees.push(PacketFee(2, 2, 1));
        packetFees.push(PacketFee(1, 1, 2));
        packetFees.push(PacketFee(1, 2, 2));
        packetFees.push(PacketFee(3, 2, 1));

        for (uint i = 0; i < packetFees.length; i++) {
            fee = packetFees[i];
            _;
        }
    }

    // claim ack fee on receving ack
    function test_ack_fee() public useFeeTestCases {
        // save balances of forward and reverse relayer payees; assert balances changes after ack
        uint balanceForwardRelayer1 = address(forwardRelayerPayee).balance;
        uint balanceReverseRelayer1 = address(reverseRelayerPayee).balance;
        address refundPayee = address(mars);
        uint balanceRefund1 = refundPayee.balance;

        vm.startPrank(relayer);
        sendPacket();

        // Incentivized ack
        IncentivizedAckPacket memory incAck = IncentivizedAckPacket({
            success: true,
            relayer: abi.encodePacked(forwardRelayerPayee),
            data: bytes('ack-data')
        });
        vm.startPrank(reverseRelayerPayee, reverseRelayerPayee);
        dispatcher.incentivizedAcknowledgement(IbcReceiver(mars), sentPacket, incAck, validProof);

        assertEq(address(forwardRelayerPayee).balance, fee.recvFee + balanceForwardRelayer1);
        assertEq(address(reverseRelayerPayee).balance, fee.ackFee + balanceReverseRelayer1);
        assertEq(refundPayee.balance, Ibc.ackRefundAmount(fee) + balanceRefund1);
        assertEq(escrow.balance, 0);

        // cannot claim ack fee twice
        vm.expectRevert();
        dispatcher.incentivizedAcknowledgement(IbcReceiver(mars), sentPacket, incAck, validProof);
    }

    // claim timeout fee on receving timeout
    function test_timeout_fee() public {
        assertEq(address(forwardRelayerPayee).balance, 0);
        assertEq(address(reverseRelayerPayee).balance, 0);

        sendPacket();

        uint balance1 = address(mars).balance;

        vm.startPrank(reverseRelayerPayee, reverseRelayerPayee);
        dispatcher.timeout(IbcReceiver(mars), sentPacket, validProof);

        assertEq(address(reverseRelayerPayee).balance, fee.timeoutFee);
        assertEq(address(forwardRelayerPayee).balance, 0);
        assertEq(escrow.balance, 0);
        assertEq(address(mars).balance, balance1 + Ibc.timeoutRefundAmount(fee));

        // cannot claim timeout fee twice
        vm.expectRevert();
        dispatcher.timeout(IbcReceiver(mars), sentPacket, validProof);
    }
}
