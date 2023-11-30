// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "../contracts/Ibc.sol";
import {Dispatcher, InitClientMsg, UpgradeClientMsg} from "../contracts/Dispatcher.sol";
import {IbcEventsEmitter} from "../contracts/IbcDispatcher.sol";
import {Escrow} from "../contracts/Escrow.sol";
import {IbcReceiver} from "../contracts/IbcReceiver.sol";
import "../contracts/IbcVerifier.sol";
import "../contracts/Verifier.sol";
import "../contracts/Mars.sol";
import "../contracts/OpConsensusStateManager.sol";
import "./Dispatcher.base.t.sol";

contract ClientTestBase is Base {
    function setUp() public virtual {
        dispatcher = new Dispatcher(verifier, escrow, portPrefix, opConsensusStateManager);
    }
}

contract DispatcherCreateClientTest is ClientTestBase {
    function test_success() public {
        dispatcher.createClient(initClientMsg);
    }

    function test_mustByOwner() public {
        vm.prank(deriveAddress("non-onwer"));
        expectRevertNonOwner();
        dispatcher.createClient(initClientMsg);
    }

    function test_onlyOnce() public {
        dispatcher.createClient(initClientMsg);
        vm.expectRevert(abi.encodeWithSignature("clientAlreadyCreated()"));
        dispatcher.createClient(initClientMsg);
    }
}

contract DispatcherUpdateClientTest is ClientTestBase {
    function setUp() public override {
        super.setUp();
        dispatcher.createClient(initClientMsg);
    }

    function test_updateConsensusState_success() public {
        dispatcher.updateClientWithConsensusState(trustedState, proof);
    }

    function test_updateConsensusState_invalid() public {
        vm.expectRevert(abi.encodeWithSignature("consensusStateVerificationFailed()"));
        dispatcher.updateClientWithConsensusState(untrustedState, proof);

        vm.expectRevert(abi.encodeWithSignature("consensusStateVerificationFailed()"));
        ConsensusState memory invalidConsensusState;
        dispatcher.updateClientWithConsensusState(invalidConsensusState, proof);
    }
}

contract DispatcherUpgradeClientTest is ClientTestBase {
    function setUp() public override {
        super.setUp();
        dispatcher.createClient(initClientMsg);
        dispatcher.updateClientWithConsensusState(trustedState, proof);
    }

    function test_success() public {
        dispatcher.upgradeClient(UpgradeClientMsg(bytes("upgradeOptimisticConsensusState"), trustedState));
    }

    function test_ownerOnly() public {
        vm.prank(deriveAddress("non-onwer"));
        expectRevertNonOwner();
        dispatcher.upgradeClient(UpgradeClientMsg(bytes("upgradeOptimisticConsensusState"), trustedState));
    }
}
