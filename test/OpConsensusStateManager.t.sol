// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import 'forge-std/Test.sol';
import '../contracts/OpConsensusStateManager.sol';
import '../contracts/DummyVerifier.sol';

contract OptimisticConsensusStateManagerTest is Test {
    OptimisticConsensusStateManager manager;
    DummyVerifier verifier;
    
    function setUp() public {
        verifier = new DummyVerifier();
        manager = new OptimisticConsensusStateManager(1, verifier);
    }

    function test_addOpConsensusState_newOpConsensusStateCreatedWithPendingStatus() public {
        (, bool ended) = manager.addOpConsensusState(1, 1);
        assertEq(false, ended);
    }

    function test_addOpConsensusState_addingAlreadyTrustedOpConsensusStateIsNoop() public {
        manager.addOpConsensusState(1, 1);

        // fast forward time.
        vm.warp(block.timestamp + 100);

        // the fraud proof window has passed.
        manager.addOpConsensusState(1, 1);

        (, , bool ended) = manager.getState(1);
        assertEq(true, ended);
    }

    function test_addOpConsensusState_addingPendingOpConsensusStateWithDifferentValuesIsError() public {
        manager.addOpConsensusState(1, 1);

        vm.expectRevert(bytes('cannot update a pending optimistic consensus state with a different appHash, please submit fraud proof instead'));
        manager.addOpConsensusState(1, 2);
    }

    function test_addOpConsensusState_addingSameOpConsensusStateIsNoop() public {
        manager.addOpConsensusState(1, 1);

        (, uint256 originalFraudProofEndTime, ) = manager.getState(1);

        vm.warp(block.timestamp + 1);

        // adding the same appHash later doesn't update the fraud
        // proof end time.
        manager.addOpConsensusState(1, 1);
        (, uint256 newFraudProofEndTime, ) = manager.getState(1);
        assertEq(originalFraudProofEndTime, newFraudProofEndTime);
    }

    function test_zero_proof_window() public {
        verifier = new DummyVerifier();
        manager = new OptimisticConsensusStateManager(0, verifier);
        manager.addOpConsensusState(1, 1);
        (, uint256 endTime, bool ended) = manager.getState(1);
        assertEq(true, ended);
     }

    function test_getState_nonExist() public {
        (uint256 appHash, , bool ended) = manager.getState(1);
        assertEq(0, appHash);
        assertEq(false, ended);
    }
}
