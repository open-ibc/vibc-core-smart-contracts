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

import {ChannelOrder} from "../libs/Ibc.sol";

struct GasFee {
    uint256 gasLimit;
    uint256 gasPrice;
}

struct SendpacketFeeDeposited {
    uint256[2] gasLimits;
    uint256[2] gasPrices;
}

interface IFeeVault {
    event SendPacketFeeDeposited(bytes32 channelId, uint64 sequence, uint256[2] gasLimits, uint256[2] gasPrices);
    event OpenChannelFeeDeposited(
        address sourceAddress,
        string version,
        ChannelOrder ordering,
        string[] connectionHops,
        string counterpartyPortId,
        uint256 feeAmount
    );

    error SenderNotDispatcher();
    error NoFeeSent();
    error IncorrectFeeSent(uint256 expected, uint256 sent);

    function depositSendPacketFee(
        bytes32 channelId,
        uint64 sequence,
        uint256[2] calldata gasLimits,
        uint256[2] calldata gasPrices
    ) external payable;

    function depositOpenChannelFee(
        address sender,
        string memory version,
        ChannelOrder ordering,
        string[] calldata connectionHops,
        string memory counterpartyPortId
    ) external payable;

    function withdrawFeesToOwner() external;
}
