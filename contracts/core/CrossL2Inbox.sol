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

/**
 * @title CrossL2Inbox
 * @notice Verifies initiating message proofs
 * @author Polymer Labs
 */
contract CrossL2Inbox {
    using RLPReader for RLPReader.RLPItem;
    using RLPReader for bytes;

    /// @notice Validates a cross chain message on the destination chain
    ///         and emits an ExecutingMessage event. This function is useful
    ///         for applications that understand the schema of the _message payload and want to
    ///         process it in a custom way.
    /// @param _root receipt root.
    /// @param _key receipt index.
    /// @param _value raw receipt.
    /// @param _receiptProof receipt proof.
    /// @param _logHash Hash of the raw log.
    /// @param _logIdx raw log index.
    function validateMessageWithProof(
        bytes32 _root, 
        bytes calldata _key, 
        bytes calldata _value, 
        bytes[] calldata _receiptProof, 
        bytes32 _logHash,
        uint256 _logIdx
    ) external returns (bool) {
        bool valid = MerkleTrie.verifyInclusionProof(
            _key,
            _value,
            _receiptProof,
            _root
        );

        if (!valid) {
            return valid;
        }

        // The first byte is a RLP encoded receipt type so slice it off.
        RLPReader.RLPItem[] memory items = RLPReader.readList(Bytes.slice(_value, 1, _value.length-1));
        /*
            // RLP encoded receipt has the following structure. Logs are the 4th RLP list item.
            type ReceiptRLP struct {
                    PostStateOrStatus []byte
                    CumulativeGasUsed uint64
                    Bloom             Bloom
                    Logs              []*Log
            }
        */ 
        RLPReader.RLPItem[] memory logs = RLPReader.readList(items[3]);

        // Raw log bytes shoul
        return keccak256(RLPReader.readRawBytes(logs[_logIdx])) == _logHash;
    }
}
