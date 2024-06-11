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

struct SendpacketFeeDeposited {
    uint256[2] gasLimits;
    uint256[2] gasPrices;
}

struct ChannelOpenFeeDeposited {
    uint256[3] gasLimits;
    uint256[3] gasPrices;
}

struct ChannelCloseFeeDeposited {
    uint256[1] gasLimits;
    uint256[1] gasPrices;
}

interface IFeeVault {
    error SenderNotDispatcher();
    error IncorrectFeeSent();

    event sendPacketFeeDeposited(address beneficiary, uint256[2] gasLimits, uint256[2] gasPrices);
    event openChannelFeeDeposited(address beneficiary, uint256[3] gasLimits, uint256[3] gasPrices);
    event closeChannelFeeDeposited(address beneficiary, uint256[1] gasLimits, uint256[1] gasPrices);

    function depositSendPacketFee(bytes32 channelId, uint64 sequence, uint256[2] calldata gasLimits) external payable;

    function depositOpenChannelFee(
        address sender,
        string memory destPortId,
        uint256[3] calldata gasLimits,
        uint256[3] calldata gasPrices
    ) external payable;
    function depositCloseChannelFee(bytes32 channelId, uint256[1] calldata gasLimits, uint256[1] calldata gasPrices)
        external
        payable;
    function withdraw() external;
}
