// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import './ConsensusStateManager.sol';

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

    function verifyMembership(Proof calldata, bytes memory, bytes memory, string memory) external view {}

    function verifyNonMembership(Proof calldata, bytes memory, string memory) external view {}
}
