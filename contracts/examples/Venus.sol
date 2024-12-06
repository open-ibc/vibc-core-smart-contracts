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
 * @title Mars
 * @notice Mars is a simple IBC receiver contract that receives packets and sends acks.
 * @dev This contract is used for only testing IBC functionality and as an example for dapp developers on how to
 * integrate with the vibc protocol.
 */
contract Venus {
    ICrossL2Prover public immutable prover ;

    event SuccessfulReceipt(bytes receiptIndex, bytes receiptRLP);
    event SuccessfulEvent(uint eventIndex, bytes eventBytes);

    error invalidReceiptProof();
    error invalidEventProof();

    constructor(ICrossL2Prover _prover) {
        prover = _prover;
    }

    /**
     * @notice trigger a block for a given
     * @notice If you want polymer to relay txs for you, use triggerChannelInitWithFee instead.
     */
    function checkReceiptProof(bytes calldata receiptIndex, bytes calldata receiptRLPEncodedBytes, bytes calldata proof)
        external
    {
        if (!prover.validateReceipt(receiptIndex, receiptRLPEncodedBytes, proof)) {
            revert invalidReceiptProof();
        }

        emit SuccessfulReceipt(receiptIndex, receiptRLPEncodedBytes);
    }

    function checkEventProof(
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

        if(!prover.validateEvent(receiptIndex, receiptRLPEncodedBytes, logIndex, logBytes, proof)){
            revert invalidEventProof();
        }

        emit SuccessfulEvent(logIndex, logBytes);
    }
}
