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
 * integrate with polymer's proof api. And some necessary validation which needs to occur
 */
contract Venus {
    ICrossL2Prover public immutable prover;
    address public immutable counterParty; // The dapp on the counter party chain we wish to prove on this local chain
    bytes32 public lastReceivedTransmission; // Last received arguments from transmitted event
    bytes32 public immutable chainId;

    // The event that we emit on source chain to be proven on this local chain
    event TransmitToHouston(bytes32 message, uint64 timestamp);

    // Event that we emit on this local chain to indicate that we have received an event from the source chain
    event TransmissionReceived(bytes32 message, uint64 timestamp);

    event SuccessfulReceipt(bytes32 srcChainId, bytes receiptRLP);
    event ValidCounterpartyEvent(address counterParty, bytes[] topics, bytes unindexed);

    error invalidProverAddress();
    error invalidReceiptProof();
    error invalidChainId(); // The chain id of the proof is not the expected chain id. It might be a valid event, but is
        // on a different chain than we were expecting.
    error invalidEventProof();
    error invalidEventSender();
    error invalidCounterpartyEvent();

    constructor(ICrossL2Prover _prover, address _counterParty, bytes32 _chainId) {
        if (address(_prover) == address(0)) {
            revert invalidProverAddress();
        }
        prover = _prover;
        counterParty = _counterParty;
        chainId = _chainId;
    }

    /**
     * * @notice Validates a receipt proof using the ICrossL2Prover contract
     * @param proof The proof to validate
     * @notice emits the src chain and receipt rlp encoded bytes and the receipt index if the proof is valid, reverts
     * otherwise.
     */
    function receiveReceipt(bytes calldata proof) external {
        (bytes32 srcChainId, bytes memory receiptRLP) = prover.validateReceipt(proof);
        emit SuccessfulReceipt(srcChainId, receiptRLP);
    }

    /**
     * * @notice Validates a generic event within a receipt proof using the ICrossL2Prover contract
     * @notice reverts if the
     * * @param proof The proof to validate, from which the rlp encoded bytes are fetched
     */
    function receiveEvent(
        uint256 logIndex,
        bytes calldata proof,
        address expectedEmitter,
        bytes[] memory expectedTopics,
        bytes calldata expectedUnindexedData
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

        if (!Bytes.equal(abi.encode(topics), abi.encode(expectedTopics))) {
            revert invalidCounterpartyEvent();
        }
        if (!Bytes.equal(unindexedData, expectedUnindexedData)) {
            revert invalidCounterpartyEvent();
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
