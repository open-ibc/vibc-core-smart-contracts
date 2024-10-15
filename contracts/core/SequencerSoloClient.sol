// SPDX-License-Identifier: Apache-2.0
/*
 * Copyright 2024, Polymer Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

pragma solidity 0.8.15;

import {Ics23Proof} from "../interfaces/IProofVerifier.sol";
import {ISignatureVerifier} from "../interfaces/ISignatureVerifier.sol";
import {ILightClient, LightClientType} from "../interfaces/ILightClient.sol";
import {L1Block} from "optimism/L2/L1Block.sol";

/**
 * @title SequencerSoloClient
 * @author Polymer Labs
 * @dev This light client implementation verifies a single signature by the polymer p2p-signer, and updates the light
 * client if valid signatures are provided.
 */
contract SequencerSoloClient is ILightClient {
    LightClientType public constant LIGHT_CLIENT_TYPE = LightClientType.SequencerLightClient; // Stored as a constant
        // for cheap on-chain use

    ISignatureVerifier public immutable verifier;
    L1Block public immutable l1BlockProvider;

    uint8 private _LightClientType = uint8(LIGHT_CLIENT_TYPE); // Also redundantly stored as a private mutable type in
        // case it needs to be accessed in any proofs

    // consensusStates maps from the height to the appHash, and is modified during state updates if it is a valid state
    // update.
    mapping(uint256 => uint256) public consensusStates;

    error InvalidL1Origin();
    error CannotUpdateClientWithDifferentAppHash();
    error AppHashHasNotPassedFraudProofWindow();
    error NonMembershipProofsNotYetImplemented();
    error NoConsensusStateAtHeight(uint256 height);

    constructor(ISignatureVerifier verifier_, L1Block _l1BlockProvider) {
        verifier = verifier_;
        l1BlockProvider = _l1BlockProvider;
    }

    /**
     * @dev Adds an appHash to the internal store, after verifying the client update proof associated with the light
     * client implementation.
     * @param peptideAppHash App hash (state root) to be verified
     * @param proof An array of bytes that contain the l1blockhash and the sequencer's signature. The first 32 bytes of
     * this argument should be the l1BlockHash, and the remaining bytes should be the sequencer signature which attests
     * to the peptide AppHash
     * for that l1BlockHash
     */
    function updateClient(bytes calldata proof, uint256 peptideHeight, uint256 peptideAppHash) external override {
        if (l1BlockProvider.hash() != bytes32(proof[:32])) {
            revert InvalidL1Origin();
        }
        _updateClient(proof, peptideHeight, peptideAppHash);
    }

    /**
     * @inheritdoc ILightClient
     */
    function getState(uint256 height) external view returns (uint256 appHash) {
        return _getState(height);
    }

    /**
     * @dev Checks if the current trusted optimistic consensus state
     * can be used to perform the membership test and if so, verifies the proof
     * @dev reverts if the proof is not valid (i.e. if the key is not included in the proof)
     */
    function verifyMembership(Ics23Proof calldata proof, bytes calldata key, bytes calldata expectedValue)
        external
        view
    {
        // a proof generated at height H can only be verified against state root (app hash) from block H - 1.
        // this means the relayer must have updated the contract with the app hash from the previous block and
        // that is why we use proof.height - 1 here.
        verifier.verifyMembership(bytes32(_getState(proof.height - 1)), key, expectedValue, proof);
    }

    /**
     * @dev Verifies that the given key is not included in the proof
     */
    function verifyNonMembership(Ics23Proof calldata, bytes calldata) external pure {
        revert NonMembershipProofsNotYetImplemented();
    }

    function _updateClient(bytes calldata proof, uint256 peptideHeight, uint256 peptideAppHash) internal {
        if (consensusStates[peptideHeight] != 0) {
            // Note: we don't cache consensusStates[peptideHeight] in mem for gas savings here because this if statement
            // would not be triggered too often, and would make the more frequent branch more expensive due to mem
            // allocation.
            if (consensusStates[peptideHeight] != peptideAppHash) {
                revert CannotUpdateClientWithDifferentAppHash();
            }
            return;
        }
        verifier.verifyStateUpdate(peptideHeight, bytes32(peptideAppHash), bytes32(proof[:32]), proof[32:]);
        consensusStates[peptideHeight] = peptideAppHash;
    }

    /**
     * @dev Returns the internal state of the light client at a given height.
     */
    function _getState(uint256 height) internal view returns (uint256 appHash) {
        return consensusStates[height];
    }
}
