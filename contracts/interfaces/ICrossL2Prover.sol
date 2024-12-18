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
        string srcChainId; // The chain id of the source chain the event was emitted on
        bytes receiptIdx; // The rlp encoded index of the receipt within the block. This is the same as the key of the
            // receipt MMPT
    }

    /**
     *   @notice Validates a cross chain receipt from a couterparty chain
     *   @notice Some of the data needed to call this method will be returned by polymer's proof api. This method should
     * only be
     * called after we are certain that polymer's relayer has called updateClient on this contract to update this light
     * client
     * @notice This method reverts if a proof is invalid.
     * @param proof: The proof data needed to prove the receipt. This is returned from Polymer's proof API and should
     * be passed in as is. This is generally an opaque bytes object but it is constructed through abi encoding the proof
     * fields from the above EventProof struct in this interface. See the struct for per-property documentation.
     * @return srcChainId The source chain id the proof was for. Dapps which call this method should always validate
     * this
     * chain id is indeed expected to prevent spoofing source chains (e.g. a proof might be valid for an event occurring
     * on another chain, but not for the chain you expect)
     * @return receiptRLP The raw rlp encoded bytes for the receipt, which can further be parsed to extract all logs for
     * the given receipt. This is equivalent to the leaf value stored in the receipt MMPT and is just an RLP encoded
     * object of the type:
     * type ReceiptRLP struct {
     *     PostStateOrStatus []byte
     *     CumulativeGasUsed uint64
     *     Bloom             Bloom
     *     Logs              []*Log
     *     }
     */
    function validateReceipt(bytes calldata proof)
        external
        view
        returns (string memory srcChainId, bytes calldata receiptRLP);

    /**
     * @notice A a log at a given raw rlp encoded receipt at a given logIndex within the receipt.
     * @notice the receiptRLP should first be validated by calling validateReceipt.
     * @param logIndex: The index of the log in the logs array of the receipt. NOTE: This is not the global log index
     * within
     * the block, only the log index within the receipt.
     * @param proof: The proof of a given rlp bytes for the receipt, returned from the receipt MMPT of a block.
     * @return chainId The chainID that the proof proves the log for
     * @return emittingContract The address of the contract that emitted the log on the source chain
     * @return topics The topics of the event. First topic is the event signature that can be calculated by
     * Event.selector. The remaining elements in this array are the indexed parameters of the event.
     * @return unindexedData // The abi encoded non-indexed parameters of the event.
     */
    function validateEvent(uint256 logIndex, bytes calldata proof)
        external
        view
        returns (string memory chainId, address emittingContract, bytes[] calldata topics, bytes calldata unindexedData);

    /**
     * Returns the peptide at a given apphash at a given height,
     */
    function getState(uint256 height) external view returns (uint256);
}
