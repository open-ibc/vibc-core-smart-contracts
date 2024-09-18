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
import {ISignatureVerifier} from "../interfaces/ISignatureVerifier.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {AppStateVerifier} from "../base/AppStateVerifier.sol";

/**
 * @title OptimisticProofVerifier
 * @notice Verifies proofs related to Optimistic Rollup state updates
 * @author Polymer Labs
 */
contract SequencerSignatureVerifier is AppStateVerifier, ISignatureVerifier {
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for bytes;

    // @notice known L2 Output Oracle contract address to verify state update proofs against
    address public immutable SEQUENCER;
    bytes32 public immutable CHAIN_ID; // Chain ID of the L2 chain for which the sequencer signs over

    constructor(address sequencer_, bytes32 chainId_) {
        SEQUENCER = sequencer_;
        CHAIN_ID = chainId_;
    }

    function verifyStateUpdate(uint256 l2BlockNumber, bytes32 appHash, bytes32 l1BlockHash, bytes calldata signature)
        external
        view
    {
        _verifySequencerSignature(l2BlockNumber, appHash, l1BlockHash, signature);
    }

    function _verifySequencerSignature(
        uint256 l2BlockNumber,
        bytes32 appHash,
        bytes32 l1BlockHash,
        bytes calldata signature
    ) internal view {
        // Signature will be the keccak256 hash of a + b, where:
        // a =  (empty bytes 32 corresponding to the Domain) + ( peptide chain Id )
        // b = keccak((l2 block number ) + ( l2 apphash) + ( l1 origin hash))
        if (
            ECDSA.recover(
                keccak256(
                    bytes.concat(
                        bytes32(0), CHAIN_ID, keccak256(bytes.concat(bytes32(l2BlockNumber), appHash, l1BlockHash))
                    )
                ),
                signature
            ) != SEQUENCER
        ) {
            revert InvalidSequencerSignature();
        }
    }
}
