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
//import {RLPWriter} from "optimism/libraries/rlp/RLPWriter.sol";

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
    /// @param _receiptProof      receipt proof.
    /// @param _rawLog Raw log or message payload to call target with.
    function validateMessageWithProof(bytes[] calldata _receiptProof, bytes calldata _rawLog) external returns (bool) {
        if (!containsBytes(_receiptProof[_receiptProof.length-1], _rawLog)) {
            return false;
        }

        bytes32 currentHash = keccak256(_receiptProof[_receiptProof.length-1]);
        // Check leaf -> root relationship
        for (int idx = int256(_receiptProof.length)-2; idx >= 0; idx--) {
            uint256 i = uint256(idx);
            if (containsBytes32(_receiptProof[i], currentHash)) {
                currentHash = keccak256(_receiptProof[i]);
                continue;
            }
            return false;
        }

        return true;
    }

    // Contains function for `bytes` arrays
    function containsBytes32(bytes calldata left, bytes32 right) internal pure returns (bool) {
        if (right.length > left.length) {
            return false;
        }
        for (uint i = 0; i <= left.length - right.length; i++) {
            bool found = true;
            for (uint j = 0; j < right.length; j++) {
                if (left[i + j] != right[j]) {
                    found = false;
                    break;
                }
            }
            if (found) {
                return true;
            }
        }
        return false;
    }

    function containsBytes(bytes calldata left, bytes calldata right) internal pure returns (bool) {
        if (right.length > left.length) {
            return false;
        }
        for (uint i = 0; i <= left.length - right.length; i++) {
            bool found = true;
            for (uint j = 0; j < right.length; j++) {
                if (left[i + j] != right[j]) {
                    found = false;
                    break;
                }
            }
            if (found) {
                return true;
            }
        }
        return false;
    }
}
