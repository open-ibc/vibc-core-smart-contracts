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
// solhint-disable-next-line
import {ILightClient} from "../interfaces/ILightClient.sol"; // we need this import to generate docs
import {L1Block} from "optimism/L2/L1Block.sol";
import {IOptimisticLightClient} from "../interfaces/IOptimisticLightClient.sol";

/**
 * @title OptimisticLightClient
 * @author Polymer Labs
 * @dev This specific light client implementation uses the same client that is used in the op-stack
 */
contract OptimisticLightClient is IOptimisticLightClient {
    // consensusStates maps from the peptideHeight to the peptideAppHash.
    mapping(uint256 => uint256) public consensusStates;

    // fraudProofEndtime maps from the peptideAppHash to the fraud proof end time.
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
    function updateClient(bytes calldata proof, uint256 peptideHeight, uint256 peptideAppHash) external override {
        (L1Header memory l1header, OpL2StateProof memory stateProof) = abi.decode(proof, (L1Header, OpL2StateProof));
        uint256 hash = consensusStates[peptideHeight];
        if (hash == 0) {
            // if this is a new peptideAppHash we need to verify the provided proof. This method will revert in case
            // of invalid proof.
            verifier.verifyStateUpdate(
                l1header, stateProof, bytes32(peptideAppHash), l1BlockProvider.hash(), l1BlockProvider.number()
            );

            // a new peptideAppHash
            consensusStates[peptideHeight] = peptideAppHash;
            fraudProofEndtime[peptideAppHash] = block.timestamp + fraudProofWindowSeconds;
        } else if (hash != peptideAppHash) {
            revert CannotUpdatePendingOptimisticConsensusState();
        }
    }

    /**
     * @inheritdoc ILightClient
     */
    function getState(uint256 peptideHeight) external view returns (uint256 peptideAppHash) {
        return _getState(peptideHeight);
    }

    /**
     * @inheritdoc IOptimisticLightClient
     */
    function getStateAndEndTime(uint256 peptideHeight)
        external
        view
        returns (uint256 peptideAppHash, uint256 fraudProofEndTime, bool ended)
    {
        return _getStateAndEndTime(peptideHeight);
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
        (uint256 peptideAppHash,, bool ended) = _getStateAndEndTime(proof.height - 1);
        if (!ended) revert AppHashHasNotPassedFraudProofWindow();
        verifier.verifyMembership(bytes32(peptideAppHash), key, expectedValue, proof);
    }

    /**
     * @inheritdoc ILightClient
     */
    function verifyNonMembership(Ics23Proof calldata proof, bytes calldata key) external view {
        (uint256 peptideAppHash,, bool ended) = _getStateAndEndTime(proof.height - 1);
        if (!ended) revert AppHashHasNotPassedFraudProofWindow();
        verifier.verifyNonMembership(bytes32(peptideAppHash), key, proof);
    }

    /**
     * @inheritdoc IOptimisticLightClient
     */
    function getFraudProofEndtime(uint256 peptideHeight) external view returns (uint256 fraudProofEndTime) {
        uint256 hash = consensusStates[peptideHeight];
        return fraudProofEndtime[hash];
    }

    function _getState(uint256 height) internal view returns (uint256 peptideAppHash) {
        peptideAppHash = consensusStates[height];
    }

    /**
     * @dev Returns the internal state of the light client at a given peptideHeight.
     */
    function _getStateAndEndTime(uint256 peptideHeight)
        internal
        view
        returns (uint256 peptideAppHash, uint256 fraudProofEndTime, bool ended)
    {
        uint256 hash = consensusStates[peptideHeight];
        return (hash, fraudProofEndtime[hash], hash != 0 && block.timestamp >= fraudProofEndtime[hash]);
    }
}
