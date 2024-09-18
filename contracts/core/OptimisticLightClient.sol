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

import {L1Header, IProofVerifier, OpL2StateProof, Ics23Proof} from "../interfaces/IProofVerifier.sol";
import {ILightClient} from "../interfaces/ILightClient.sol";
import {L1Block} from "optimism/L2/L1Block.sol";

/**
 * @title OptimisticLightClient
 * @author Polymer Labs
 * @dev This specific light client implementation uses the same client that is used in the op-stack
 */
contract OptimisticLightClient is ILightClient {
    // consensusStates maps from the height to the appHash.
    mapping(uint256 => uint256) public consensusStates;

    // fraudProofEndtime maps from the appHash to the fraud proof end time.
    mapping(uint256 => uint256) public fraudProofEndtime;
    uint256 public fraudProofWindowSeconds;
    IProofVerifier public verifier;
    L1Block public l1BlockProvider;

    error CannotUpdatePendingOptimisticConsensusState();
    error AppHashHasNotPassedFraudProofWindow();

    constructor(uint32 fraudProofWindowSeconds_, IProofVerifier verifier_, L1Block _l1BlockProvider) {
        fraudProofWindowSeconds = fraudProofWindowSeconds_;
        verifier = verifier_;
        l1BlockProvider = _l1BlockProvider;
    }

    /**
     * @inheritdoc ILightClient
     */
    function updateClient(bytes calldata l1headerbytes, bytes calldata proof, uint256 height, uint256 appHash)
        external
        override
        returns (uint256 fraudProofEndTime, bool ended)
    {
        L1Header memory l1header = abi.decode(l1headerbytes, (L1Header));
        uint256 hash = consensusStates[height];
        if (hash == 0) {
            // if this is a new apphash we need to verify the provided proof. This method will revert in case
            // of invalid proof.
            verifier.verifyStateUpdate(
                l1header,
                abi.decode(proof, (OpL2StateProof)),
                bytes32(appHash),
                l1BlockProvider.hash(),
                l1BlockProvider.number()
            );

            // a new appHash
            consensusStates[height] = appHash;
            uint256 endTime = block.timestamp + fraudProofWindowSeconds;
            fraudProofEndtime[appHash] = endTime;
            return (endTime, false);
        }

        if (hash == appHash) {
            uint256 endTime = fraudProofEndtime[hash];
            return (endTime, block.timestamp >= endTime);
        }

        revert CannotUpdatePendingOptimisticConsensusState();
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
        (uint256 appHash,, bool ended) = getInternalState(proof.height - 1);
        if (!ended) revert AppHashHasNotPassedFraudProofWindow();
        verifier.verifyMembership(bytes32(appHash), key, expectedValue, proof);
    }

    /**
     * @inheritdoc ILightClient
     */
    function verifyNonMembership(Ics23Proof calldata proof, bytes calldata key) external view {
        (uint256 appHash,, bool ended) = getInternalState(proof.height - 1);
        if (!ended) revert AppHashHasNotPassedFraudProofWindow();
        verifier.verifyNonMembership(bytes32(appHash), key, proof);
    }

    /**
     * @inheritdoc ILightClient
     */
    function getFraudProofEndtime(uint256 height) external view returns (uint256 fraudProofEndTime) {
        uint256 hash = consensusStates[height];
        return fraudProofEndtime[hash];
    }

    /**
     * @dev Returns the internal state of the light client at a given height.
     */
    function getInternalState(uint256 height)
        public
        view
        returns (uint256 appHash, uint256 fraudProofEndTime, bool ended)
    {
        uint256 hash = consensusStates[height];
        return (hash, fraudProofEndtime[hash], hash != 0 && block.timestamp >= fraudProofEndtime[hash]);
    }
}
