// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import './IbcVerifier.sol';
import './IbcReceiver.sol';


function isUpdateClientMsgFilled(UpdateClientMsg memory message) pure returns (bool) {
    // check that each field of the ConsensusState struct is not zero
    if (message.consensusState.app_hash == 0 ||
    message.consensusState.valset_hash == 0 ||
    message.consensusState.time == 0 ||
        message.consensusState.height == 0) {
        return false;
    }

    // check that each field of the ZkProof struct is not zero
    for (uint256 i = 0; i < 2; i++) {
        if (message.zkProof.a[i] == 0) return false;
        if (message.zkProof.c[i] == 0) return false;
        for (uint256 j = 0; j < 2; j++) {
            if (message.zkProof.b[i][j] == 0) return false;
        }
    }

    // if none of the fields were zero, the struct is considered filled
    return true;
}


contract DummyVerifier is ZKMintVerifier {
    function verifyUpdateClientMsg(
        ConsensusState calldata trustedState,
        UpdateClientMsg calldata updateClientMsg
    ) external pure override returns (bool) {
        ConsensusState memory untrustedState = updateClientMsg.consensusState;

        require(isUpdateClientMsgFilled(updateClientMsg), 'Invalid update client message');
        return untrustedState.height >= trustedState.height;
    }

    function verifyMembership(
        ConsensusState calldata consensusState,
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
        ConsensusState calldata consensusState,
        Proof calldata proof,
        bytes calldata key
    ) external pure override returns (bool) {
        require(key.length > 0, 'Key cannot be empty');

        // TODO: replace with real merkle verification logic
        // For now, a dummy proof verification is implemented
        return proof.proof.length > 0;
    }
}
