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

import {L1Header, Ics23Proof} from "../interfaces/IProofVerifier.sol";
import {ISignatureVerifier} from "../interfaces/ISignatureVerifier.sol";
import {ILightClient} from "../interfaces/ILightClient.sol";

/**
 * @title SubfinalityLightClient
 * @author Polymer Labs
 * @dev This specific light client implementation uses the same client that is used in the op-stack
 */
contract SequencerSoloClient is ILightClient {
    // consensusStates maps from the height to the appHash.
    mapping(uint256 => uint256) public consensusStates;

    ISignatureVerifier public immutable verifier;

    error CannotUpdatePendingOptimisticConsensusState();
    error AppHashHasNotPassedFraudProofWindow();
    error NonMemberShipProofsNotYetImplemented();
    error NoConsensusStateAtHeight(uint256 height);

    constructor(ISignatureVerifier verifier_) {
        verifier = verifier_;
    }

    /**
     * @inheritdoc ILightClient
     * @param signature The signature of the sequencer that the client verifies the signing of
     */
    function updateClient(L1Header calldata l1header, bytes calldata signature, uint256 height, uint256 appHash)
        external
        override
        returns (uint256 fraudProofEndTime, bool ended)
    {
        uint256 hash = consensusStates[height];
        if (hash != appHash) {
            revert CannotUpdatePendingOptimisticConsensusState();
        }

        verifier.verifyStateUpdate(l1header, bytes32(appHash), height, signature);

        consensusStates[height] = appHash;

        ended = true;
        fraudProofEndTime = 0; // No fraud proof window for sequencer client so we're always past fraud proof time
    }

    /**
     * @inheritdoc ILightClient
     */
    function getState(uint256 height) external view returns (uint256 appHash, uint256 fraudProofEndTime, bool ended) {
        return getInternalState(height);
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
        (uint256 appHash,,) = getInternalState(proof.height - 1);
        verifier.verifyMembership(bytes32(appHash), key, expectedValue, proof);
    }

    /**
     * @inheritdoc ILightClient
     */
    function getFraudProofEndtime(uint256 height) external view returns (uint256 fraudProofendTime) {
        if (consensusStates[height] == 0) {
            revert NoConsensusStateAtHeight(height);
        }

        fraudProofendTime = 0;
        // Always return 0 since no end time for proofs
    }

    /**
     * @inheritdoc ILightClient
     */
    function verifyNonMembership(Ics23Proof calldata, bytes calldata) external pure {
        revert NonMemberShipProofsNotYetImplemented();
    }

    /**
     * @dev Returns the internal state of the light client at a given height.
     */
    function getInternalState(uint256 height)
        public
        view
        returns (uint256 appHash, uint256 fraudProofEndTime, bool ended)
    {
        return (consensusStates[height], 0, true);
    }
}
