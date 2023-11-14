// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import './IbcReceiver.sol';

// ConsensusState contains the complete information that can be used
// to verify the next update for consensus state.
//
// TODO: add more fields, such as public keys of signers, their voting
// powers, etc.
struct ConsensusState {
    uint256 app_hash;
    uint256 valset_hash;
    uint256 time;
    uint256 height;
}

struct ZkProof {
    uint256[2] a;
    uint256[2][2] b;
    uint256[2] c;
}

interface ZKMintVerifier {
    function verifyConsensusState(
        ConsensusState calldata lastConsensusState,
        ConsensusState calldata newConsensusState,
        ZkProof calldata proof
    ) external view returns (bool);

    function verifyMembership(
        uint256 appHash,
        Proof calldata proof,
        bytes calldata key,
        bytes calldata expectedValue
    ) external pure returns (bool);

    function verifyNonMembership(
        uint256 appHash,
        Proof calldata proof,
        bytes calldata key
    ) external pure returns (bool);
}
