// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import './IbcVerifier.sol';
import './IbcReceiver.sol';

function isConsensusStateFilled(ConsensusState calldata state) pure returns (bool) {
    if (state.app_hash == 0 || state.valset_hash == 0 || state.time == 0 || state.height == 0) {
        return false;
    }

    return true;
}

contract DummyVerifier is ZKMintVerifier {
    function verifyConsensusState(
        ConsensusState calldata trustedState,
        ConsensusState calldata untrustedState,
        ZkProof calldata proof
    ) external pure override returns (bool) {
        require(isConsensusStateFilled(untrustedState), 'empty non-optimistic consensus state');
        return untrustedState.height >= trustedState.height;
    }

    function verifyMembership(
        uint256 appHash,
        Proof calldata proof,
        bytes calldata key,
        bytes calldata expectedValue
    ) external pure override returns (bool) {
        require(key.length > 0, 'Key cannot be empty');
        require(expectedValue.length > 0, 'Expected value cannot be empty');

        // TODO: replace with real merkle verification logic
        // For now, a dummy proof verification is implemented
        return proof.proof.length > 0;
    }

    function verifyNonMembership(
        uint256 appHash,
        Proof calldata proof,
        bytes calldata key
    ) external pure override returns (bool) {
        require(key.length > 0, 'Key cannot be empty');

        // TODO: replace with real merkle verification logic
        // For now, a dummy proof verification is implemented
        return proof.proof.length > 0;
    }
}
