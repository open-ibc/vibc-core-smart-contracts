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

contract GasAudit {
    mapping(bytes32 => bool) public channelIds1;
    mapping(string => bool) public channelIds2;

    event OpenIbcChannel1(
        address indexed portAddress, bytes32 indexed channelId, string counterpartyPortId, bytes32 coutnerpartyChannelId
    );

    event OpenIbcChannel2(
        address indexed portAddress, string channelId, string counterpartyPortId, string coutnerpartyChannelId
    );

    function callWithBytes32(
        address portAddress,
        bytes32 channelId,
        string calldata counterpartyPortId,
        bytes32 coutnerpartyChannelId
    ) external {
        channelIds1[channelId] = true;
        emit OpenIbcChannel1(portAddress, channelId, counterpartyPortId, coutnerpartyChannelId);
    }

    function callWithString(
        address portAddress,
        string calldata channelId,
        string calldata counterpartyPortId,
        string calldata coutnerpartyChannelId
    ) external {
        channelIds2[channelId] = true;
        emit OpenIbcChannel2(portAddress, channelId, counterpartyPortId, coutnerpartyChannelId);
    }
}
