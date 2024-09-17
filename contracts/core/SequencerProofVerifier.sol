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

import {RLPReader} from "optimism/libraries/rlp/RLPReader.sol";
import {RLPWriter} from "optimism/libraries/rlp/RLPWriter.sol";
import {L1Header, Ics23Proof, OpIcs23Proof} from "../interfaces/IProofVerifier.sol";
import {ISignatureVerifier} from "../interfaces/ISignatureVerifier.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

/**
 * @title OptimisticProofVerifier
 * @notice Verifies proofs related to Optimistic Rollup state updates
 * @author Polymer Labs
 */
contract SequencerSignatureVerifier is ISignatureVerifier {
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for bytes;

    // @notice known L2 Output Oracle contract address to verify state update proofs against
    address public immutable sequencer;

    error InvalidSequencerSignature();

    constructor(address sequencer_) {
        sequencer = sequencer_;
    }

    function verifyStateUpdate(L1Header calldata l1header, bytes32 appHash, uint256 l2Height, bytes calldata signature)
        external
        view
    {
        _verifySequencerSignature(
            l2Height, appHash, keccak256(RLPWriter.writeList(l1header.header)), l1header.stateRoot, signature
        );
    }

    /**
     * @dev Prove that a given state is not part of a proof
     * @dev this method is mainly used for packet timeouts, which is currently not implemented
     */
    function verifyNonMembership(bytes32, bytes calldata, Ics23Proof calldata) external pure {
        revert MethodNotImplemented();
    }

    /**
     * @inheritdoc ISignatureVerifier
     * @dev verifies a chain of ICS23 proofs
     * Each computed subroot starting from index 0 must match the value of the next proof (hence chained proofs).
     * The cosmos SDK and ics23 support chained proofs to switch between different proof specs.
     * Custom proof specs are not supported here. Only Iavl and Tendermint or similar proof specs are supported.
     */
    function verifyMembership(bytes32 appHash, bytes calldata key, bytes calldata value, Ics23Proof calldata proofs)
        external
        pure
    {
        // first check that the provided proof indeed proves the keys and values.
        if (keccak256(key) != keccak256(proofs.proof[0].key)) {
            revert InvalidProofKey();
        }
        if (keccak256(value) != keccak256(proofs.proof[0].value)) revert InvalidProofValue();
        // proofs are chained backwards. First proof in the list (proof[0]) corresponds to the packet proof, meaning
        // that can be checked against the next subroot value (i.e. ibc root). Once the first proof is verified,
        // we can check the second that corresponds to the ibc proof, that is checked against the app hash (app root)
        if (bytes32(proofs.proof[1].value) != _verify(proofs.proof[0])) revert InvalidPacketProof();
        if (appHash != _verify(proofs.proof[1])) revert InvalidIbcStateProof();
    }

    function _verifySequencerSignature(
        uint256 L2BlockNumber,
        bytes32 AppHash,
        bytes32 L1BlockHash,
        bytes32 L1StateRoot,
        bytes calldata signature
    ) internal view {
        // TODO: check that erecover should indeed ensure that the signature signs over the hash
        if (
            ECDSA.recover(keccak256(bytes.concat(bytes32(L2BlockNumber), AppHash, L1BlockHash, L1StateRoot)), signature)
                != sequencer
        ) {
            revert InvalidSequencerSignature();
        }
    }

    /**
     * @dev Verifies an ICS23 proof through the root hash based on the provided proof.
     * @dev This code was adapted from the ICS23 membership verification found here:
     * https://github.com/cosmos/ics23/blob/go/v0.10.0/go/ics23.go#L36
     * @param proof The ICS23 proof to be verified.
     * @return computed The computed root hash.
     */
    function _verify(OpIcs23Proof calldata proof) internal pure returns (bytes32 computed) {
        bytes32 hashedData = sha256(proof.value);
        computed = sha256(
            abi.encodePacked(
                proof.prefix, _encodeVarint(proof.key.length), proof.key, _encodeVarint(hashedData.length), hashedData
            )
        );

        for (uint256 i = 0; i < proof.path.length; i++) {
            computed = sha256(abi.encodePacked(proof.path[i].prefix, computed, proof.path[i].suffix));
        }
    }

    /**
     * @dev Encodes an integer value into a variable-length integer format.
     * @param value The integer value to be encoded.
     * @return encoded The encoded bytes array.
     */
    function _encodeVarint(uint256 value) internal pure returns (bytes memory encoded) {
        bytes memory result;
        while (value >= 0x80) {
            bytes.concat(result, bytes1(uint8((value & 0x7F) | 0x80)));
            value >>= 7;
        }
        encoded = bytes.concat(result, bytes1(uint8(value)));
    }
}
