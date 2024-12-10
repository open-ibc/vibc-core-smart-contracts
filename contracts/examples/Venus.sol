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

pragma solidity ^0.8.9;

import {ICrossL2Prover} from "../interfaces/ICrossL2Prover.sol";

/**
 * @title Venus
 * @notice Venus is a simple polymer proof api contract that proves an event happened on another chain.
 * @dev This contract is used for only testing and as an example for dapp developers on how to
 * integrate with polymer's proof api.
 */
contract Venus {
    ICrossL2Prover public immutable prover;

    event SuccessfulReceipt(bytes receiptIndex, bytes receiptRLP);
    event SuccessfulEvent(uint256 eventIndex, address sender);
    event SendHealthCheckEvent(bytes32 id, uint256 sourceChainID, uint256 destChainID);

    error invalidProverAddress();
    error invalidReceiptProof();
    error invalidEventProof();

    constructor(ICrossL2Prover _prover) {
        if (address(_prover) == address(0)) {
            revert invalidProverAddress();
        }
        prover = _prover;
    }

    /**
     * * @notice Validates a receipt proof using the ICrossL2Prover contract
     * * @param receiptIndex The index of the receipt
     * * @param receiptRLPEncodedBytes The RLP encoded receipt
     * * @param proof The proof to validate
     */
    function receiveReceipt(bytes calldata receiptIndex, bytes calldata receiptRLPEncodedBytes, bytes calldata proof)
        external
    {
        if (!prover.validateReceipt(receiptIndex, receiptRLPEncodedBytes, proof)) {
            revert invalidReceiptProof();
        }

        emit SuccessfulReceipt(receiptIndex, receiptRLPEncodedBytes);
    }

    /**
     * * @notice Validates an event within a receipt proof using the ICrossL2Prover contract
     * * @param receiptIndex The index of the receipt
     * * @param receiptRLPEncodedBytes The RLP encoded receipt
     * * @param proof The proof to validate
     */
    function receiveEvent(
        bytes calldata receiptIndex,
        bytes calldata receiptRLPEncodedBytes,
        uint256 logIndex,
        bytes calldata logBytes,
        bytes calldata proof
    ) external {
        // First we validate receipt to have a more helpful error message if the receipt itself is incorrect

        if (!prover.validateReceipt(receiptIndex, receiptRLPEncodedBytes, proof)) {
            revert invalidReceiptProof();
        }

        if (!prover.validateEvent(receiptIndex, receiptRLPEncodedBytes, logIndex, logBytes, proof)) {
            revert invalidEventProof();
        }

        emit SuccessfulEvent(logIndex, msg.sender);
    }

    function emitEvent(uint256 sourceChainID, uint256 destinationChainID) external payable {
        bytes32 uniqueID = keccak256(
            abi.encodePacked(sourceChainID, destinationChainID, block.timestamp, block.number)
        );

        emit SendHealthCheckEvent(uniqueID, sourceChain, destinationChain, msg.value);
    }
}
