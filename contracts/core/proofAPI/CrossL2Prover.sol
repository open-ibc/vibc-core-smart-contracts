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

    string public clientType;
    ISignatureVerifier public immutable verifier;

    // Store peptide apphashes for a given height
    mapping(uint256 => uint256) public peptideAppHashes;

    error CannotUpdateClientWithDifferentAppHash();

    constructor(ISignatureVerifier verifier_, string memory clientType_) {
        verifier = verifier_;
        clientType = clientType_;
    }

    /**
     * @inheritdoc ICrossL2Prover
     */
    function validateReceipt(bytes calldata proof) public view returns (bytes32 srcChainID, bytes memory receiptRLP) {
        (
            Ics23Proof memory peptideAppProof,
            bytes[] memory receiptMMPTProof,
            bytes32 receiptRoot,
            uint64 eventHeight,
            string memory srcChainId,
            bytes memory receiptIndex
        ) = abi.decode(proof, (Ics23Proof, bytes[], bytes32, uint64, string, bytes));
        // Before we can trust the receipt root, we first need to verify that the receipt root is indeed stored
        // on peptide at the given clientID and height.

        // VerifyMembership verifies the receipt root  through an ics23 proof of peptide state that attests that the
        // given eventHeight has the receipt root at the peptide height
        this.verifyMembership(
            bytes32(_getPeptideAppHash(peptideAppProof.height - 1)), // a proof generated at height H can only be
                // verified against state root (app hash) from block H - 1. this means the relayer must have updated the
                // contract with the app hash from the previous block and that is why we use proof.height - 1 here.
            Ibc.receiptRootKey(srcChainId, clientType, eventHeight),
            abi.encodePacked(receiptRoot),
            peptideAppProof
        );

        // Now that verifyMembership passed, we can now trust the receiptRoot.
        // Now we just simply have to prove that raw receipt is indeed part of the receipt root at the given receipt
        // index.
        // This is done through a Merkle proof.

        return (srcChainID, MerkleTrie.get(receiptIndex, receiptMMPTProof, receiptRoot));
    }

    function validateLog(uint256 logIndex, bytes calldata proof)
        external
        view
        returns (bytes32 chainId, address emittingContract, bytes[] memory topics, bytes memory unindexedData)
    {
        bytes memory receiptRLP;
        (chainId, receiptRLP) = validateReceipt(proof);
        // The first byte is a RLP encoded receipt type so slice it off.
        RLPReader.RLPItem[] memory receipt = RLPReader.readList(Bytes.slice(receiptRLP, 1, receiptRLP.length - 1));
        /*
            // RLP encoded receipt has the following structure. Logs are the 4th RLP list item.
            type ReceiptRLP struct {
                    PostStateOrStatus []byte
                    CumulativeGasUsed uint64
                    Bloom             Bloom
                    Logs              []*Log
            }
        */

        // Each log itself is an rlp encoded datatype of 3 properties:
        // type Log struct {
        //         senderAddress bytes // contract address where this log was emitted from
        //         topics bytes        // Array of indexed topics. The first element is the 32-byte selector of the
        // event (can use TransmitToHouston.selector), and the following  elements in this array are the abi encoded
        // arguments individually
        //         topics data         // abi encoded raw bytes of unindexed data
        // }
        RLPReader.RLPItem[] memory log = RLPReader.readList(RLPReader.readList(receipt[3])[logIndex]);

        emittingContract = address(uint160(uint256(bytes32(RLPReader.readBytes(log[0])))));
        RLPReader.RLPItem[] memory encodedTopics = RLPReader.readList(log[1]);
        unindexedData = (RLPReader.readBytes(log[2])); // This is the raw unindexed data. in this case it's
            // just an abi encoded uint64

        for (uint256 i = 0; i < encodedTopics.length; i++) {
            topics[i] = RLPReader.readBytes(encodedTopics[i]);
        }
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
