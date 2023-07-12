// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import './IbcReceiver.sol';


struct ConsensusState {
    uint256 app_hash;
    uint256 valset_hash;
    uint256 time;
    uint256 height;
}

struct ZkProof {
    uint256[2]  a;
    uint256[2][2] b;
    uint256[2]  c;
}

// UpdateClientMsg is used to update an existing Polymer client on an EVM chain
struct UpdateClientMsg {
    ConsensusState consensusState;
    ZkProof zkProof;
}

interface ZKMintVerifier {
    function verifyUpdateClientMsg(
        ConsensusState calldata lastConsensusState,
        UpdateClientMsg calldata newConsensusState
    ) external view returns (bool);

    function verifyMembership(
        ConsensusState calldata consensusState,
        Proof calldata proof,
        bytes calldata key,
        bytes calldata expectedValue
    ) external pure returns (bool);

    function verifyNonMembership(
        ConsensusState calldata consensusState,
        Proof calldata proof,
        bytes calldata key
    ) external pure returns (bool);
}
