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

    // The event that we emit on source chain to be proven on this local chain
    event TransmitToHouston(bytes32 message, uint64 timestamp);

    // Event that we emit on this local chain to indicate that we have received an event from the source chain
    event TransmissionReceived(bytes32 message, uint64 timestamp);

    event SuccessfulReceipt(bytes receiptIndex, bytes receiptRLP);
    event SuccessfulEvent(uint256 eventIndex, address sender);

    error invalidProverAddress();
    error invalidReceiptProof();
    error invalidEventProof();
    error invalidEventSender();
    error invalidCounterpartyEvent();

    constructor(ICrossL2Prover _prover, address _counterParty) {
        if (address(_prover) == address(0)) {
            revert invalidProverAddress();
        }
        prover = _prover;
        counterParty = _counterParty;
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

    /**
     * * @notice  Receive a specific houston transmission method. This is a useful example to illustrate how to
     * deconstruct the event type and arguments within the evm layer
     * @param receiptIndex The index of the receipt
     * @param receiptRLPEncodedBytes The RLP encoded receipt
     * @param logIndex The index within the tx of where the hello earth is logged in the transaction. Note: this is
     * different than the index of the log within the block
     * @param proof The proof to validate - returned by polymer proof api
     */
    function receiveTransmissionEvent(
        bytes calldata receiptIndex,
        bytes calldata receiptRLPEncodedBytes,
        uint256 logIndex,
        bytes calldata proof
    ) external {
        // First, we validate that the receipt itself is in the receipt MMPT of the counter party chain's block
        if (!prover.validateReceipt(receiptIndex, receiptRLPEncodedBytes, proof)) {
            revert invalidReceiptProof();
        }

        // Now that we have validated the receipt, we can trust the rlp encoded receipt bytes. Now we unpack the event
        // data from these rlp encoded receipt bytes and validate it.

        // The first byte is a RLP encoded receipt type so slice it off.
        RLPReader.RLPItem[] memory receipt =
            RLPReader.readList(Bytes.slice(receiptRLPEncodedBytes, 1, receiptRLPEncodedBytes.length - 1));

        // Logs are the 4th element in the receipt rlp array so we access it directly
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

        // Each log itself is an rlp encoded datatype of 3 properties:
        // type Log struct {
        //         senderAddress bytes // contract address where this log was emitted from
        //         topics bytes        // Array of indexed topics. The first element is the 32-byte selector of the
        // event (can use TransmitToHouston.selector), and the following  elements in this array are the abi encoded
        // arguments individually
        //         topics data         // abi encoded raw bytes of unindexed data
        // }
        RLPReader.RLPItem[] memory log = RLPReader.readList(logs[logIndex]);

        bytes memory emittingAddress = RLPReader.readBytes(log[0]);
        RLPReader.RLPItem[] memory topics = RLPReader.readList(log[1]);
        bytes memory unindexed = (RLPReader.readBytes(log[2])); // This is the raw unindexed data. in this case it's
            // just an abi encoded uint64

        if (address(uint160(uint256(bytes32(emittingAddress)))) != counterParty) {
            // If this triggers, we have received a valid event from the source chain with a valid proof,
            // but it was emitted from a wrong address. This would likely be someone trying to spoof another contract's
            // event.
            // This validation is important since any contract can emit any event.
            revert invalidEventSender();
        }

        // First bytes of topics is always the 32-byte selector calculated by hash of the log string. We want to make
        // sure it is the right type of event before we try to decode the data.
        if (bytes32(RLPReader.readBytes(topics[0])) != TransmitToHouston.selector) {
            revert invalidCounterpartyEvent();
        }

        // Now that we have verified the event type and the sender is correct, we can trust the log itself occured from
        // a smart contract we are familiar with. All that is left to do is to unpack the args and use them.

        // The abi encoded string is the only indexed arg, it will be the second arg in the topics, we can store it in
        // this local contract
        bytes32 transmission = bytes32(RLPReader.readBytes(topics[2]));

        // We get the unindexed arguments from the generic data property of the log
        uint64 transmissionTime = abi.decode(RLPReader.readBytes(unindexed), (uint64));

        lastReceivedTransmission = transmission;
        emit TransmissionReceived(transmission, transmissionTime);
    }
}
