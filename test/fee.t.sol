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

    // pay extra packet fee after packet creation
    function test_payPacketFeeAsync() public useFeeTestCases {
        // merge fees to calculate expected packet fee
        PacketFee memory extraFee = PacketFee(1, 2, 3);
        PacketFee memory expectedFee = mergeFees(fee, extraFee);

        for (uint64 index = 0; index < 3; index++) {
            // create packet and pay escrow fee
            sendPacket();

            // pay packet fee asynchronously
            dispatcher.payPacketFeeAsync{value: Ibc.calcEscrowFee(extraFee)}(
                address(mars),
                channelId,
                sentPacket.sequence,
                extraFee
            );

            // check that the total packet fee is correct
            PacketFee memory actualFee = dispatcher.getTotalPacketFees(address(mars), channelId, sentPacket.sequence);
            assertEq(keccak256(abi.encode(actualFee)), keccak256(abi.encode(expectedFee)), 'Packet fee is incorrect');
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
    function test_timeout_fee() public useFeeTestCases {
        uint balanceReverseRelayer1 = address(reverseRelayerPayee).balance;
        address refundPayee = address(mars);
        uint balanceRefund1 = refundPayee.balance;

        vm.startPrank(relayer);
        sendPacket();

        vm.startPrank(reverseRelayerPayee, reverseRelayerPayee);
        dispatcher.timeout(IbcReceiver(mars), sentPacket, validProof);

        assertEq(address(reverseRelayerPayee).balance, fee.timeoutFee + balanceReverseRelayer1);
        // no forward relayer is invovled in timeout case
        assertEq(address(forwardRelayerPayee).balance, 0);
        assertEq(refundPayee.balance, Ibc.timeoutRefundAmount(fee) + balanceRefund1);
        assertEq(escrow.balance, 0);

        // cannot claim timeout fee twice
        vm.expectRevert();
        dispatcher.timeout(IbcReceiver(mars), sentPacket, validProof);
    }
}
