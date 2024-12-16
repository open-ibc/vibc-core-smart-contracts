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
import {IClientUpdates} from "./ILightClient.sol";

/**
 * @title ICrossL2Prover
 * @author Polymer Labs
 * @notice A contract that can prove peptides state. Since peptide is an aggregator of many chains' states, this
 * contract can in turn be used to prove any arbitrary events and/or storage on counterparty chains.
 */
interface ICrossL2Prover is IClientUpdates {
    struct EventProof {
        Ics23Proof iavlProof; // Proof that the receipt root exists within an iavl tree at a given apphash stored in
            // this contract
        bytes[] receiptMMPTProof; // A modified Merkle Patricia tree proof that proves the raw RLP receipt bytes are
            // included in the receipt root
        bytes32 receiptRoot; // The receipt root that is is proven against in the MMPT proof
        uint64 eventHeight; // The source L2 block the receipt or event we are trying to prove was emitted in. The
            // receiptRoot of the block at this eight should be the the same as the receipt root given in this struct
    }

    // An event decoded from a counterparty chain.
    struct DecodedEvent {
        address emittingContractAddress;
        bytes[] topics; // The topics of the event. First topic is the event signature that can be calculated by
            // Event.selector. The remaining elements in this array are the indexed parameters of the event.
        bytes data; // The data of the event. This is the abi encoded non-indexed parameters of the event.
    }

    /**
     *   @notice Validates a cross chain receipt from a couterparty chain
     *   @notice Some of the data needed to call this method will be returned by polymer's proof api. This method should
     * only be
     * called after we are certain that polymer's relayer has called updateClient on this contract to update this light
     * client
     * @notice This method reverts if a proof is invalid.
     * @param receiptIndex: The index of the receipt in the block.  This is the same as the key the receipt bytes are
     * stored in the MMPT.
     * @param proof: The proof data needed to prove the receipt. This is returned from Polymer's proof API and should
     * be passed in as is. This is generally an opaque bytes object but it is constructed through abi encoding the proof
     * fields from the above EventProof struct in this interface. See the struct for per-property documentation.
     * @return chainId The source chain id the proof was for. Dapps which call this method should always validate this
     * chain id is indeed expected to prevent spoofing source chains (e.g. a proof might be valid for an event occurring
     * on another chain, but not for the chain you expect)
     * @return ReceiptRLP The raw rlp encoded bytes for the receipt, which can further be parsed to extract all logs for
     * the given receipt. This is equivalent to the leaf value stored in the receipt MMPT and is just an RLP encoded
     * object of the type:
     * type ReceiptRLP struct {
     *     PostStateOrStatus []byte
     *     CumulativeGasUsed uint64
     *     Bloom             Bloom
     *     Logs              []*Log
     *     }
     */
    function validateReceipt(bytes calldata receiptIndex, bytes calldata proof)
        external
        returns (bytes calldata chainId, bytes calldata receiptRLPEncodedBytes);

    /**
     *   @notice Validates a cross chain event from a couterparty chain
     *   @notice Some of the data needed to call this method will be returned by polymer's proof api. This method should
     * only be called after we are certain that polymer's relayer has called updateClient on this contract to update
     * this light client
     * @param receiptIndex: The index of the receipt in the block.  This is the same as the key the receipt bytes are
     * stored in the MMPT.
     * @param receiptRLPEncodedBytes: The raw rlp encoded bytes of the whole receipt object we are trying to prove, this
     * is the value stored in the MMPT
     * @param logIndex: The index of the log in the logs array of the receipt. NOTE: This is not the log index within
     * the block, only the log index within the receipt.
     * @param logBytes: The raw bytes of the log we are trying to prove. This is the value stored in the MMPT.
     * @param proof: The proof data needed to prove the receipt. This is returned from Polymer's proof API and should
     * be passed in as is. This is generally an opaque bytes object but it is constructed through abi encoding the proof
     * fields from the above EventProof struct in this interface. See the struct for per-property documentation.
     */
    function validateEvent(bytes calldata receiptIndex, uint256 logIndex, bytes calldata proof)
        external
        returns (bytes calldata chainId, bytes calldata DecodedEvent);

    /**
     * Returns the peptide at a given apphash at a given height,
     */
    function getState(uint256 height) external view returns (uint256);
}
