// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {ChannelOpenTestBaseSetup} from "./Dispatcher/Dispatcher.t.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import "forge-std/Test.sol";

contract FeeVaultTest is ChannelOpenTestBaseSetup {
    address notOwner = vm.addr(2);
    address owner;
    Ownable feeVaultOwnable; // FeeVault address but from Ownable interface
    uint256 feePerGreet = 600_000 * 60 gwei + 700_000 * 70 gwei;

    function setUp() public override {
        super.setUp();

        feeVaultOwnable = Ownable(address(feeVault));
        greetMarsWithFee();
    }

    // Fuzz test for always reverting if incorrect fee is sent.
    function testFuzz_sendPacket_FeesAddUpToValue(
        uint104 gasFee1,
        uint104 gasFee2,
        uint104 gasLimit1,
        uint104 gasLimit2,
        bool shouldRevert,
        uint56 fuzz
    ) public {
        gasFee1 = uint104(bound(gasFee1, 1, 2 ** 104 - 1));
        gasLimit1 = uint104(bound(gasLimit1, 1, 2 ** 104 - 1));
        fuzz = uint56(bound(fuzz, 1, 2 ** 56 - 1));
        fuzz = uint56(bound(fuzz, 1, gasFee1));

        uint256 feesToSend = uint256(gasFee1) * uint256(gasLimit1) + uint256(gasFee2) * uint256(gasLimit2);
        vm.deal(address(this), feesToSend + fuzz);

        if (shouldRevert) {
            // Fuzz the fees to make sure it reverts
            if (fuzz % 2 == 0) {
                feesToSend += uint256(fuzz);
            } else {
                feesToSend -= uint256(fuzz);
            }
            vm.expectRevert();
        }
        feeVault.depositSendPacketFee{value: feesToSend}(
            channelId, uint64(1), [uint256(gasFee1), uint256(gasFee2)], [uint256(gasLimit1), uint256(gasLimit2)]
        );
    }

    function test_withdrawAlwaysGoesToOwner() public {
        assertEq(address(feeVault).balance, feePerGreet);
        uint256 startingBalance = address(this).balance;
        feeVault.withdrawFeesToOwner();
        assertEq(address(this).balance, startingBalance + feePerGreet);

        greetMarsWithFee(); // send more fees to feeVault
        // Non Owners can call but it should sill go to the owner (i.e. this address)
        vm.prank(notOwner);
        feeVault.withdrawFeesToOwner();
        assertEq(address(feeVault).balance, 0);
        assertEq(address(this).balance, startingBalance + (feePerGreet * 2));
    }

    function greetMarsWithFee() internal {
        vm.deal(address(mars), feePerGreet);
        vm.prank(address(mars));
        vm.expectEmit(false, true, true, false, address(feeVault)); // Ignore the emitted seuqence since we don't know
            // what it will be before we call sendpacket (vm.expect emit assumes the first event will always be
            // chedcked, so to avoid checking second argument we pass false for the first argument)
        emit SendPacketFeeDeposited(
            channelId, 1, [uint256(600_000), uint256(700_000)], [uint256(60 gwei), uint256(70 gwei)]
        );
        mars.greetWithFee{value: feePerGreet}(
            "hello", channelId, maxTimeout, [uint256(600_000), uint256(700_000)], [uint256(60 gwei), uint256(70 gwei)]
        );
    }

    // This test contract needs to have a receive function to accept funds sent from the FeeVault
    // Note: our multisig should have a fallback to accept funds in general
    receive() external payable {}
}
