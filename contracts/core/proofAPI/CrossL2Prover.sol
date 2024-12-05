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

import {RLPReader} from "optimism/libraries/rlp/RLPReader.sol";
import {MerkleTrie} from "optimism/libraries/trie/MerkleTrie.sol";
import {Bytes} from "optimism/libraries/Bytes.sol";
import {Ibc} from "../../libs/Ibc.sol";
import {AppStateVerifier} from "../../base/AppStateVerifier.sol";
import {ICrossL2Prover} from "../../interfaces/ICrossL2Prover.sol";
import {Ics23Proof} from "../../interfaces/IProofVerifier.sol";
import {ISignatureVerifier} from "../../interfaces/ISignatureVerifier.sol";
import {LightClientType} from "../../interfaces/ILightClient.sol";

contract CrossL2Prover is AppStateVerifier, ICrossL2Prover {
    LightClientType public constant LIGHT_CLIENT_TYPE = LightClientType.SequencerLightClient; // Stored as a constant
        // for cheap on-chain use

    string public peptideClientId;
    ISignatureVerifier public immutable verifier;

    // Store peptide apphashes for a given height
    mapping(uint256 => uint256) public peptideAppHashes;

    error CannotUpdateClientWithDifferentAppHash();

    constructor(ISignatureVerifier verifier_, string memory peptideClientId_) {
        verifier = verifier_;
        peptideClientId = peptideClientId_;
    }

    /**
     * @inheritdoc ICrossL2Prover
     */
    function validateReceipt(bytes calldata receiptIndex, bytes calldata receiptRLPEncodedBytes, bytes calldata proof)
        public
        view
        returns (bool valid)
    {
        (Ics23Proof memory peptideAppProof, bytes[] memory receiptMMPTProof, bytes32 receiptRoot, uint64 eventHeight) =
            abi.decode(proof, (Ics23Proof, bytes[], bytes32, uint64));
        // Before we can trust the receipt root, we first need to verify that the receipt root is indeed stored
        // on peptide at the given clientID and height.

        // VerifyMembership verifies the receipt root  through an ics23 proof of peptide state that attests that the
        // given eventHeight has the receipt root at the peptide height
        this.verifyMembership(
            bytes32(_getPeptideAppHash(peptideAppProof.height)),
            Ibc.receiptRootKey(peptideClientId, eventHeight),
            abi.encodePacked(receiptRoot),
            peptideAppProof
        );

        // Now that verifyMembership passed, we can now trust the receiptRoot.
        // Now we just simply have to prove that raw receipt is indeed part of the receipt root at the given receipt
        // index.
        // This is done through a Merkle proof.
        valid = MerkleTrie.verifyInclusionProof(
            abi.encodePacked(receiptIndex), receiptRLPEncodedBytes, receiptMMPTProof, receiptRoot
        );
    }

    function validateEvent(
        bytes calldata receiptIndex,
        bytes calldata receiptRLPEncodedBytes,
        uint256 logIndex,
        bytes calldata logBytes,
        bytes calldata proof
    ) external view returns (bool) {
        // First we validate an event by validating the receipt
        if (!validateReceipt(receiptIndex, receiptRLPEncodedBytes, proof)) {
            return false;
        }

        // The first byte is a RLP encoded receipt type so slice it off.
        RLPReader.RLPItem[] memory receipt =
            RLPReader.readList(Bytes.slice(receiptRLPEncodedBytes, 1, receiptRLPEncodedBytes.length - 1));
        /*
            // RLP encoded receipt has the following structure. Logs are the 4th RLP list item.
            type ReceiptRLP struct {
                    PostStateOrStatus []byte
                    CumulativeGasUsed uint64
                    Bloom             Bloom
                    Logs              []*Log
            }
        */
        RLPReader.RLPItem[] memory logs = RLPReader.readList(receipt[3]);

        // Then we decode the rlp bytes and check that the logBytes is the same as the log at the given index
        return keccak256(RLPReader.readRawBytes(logs[logIndex])) == keccak256(logBytes);
    }

    /**
     * @dev Adds an appHash to the internal store, after verifying the client update proof associated with the light
     * client implementation.
     * @param peptideAppHash App hash (state root) to be verified
     * @param proof An array of bytes that contain the l1blockhash and the sequencer's signature. The first 32 bytes of
     * this argument should be the l1BlockHash, and the remaining bytes should be the sequencer signature which attests
     * to the peptide AppHash
     * for that l1BlockHash
     */
    function updateClient(bytes calldata proof, uint256 peptideHeight, uint256 peptideAppHash) external {
        _updateClient(proof, peptideHeight, peptideAppHash);
    }

    function getState(uint256 height) external view returns (uint256) {
        return _getPeptideAppHash(height);
    }

    function _getPeptideAppHash(uint256 _height) internal view returns (uint256) {
        return peptideAppHashes[_height];
    }

    function _updateClient(bytes calldata proof, uint256 peptideHeight, uint256 peptideAppHash) internal {
        if (peptideAppHashes[peptideHeight] != 0) {
            // Note: we don't cache peptideAppHash[peptideHeight] in mem for gas savings here because this if
            // statement
            // would not be triggered too often, and would make the more frequent branch more expensive due to mem
            // allocation.
            if (peptideAppHashes[peptideHeight] != peptideAppHash) {
                revert CannotUpdateClientWithDifferentAppHash();
            }
            return;
        }
        verifier.verifyStateUpdate(peptideHeight, bytes32(peptideAppHash), bytes32(proof[:32]), proof[32:]);
        peptideAppHashes[peptideHeight] = peptideAppHash;
    }
}
