// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./ConsensusStateManager.sol";
/**
 * @title DummyConsensusStateManager
 * @dev This contract is a dummy implementation of a consensus state manager.
 *      It should only be used for testing purposes.
 *      The logic for checking if the proof length is greater than zero is naive.
 */

contract DummyConsensusStateManager is ConsensusStateManager {
    constructor() {}

    function addOpConsensusState(uint256, uint256) external pure returns (uint256, bool) {
        return (0, false);
    }

    function getState(uint256) external pure returns (uint256, uint256, bool) {
        return (0, 0, false);
    }

    function getFraudProofEndtime(uint256) external pure returns (uint256) {
        return 0;
    }

    function verifyMembership(Proof calldata proof, bytes memory key, bytes memory expectedValue, string memory message)
        external
        view
    {
        require(proof.proof.length > 0, message);
    }

    function verifyNonMembership(Proof calldata proof, bytes memory key, string memory message) external view {
        require(proof.proof.length > 0, message);
    }
}
