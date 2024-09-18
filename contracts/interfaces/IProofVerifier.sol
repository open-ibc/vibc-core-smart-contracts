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

// OpIcs23ProofPath represents a commitment path in an ICS23 proof, which consists of a commitment prefix and a suffix.
struct OpIcs23ProofPath {
    bytes prefix;
    bytes suffix;
}

// OpIcs23Proof represents an ICS23 proof
struct OpIcs23Proof {
    OpIcs23ProofPath[] path;
    bytes key;
    bytes value;
    bytes prefix;
}

// the Ics23 proof related structs are used to do membership verification. These are not the actual Ics23
// format but a "solidity friendly" version of it - data is the same just packaged differently
struct Ics23Proof {
    OpIcs23Proof[] proof;
    uint256 height;
}

// This is the proof we use to verify the apphash (state) updates.
struct OpL2StateProof {
    bytes[] accountProof;
    bytes[] outputRootProof;
    bytes32 l2OutputProposalKey;
    bytes32 l2BlockHash;
}

// The `header` field is a list of RLP encoded L1 header fields. Both stateRoot and number are not
// encoded for easy usage. They must match with their RLP encoded counterparty versions.
struct L1Header {
    bytes[] header;
    bytes32 stateRoot;
    uint64 number;
}


interface IAppStateVerifier {
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

/**
 * @title IProofVerifier
 * @author Polymer Labs
 * @notice An interface that abstracts away proof verification logic for light clients
 */
interface IProofVerifier is IAppStateVerifier {
    /**
     * @dev verifies if a state update (apphash) is valid, given the provided proofs.
     *      Reverts in case of failure.
     *
     * @param l1header RLP "encoded" version of the L1 header that matches with the trusted hash and number
     * @param proof l2 state proof. It includes the keys, hashes and storage proofs required to verify the app hash
     * @param appHash  l2 app hash (state root) to be verified
     * @param trustedL1BlockHash trusted L1 block hash. Provided L1 header must match with it.
     * @param trustedL1BlockNumber trusted L1 block number. Provided L1 header must match with it.
     */
    function verifyStateUpdate(
        L1Header calldata l1header,
        OpL2StateProof calldata proof,
        bytes32 appHash,
        bytes32 trustedL1BlockHash,
        uint64 trustedL1BlockNumber
    ) external view;


}
