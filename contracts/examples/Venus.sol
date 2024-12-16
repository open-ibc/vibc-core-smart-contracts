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
import {Bytes} from "optimism/libraries/Bytes.sol";

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

        // Now that we have validated the receipt, we can trust the rlp encoded receipt bytes. Now we unpack the event
        // data from these rlp encoded receipt bytes and validate it.
        (bytes32 proofChainId, address emittingContract, bytes[] memory topics, bytes memory unindexedData) =
            prover.validateEvent(logIndex, proof);

        if (chainId != proofChainId) {
            revert invalidChainId();
        }
        if (emittingContract != expectedEmitter) {
            revert invalidEventSender();
        }

        if (!prover.validateEvent(receiptIndex, receiptRLPEncodedBytes, logIndex, logBytes, proof)) {
            revert invalidEventProof();
        }

        emit ValidCounterpartyEvent(emittingContract, topics, expectedUnindexedData);
    }

    /**
     * * @notice  Receive a specific houston transmission method. This is a useful example to illustrate how to
     * deconstruct the event type and arguments within the evm layer
     * @param logIndex The index within the tx of where the hello earth is logged in the transaction. Note: this is
     * different than the index of the log within the block
     * @param proof The proof to validate - returned by polymer proof api. This contains the proof to fetch a rlp
     * encoded byte at a given index.
     */
    function receiveTransmissionEvent(uint256 logIndex, bytes calldata proof) external {
        // First, we validate the proof and log in one go, but have to validate the counterparty chain id.
        (bytes32 proofChainId, address emittingContract, bytes[] memory topics, bytes memory unindexedData) =
            prover.validateEvent(logIndex, proof);

        // Once we validate the chain id, we can Now we unpack the event
        if (chainId != proofChainId) {
            revert invalidChainId();
        }

        if (emittingContract != counterParty) {
            // If this triggers, we have received a valid event from the source chain with a valid proof,
            // but it was emitted from a wrong address. This would likely be someone trying to spoof another contract's
            // event.
            // This validation is important since any contract can emit any event.
            revert invalidEventSender();
        }

        // First bytes of topics is always the 32-byte selector calculated by hash of the log string. We want to make
        // sure it is the right type of event before we try to decode the data.
        if (bytes32(topics[0]) != TransmitToHouston.selector) {
            revert invalidCounterpartyEvent();
        }

        // Now that we have verified the event type and the sender is correct, we can trust the log itself occured from
        // a smart contract we are familiar with. All that is left to do is to unpack the args and use them.

        // The abi encoded string is the only indexed arg, it will be the second arg in the topics, we can store it in
        // this local contract
        bytes32 transmission = bytes32(topics[2]);
        lastReceivedTransmission = transmission;

        // We get the unindexed arguments from the generic data property of the log
        uint64 transmissionTime = abi.decode(unindexedData, (uint64));

        emit TransmissionReceived(transmission, transmissionTime);
    }
}
