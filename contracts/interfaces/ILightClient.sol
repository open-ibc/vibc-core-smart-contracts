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

import {Ics23Proof, L1Header} from "./IProofVerifier.sol";

/**
 * @title ILightClient
 * @author Polymer Labs
 * @notice A contract that manages the merkle root(appHash) at different block heights of a chain and tracks the fraud
 * proof end time for them.
 * @notice This is used for state inclusion proofs
 */
interface ILightClient {
    /**
     * @dev Adds an appHash to the internal store and returns the fraud proof end time, and a bool flag indicating if
     * the fraud proof window has passed according to the block's time stamp.
     * @param l1header  RLP "encoded" version of the L1 header that matches with the trusted hash and number
     * @param proof l2 state proof. It includes the keys, hashes and storage proofs required to verify the app hash
     * @param appHash Peptide app hash (state root) to be verified
     * @return fraudProofEndTime The fraud proof end time.
     * @return ended A boolean indicating if the fraud proof window has passed.
     */
    function updateClient(L1Header calldata l1header, bytes calldata proof, uint256 height, uint256 appHash)
        external
        returns (uint256 fraudProofEndTime, bool ended);

    /**
     * @dev Returns the fraud proof end time at a given block
     * 0 is returned if there isn't an appHash with the given l2 height.
     */
    function getFraudProofEndtime(uint256 height) external returns (uint256 endTime);

    /**
     * @dev Checks if the current trusted optimistic consensus state
     * can be used to perform the membership test and if so, verifies the proof
     * @dev reverts if the proof is not valid (i.e. if the key is not included in the proof)
     */
    function verifyMembership(Ics23Proof calldata proof, bytes memory key, bytes memory expectedValue) external;

    /**
     * @dev Verifies that the given key is not included in the proof
     */
    function verifyNonMembership(Ics23Proof calldata proof, bytes memory key) external;

    /**
     * @dev Returns the apphash at a given block height
     */
    function getState(uint256 height) external view returns (uint256 appHash, uint256 fraudProofEndTime, bool ended);
}
