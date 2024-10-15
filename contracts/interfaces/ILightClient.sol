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
pragma solidity ^0.8.0;

import {Ics23Proof} from "./IProofVerifier.sol";

enum LightClientType {
    SimTestLightClient, // Note: not deployed on any mainnets
    OptimisticLightClient, // Our native opstack light client
    SequencerLightClient, // Our native sequencer light client, which does not check l1 origin check to cut down on
        // latency
    ReOrgResistantSequencerLightClient // Our native sequencer light client, which checks for l1 origin checks to be
        // re-org resistant

}

interface IClientUpdates {
    /**
     * @dev Adds an appHash to the internal store, after verifying the client update proof associated with the light
     * client implementation.
     * @param proof A generic byte array that contains proof data to prove the apphash client update. This can differ
     * depending on the light client type. E.g. this can be an abi.encoded struct which contains an OpL2StateProof and
     * L1Block from the IProofVerifier
     * interface.
     * @param appHash App hash (state root) to be verified
     */
    function updateClient(bytes calldata proof, uint256 height, uint256 appHash) external;

    /*
    * Returns the type of the light client, useful for relayers to know which light client implementation is at which
    address. 
    */
    function LIGHT_CLIENT_TYPE() external view returns (LightClientType);
}

interface IMembershipVerifier {
    /**
     * @dev Checks if the current trusted optimistic consensus state
     * can be used to perform the membership test and if so, verifies the proof
     * @dev reverts if the proof is not valid (i.e. if the key is not included in the proof)
     */
    function verifyMembership(Ics23Proof calldata proof, bytes calldata key, bytes memory expectedValue) external;
}

interface INonMembershipVerifier {
    /**
     * @dev Verifies that the given key is not included in the proof
     */
    function verifyNonMembership(Ics23Proof calldata proof, bytes calldata key) external;
}

/**
 * @title ILightClient
 * @author Polymer Labs
 * @notice A contract that manages the merkle root(appHash) at different block heights of a chain and tracks the fraud
 * proof end time for them.
 * @notice This is used for state inclusion proofs
 */
interface ILightClient is IClientUpdates, IMembershipVerifier, INonMembershipVerifier {
    /**
     * Returns the apphash at a given height
     */
    function getState(uint256 height) external view returns (uint256);
}
