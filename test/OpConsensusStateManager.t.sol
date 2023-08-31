// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import 'forge-std/Test.sol';
import '../contracts/OpConsensusStateManager.sol';

contract OptimisticConsensusStateManagerCaller {
    using OptimisticConsensusStateManager for OptimisticConsensusStateManager.DataStore;
    OptimisticConsensusStateManager.DataStore private opConsensusStatesStore;

    constructor() {
        opConsensusStatesStore.fraudProofWindowSeconds = 1;
    }

    function addOpConsensusState(uint256 height, uint256 appHash)
        external returns (uint256 fraudProofEndTime, bool ended){
        return opConsensusStatesStore.addOpConsensusState(height, appHash);
    }

    function getState(uint256 height) external returns (uint256 appHash, uint256 fraudProofEndTime, bool ended) {
        return opConsensusStatesStore.getState(height);
    }
}

contract OptimisticConsensusStateManagerTest is Test {
    OptimisticConsensusStateManagerCaller caller;
    
    function setUp() public {
        caller = new OptimisticConsensusStateManagerCaller();
    }

    function test_addOpConsensusState_newOpConsensusStateCreatedWithPendingStatus() public {
        (, bool ended) = caller.addOpConsensusState(1, 1);
        assertEq(false, ended);
    }

    function test_addOpConsensusState_addingAlreadyTrustedOpConsensusStateIsNoop() public {
        caller.addOpConsensusState(1, 1);

        // fast forward time.
        vm.warp(block.timestamp + 100);

        // the fraud proof window has passed.
        caller.addOpConsensusState(1, 1);

        (, , bool ended) = caller.getState(1);
        assertEq(true, ended);
    }

    function test_addOpConsensusState_addingPendingOpConsensusStateWithDifferentValuesIsError() public {
        caller.addOpConsensusState(1, 1);

        vm.expectRevert(bytes('cannot update a pending optimistic consensus state with a different appHash, please submit fraud proof instead'));
        caller.addOpConsensusState(1, 2);
    }

    function test_addOpConsensusState_addingSameOpConsensusStateIsNoop() public {
        caller.addOpConsensusState(1, 1);

        (, uint256 originalFraudProofEndTime, ) = caller.getState(1);

        vm.warp(block.timestamp + 1);

        // adding the same appHash later doesn't update the fraud
        // proof end time.
        caller.addOpConsensusState(1, 1);
        (, uint256 newFraudProofEndTime, ) = caller.getState(1);
        assertEq(originalFraudProofEndTime, newFraudProofEndTime);
    }

    function test_getState_nonExist() public {
        (uint256 appHash, , bool ended) = caller.getState(1);
        assertEq(0, appHash);
        assertEq(false, ended);
    }
}
