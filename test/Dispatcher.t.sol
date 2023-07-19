// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import 'forge-std/Test.sol';
import {Dispatcher, InitClientMsg} from '../contracts/Dispatcher.sol';
import '../contracts/IbcVerifier.sol';
import '../contracts/Verifier.sol';

contract Base is Test {
    ConsensusState trustedState =
        ConsensusState(
            80990,
            590199110038530808131927832294665177527506259518072095333098842116767351199,
            7000040,
            1000
        );
    ConsensusState untrustedState =
        ConsensusState(
            10934,
            7064503680087416120706887577693908749828198688716609274705703517077803898371,
            7002040,
            1020
        );
    InitClientMsg initClientMsg = InitClientMsg(bytes('Polymer'), trustedState);

    ZKMintVerifier verifier;
    address payable escrow = payable(vm.addr(uint256(keccak256(abi.encode('escrow')))));
}

contract DispatcherBaseTest is Test, Base {
    function setUp() public {
        verifier = new Verifier();
    }

    function test_createClient_success() public {
        Dispatcher dispatcher = new Dispatcher(verifier, escrow, 'polyibc.eth.');
        dispatcher.createClient(initClientMsg);
    }

    function test_createClient_mustByOwner() public {
        Dispatcher dispatcher = new Dispatcher(verifier, escrow, 'polyibc.eth.');
        vm.prank(vm.addr(0x1));
        vm.expectRevert('Ownable: caller is not the owner');
        dispatcher.createClient(initClientMsg);
    }

    function test_createClient_onlyOnce() public {
        Dispatcher dispatcher = new Dispatcher(verifier, escrow, 'polyibc.eth.');
        dispatcher.createClient(initClientMsg);
        vm.expectRevert('Client already created');
        dispatcher.createClient(initClientMsg);
    }
}

contract DispatcherTest is Test, Base {
    function setUp() public {}
}
