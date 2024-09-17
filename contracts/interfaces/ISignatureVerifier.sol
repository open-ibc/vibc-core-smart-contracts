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

import {L1Header, Ics23Proof} from "./IProofVerifier.sol";

/**
 * @title ISignatureVerifier
 * @author Polymer Labs
 * @notice An interface that abstracts away proof verification logic for light clients
 */
interface ISignatureVerifier {
    error InvalidL1BlockNumber();
    error InvalidL1BlockHash();
    error InvalidRLPEncodedL1BlockNumber();
    error InvalidRLPEncodedL1StateRoot();
    error InvalidAppHash();
    error InvalidProofKey();
    error InvalidProofValue();
    error InvalidPacketProof();
    error InvalidIbcStateProof();
    error MethodNotImplemented();

    function verifyStateUpdate(L1Header calldata l1header, bytes32 appHash, uint256 l2Height, bytes calldata signature)
        external;

    /**
     * @dev verifies the provided ICS23 proof given the trusted app hash. Reverts in case of failure.
     *
     * @param appHash trusted l2 app hash (state root)
     * @param key key to be proven
     * @param value value to be proven
     * @param proof ICS23 membership proof
     */
    function verifyMembership(bytes32 appHash, bytes calldata key, bytes calldata value, Ics23Proof calldata proof)
        external
        pure;

    /**
     * @dev verifies the provided ICS23 proof given the trusted app hash. Reverts in case of failure.
     *
     * @param appHash trusted l2 app hash (state root)
     * @param key key to be proven non-existing
     * @param proof ICS23 non-membership proof
     */
    function verifyNonMembership(bytes32 appHash, bytes calldata key, Ics23Proof calldata proof) external pure;
}
