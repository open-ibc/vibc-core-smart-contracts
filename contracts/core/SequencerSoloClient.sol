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
 * @title SubfinalityLightClient
 * @author Polymer Labs
 * @dev This specific light client implementation uses the same client that is used in the op-stack
 */
contract SequencerSoloClient is ILightClient {
    error InvalidL1Origin();

    LightClientType public constant LIGHT_CLIENT_TYPE = LightClientType.SequencerLightClient; // Stored as a constant
        // for cheap on-chain use
    uint8 private _LightClientType = uint8(LIGHT_CLIENT_TYPE); // Also redundantly stored as a private mutable type in
        // case it needs to be accessed in any proofs

    // consensusStates maps from the height to the appHash.
    mapping(uint256 => uint256) public consensusStates;

    ISignatureVerifier public immutable verifier;
    L1Block public immutable l1BlockProvider;

    error CannotUpdatePendingOptimisticConsensusState();
    error AppHashHasNotPassedFraudProofWindow();
    error NonMembershipProofsNotYetImplemented();
    error NoConsensusStateAtHeight(uint256 height);

    constructor(ISignatureVerifier verifier_, L1Block _l1BlockProvider) {
        verifier = verifier_;
        l1BlockProvider = _l1BlockProvider;
    }

    /**
     * @inheritdoc ILightClient
     * @param proof An array of bytes that contain the l1blockhash and the sequencer's signature. The first 32 bytes of
     * this argument should be
     * the l1BlockHash, and the remaining bytes should be the sequencer signature which attests to the peptide AppHash
     * for that l1BlockHash
     */
    function updateClient(bytes calldata proof, uint256 peptideHeight, uint256 peptideAppHash) external override {
        if (l1BlockProvider.hash() != bytes32(proof[:32])) {
            revert InvalidL1Origin();
        }
        _updateClient(proof, peptideHeight, peptideAppHash);
    }

    function _updateClient(bytes calldata proof, uint256 peptideHeight, uint256 peptideAppHash) internal {
        if (consensusStates[peptideHeight] != 0) {
            return;
        }
        verifier.verifyStateUpdate(peptideHeight, bytes32(peptideAppHash), bytes32(proof[:32]), proof[32:]);
        consensusStates[peptideHeight] = peptideAppHash;
    }

    /**
     * @inheritdoc ILightClient
     */
    function getState(uint256 height) external view returns (uint256 appHash) {
        return _getState(height);
    }
    /**
     * @inheritdoc ILightClient
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
     * @inheritdoc ILightClient
     */
    function verifyNonMembership(Ics23Proof calldata, bytes calldata) external pure {
        revert NonMembershipProofsNotYetImplemented();
    }

    /**
     * @dev Returns the internal state of the light client at a given height.
     */
    function _getState(uint256 height) internal view returns (uint256 appHash) {
        return consensusStates[height];
    }
}
