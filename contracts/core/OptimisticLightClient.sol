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
import {LightClientType} from "../interfaces/ILightClient.sol";

/**
 * @title OptimisticLightClient
 * @author Polymer Labs
 * @dev This specific light client implementation uses the same client that is used in the op-stack
 */
contract OptimisticLightClient is IOptimisticLightClient {
    LightClientType public constant LIGHT_CLIENT_TYPE = LightClientType.OptimisticLightClient; // Stored as a constant
        // for
        // cheap on-chain use
    uint8 private _LightClientType = uint8(LIGHT_CLIENT_TYPE); // Also redundantly stored as a private mutable type in
        // case it needs to be accessed in any proofs

    // consensusStates maps from the polymerHeight to the polymerAppHash.
    mapping(uint256 => uint256) public consensusStates;

    // fraudProofEndtime maps from the polymerAppHash to the fraud proof end time.
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
     * @dev Adds an appHash to the internal store, after verifying the client update proof associated with the light
     * client implementation.
     * @param proof A generic byte array that contains proof data to prove the apphash client update. This can differ
     * depending on the light client type. E.g. this can be an abi.encoded struct which contains an OpL2StateProof and
     * L1Block from the IProofVerifier
     * interface.
     * @param polymerAppHash App hash (state root) to be verified
     */
    function updateClient(bytes calldata proof, uint256 polymerHeight, uint256 polymerAppHash) external override {
        (L1Header memory l1header, OpL2StateProof memory stateProof) = abi.decode(proof, (L1Header, OpL2StateProof));
        uint256 hash = consensusStates[polymerHeight];
        if (hash == 0) {
            // if this is a new polymerAppHash we need to verify the provided proof. This method will revert in case
            // of invalid proof.
            verifier.verifyStateUpdate(
                l1header, stateProof, bytes32(polymerAppHash), l1BlockProvider.hash(), l1BlockProvider.number()
            );

            // a new polymerAppHash
            consensusStates[polymerHeight] = polymerAppHash;
            fraudProofEndtime[polymerAppHash] = block.timestamp + fraudProofWindowSeconds;
        } else if (hash != polymerAppHash) {
            revert CannotUpdatePendingOptimisticConsensusState();
        }
    }

    /**
     * @inheritdoc ILightClient
     */
    function getState(uint256 polymerHeight) external view returns (uint256 polymerAppHash) {
        return _getState(polymerHeight);
    }

    /**
     * @inheritdoc IOptimisticLightClient
     */
    function getStateAndEndTime(uint256 polymerHeight)
        external
        view
        returns (uint256 polymerAppHash, uint256 fraudProofEndTime, bool ended)
    {
        return _getStateAndEndTime(polymerHeight);
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
        (uint256 polymerAppHash,, bool ended) = _getStateAndEndTime(proof.height - 1);
        if (!ended) revert AppHashHasNotPassedFraudProofWindow();
        verifier.verifyMembership(bytes32(polymerAppHash), key, expectedValue, proof);
    }

    /**
     * @dev Verifies that the given key is not included in the proof
     */
    function verifyNonMembership(Ics23Proof calldata proof, bytes calldata key) external view {
        (uint256 polymerAppHash,, bool ended) = _getStateAndEndTime(proof.height - 1);
        if (!ended) revert AppHashHasNotPassedFraudProofWindow();
        verifier.verifyNonMembership(bytes32(polymerAppHash), key, proof);
    }

    /**
     * @inheritdoc IOptimisticLightClient
     */
    function getFraudProofEndtime(uint256 polymerHeight) external view returns (uint256 fraudProofEndTime) {
        uint256 hash = consensusStates[polymerHeight];
        return fraudProofEndtime[hash];
    }

    function _getState(uint256 height) internal view returns (uint256 polymerAppHash) {
        polymerAppHash = consensusStates[height];
    }

    /**
     * @dev Returns the internal state of the light client at a given polymerHeight.
     */
    function _getStateAndEndTime(uint256 polymerHeight)
        internal
        view
        returns (uint256 polymerAppHash, uint256 fraudProofEndTime, bool ended)
    {
        uint256 hash = consensusStates[polymerHeight];
        return (hash, fraudProofEndtime[hash], hash != 0 && block.timestamp >= fraudProofEndtime[hash]);
    }
}
