// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import './Ibc.sol';

interface ConsensusStateManager {
    /**
     * addOpConsensusState adds an appHash to internal store and
     * returns the fraud proof end time, and a bool flag indicating if
     * the fraud proof window has passed according to the block's time stamp.
     */
    function addOpConsensusState(uint256 height, uint256 appHash) external returns (uint256, bool);

    /**
     *
     */
    function getState(uint256 height) external view returns (uint256, uint256, bool);

    /**
     *
     */
    function getFraudProofEndtime(uint256 height) external view returns (uint256);

    /**
     * verifyMembership checks if the current state
     * can be used to perform the membership test and if so, it uses
     * the verifier to perform membership check.
     */
    function verifyMembership(
        Proof calldata proof,
        bytes memory key,
        bytes memory expectedValue,
        string memory message
    ) external view;

    /**
     *
     */
    function verifyNonMembership(Proof calldata proof, bytes memory key, string memory message) external view;
}
