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

    // claim ack fee on receving ack
    function test_ack_fee() public {
        assertEq(address(forwardRelayerPayee).balance, 0);
        assertEq(address(reverseRelayerPayee).balance, 0);

        sendPacket();
        uint balance1 = address(mars).balance;

        // Incentivized ack
        IncentivizedAckPacket memory incAck = IncentivizedAckPacket({
            success: true,
            relayer: abi.encodePacked(forwardRelayerPayee),
            data: bytes('ack-data')
        });
        vm.startPrank(reverseRelayerPayee, reverseRelayerPayee);
        dispatcher.incentivizedAcknowledgement(IbcReceiver(mars), sentPacket, incAck, validProof);

        assertEq(address(forwardRelayerPayee).balance, fee.recvFee);
        assertEq(address(reverseRelayerPayee).balance, fee.ackFee);
        assertEq(escrow.balance, 0);
        assertEq(address(mars).balance, Ibc.ackRefundAmount(fee) + balance1);

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
