// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {ConsensusStateManager} from "../interfaces/ConsensusStateManager.sol";
import {L1Header, OpL2StateProof, Ics23Proof} from "../interfaces/ProofVerifier.sol";

/**
 * @title DummyConsensusStateManager
 * @dev This contract is a dummy implementation of a consensus state manager.
 *      It should only be used for testing purposes.
 *      The logic for checking if the proof length is greater than zero is naive.
 */
contract DummyConsensusStateManager is ConsensusStateManager {
    error InvalidDummyMembershipProof();
    error InvalidDummyNonMembershipProof();

    constructor() {}

    function addOpConsensusState(L1Header calldata, OpL2StateProof calldata, uint256, uint256)
        external
        pure
        override
        returns (uint256 endTime, bool ended)
    {
        return (0, false);
    }

    function getState(uint256)
        external
        pure
        override
        returns (uint256 appHash, uint256 fraudProofEndtime, bool ended)
    {
        return (0, 0, false);
    }

    function getFraudProofEndtime(uint256) external pure override returns (uint256 endTime) {
        return 0;
    }

    function verifyMembership(Ics23Proof calldata proof, bytes memory, bytes memory) external pure override {
        if (proof.height == 0) revert InvalidDummyMembershipProof();
    }

    function verifyNonMembership(Ics23Proof calldata proof, bytes memory) external pure override {
        if (proof.height == 0) revert InvalidDummyNonMembershipProof();
    }
}
