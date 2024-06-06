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

import {IbcDispatcher} from "../interfaces/IbcDispatcher.sol";
import {ChannelOrder} from "../libs/Ibc.sol";

///  Contract with
abstract contract FeeSender {
    function _depositSendPacketFee(
        IbcDispatcher dispatcher,
        bytes32 channelId,
        uint64 sequence,
        uint256[2] calldata gasLimits,
        uint256[2] calldata gasPrices
    ) internal {
        dispatcher.feeVault().depositSendPacketFee{value: msg.value}(channelId, sequence, gasLimits, gasPrices);
    }

    function _depositOpenChannelFee(
        IbcDispatcher dispatcher,
        string memory version,
        ChannelOrder ordering,
        string[] calldata connectionHops,
        string calldata counterpartyPortId
    ) internal {
        dispatcher.feeVault().depositOpenChannelFee{value: msg.value}(
            address(this), version, ordering, connectionHops, counterpartyPortId
        );
    }
}
