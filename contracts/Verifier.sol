// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import './IbcVerifier.sol';
import './IbcReceiver.sol';
import "./Groth16Verifier.sol";


contract Verifier is ZKMintVerifier {
    function verifyConsensusState(
        ConsensusState calldata trustedState,
        ConsensusState calldata untrustedState,
        ZkProof calldata proof
    ) external view override returns (bool) {
        bool isVerified = Groth16Verifier.verifyProof(
            proof.a,
            proof.b,
            proof.c,
            [
             trustedState.valset_hash,
             trustedState.time,
             trustedState.height,
             untrustedState.app_hash,
             untrustedState.valset_hash,
             untrustedState.time,
             untrustedState.height
            ]
        );
        return isVerified;
    }

    function verifyMembership(
        OptimisticConsensusState calldata consensusState,
        Proof calldata proof,
        bytes calldata key,
        bytes calldata expectedValue
    ) external pure override returns (bool) {
        require(key.length > 0, 'Key cannot be empty');
        require(expectedValue.length > 0, 'Expected value cannot be empty');
        require(consensusState.height >= proof.proofHeight, 'Consensus state not sufficient for Merkle proof verification');

        // TODO: replace with real merkle verification logic
        // For now, a dummy proof verification is implemented
        return proof.proof.length > 0;
    }

    function verifyNonMembership(
        OptimisticConsensusState calldata consensusState,
        Proof calldata proof,
        bytes calldata key
    ) external pure override returns (bool) {
        require(key.length > 0, 'Key cannot be empty');

        // TODO: replace with real merkle verification logic
        // For now, a dummy proof verification is implemented
        return proof.proof.length > 0;
    }
}
