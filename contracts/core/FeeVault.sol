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

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {
    IFeeVault,
    SendpacketFeeDeposited,
    ChannelOpenFeeDeposited,
    ChannelCloseFeeDeposited
} from "../interfaces/IFeeVault.sol";

contract FeeVault is Ownable, IFeeVault {
    address public dispatcher;
    mapping(bytes32 => mapping(uint64 => SendpacketFeeDeposited)) private sendPacketFees;
    mapping(address => mapping(string => ChannelOpenFeeDeposited)) private channelOpenFees;
    mapping(bytes32 => ChannelCloseFeeDeposited) private channelCloseFees;

    modifier onlyDispatcher() {
        if (msg.sender != dispatcher) revert SenderNotDispatcher();
        _;
    }

    constructor(address _dispatcher) {
        dispatcher = _dispatcher;
    }

    function depositSendPacketFee(bytes32 channelId, uint64 sequence, uint256[2] calldata gasLimits)
        external
        payable
        onlyDispatcher
    {
        // if ((gasLimits[0] * gasPrices[0] + gasLimits[1] * gasPrices[1]) != msg.value) revert IncorrectFeeSent();
        // sendPacketFees[channelId][sequence] = new SendpacketFeeDeposited(gasLimits, gasPrices);
        // emit sendPacketFeeDeposited(msg.sender, gasLimitsgasPrices);
    }

    function depositOpenChannelFee(
        address sender,
        string memory destPortId,
        uint256[3] calldata gasLimits,
        uint256[3] calldata gasPrices
    ) external payable onlyDispatcher {
        uint256 fee = gasLimits[0] * gasPrices[0] + gasLimits[1] * gasPrices[1] + gasLimits[2] * gasPrices[2];
        if (fee != msg.value) revert IncorrectFeeSent();
        emit openChannelFeeDeposited(msg.sender, gasLimits, gasPrices);
    }

    function depositCloseChannelFee(bytes32 channelId, uint256[1] calldata gasLimits, uint256[1] calldata gasPrices)
        external
        payable
        onlyDispatcher
    {
        uint256 fee = gasLimits[0] * gasPrices[0];
        if (fee != msg.value) revert IncorrectFeeSent();

        emit closeChannelFeeDeposited(msg.sender, gasLimits, gasPrices);
    }

    function withdraw() external onlyOwner {
        payable(owner()).transfer(address(this).balance);
    }
}
