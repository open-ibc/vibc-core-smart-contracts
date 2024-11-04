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

import {Mars} from "./Mars.sol";
import {AckPacket, ChannelOrder} from "../libs/Ibc.sol";
// import {IbcReceiverBase, IbcReceiver, IbcPacket} from "../interfaces/IbcReceiver.sol";
import {IbcDispatcher} from "../interfaces/IbcDispatcher.sol";
// import {FeeSender} from "../implementation_templates/FeeSender.sol";

/**
 * @title Mars
 * @notice Mars is a simple IBC receiver contract that receives packets and sends acks.
 * @dev This contract is used for only testing IBC functionality and as an example for dapp developers on how to
 * integrate with the vibc protocol.
 */
contract MarsEvil is Mars {
    constructor(IbcDispatcher _dispatcher) Mars(dispatcher) {}

    /**
     * @notice batch trigger a channelInit to try to grief  
     */
    function griefChannelInitWithCallsBefore(
        string calldata version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        string calldata counterpartyPortId,
        uint256 numCalls
    ) external onlyOwner {
        _triggerChannelInit(version, ordering, feeEnabled, connectionHops, counterpartyPortId);
        for (uint256 i = 0; i < numCalls; i++) {
            // Grief calls before event
            dispatcher.channelOpenInit(version, ordering, feeEnabled, connectionHops, counterpartyPortId);
        }
    }

    /**
     * @notice batch trigger a channelInit to try to grief  
     */
    function griefChannelInitWithCallsAfter(
        string calldata version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        string calldata counterpartyPortId,
        uint256 numCalls
    ) external onlyOwner {
        for (uint256 i = 0; i < numCalls; i++) {
            // Grief calls before event
            dispatcher.channelOpenInit(version, ordering, feeEnabled, connectionHops, counterpartyPortId);
        }
        _triggerChannelInit(version, ordering, feeEnabled, connectionHops, counterpartyPortId);
    }

    function _triggerChannelInit(
        string calldata version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        string calldata counterpartyPortId
    ) internal {
        IbcDispatcher _dispatcher = dispatcher; // cache for gas savings to avoid 2 SLOADS
        _dispatcher.channelOpenInit(version, ordering, feeEnabled, connectionHops, counterpartyPortId);
        _depositOpenChannelFee(_dispatcher, version, ordering, connectionHops, counterpartyPortId);
    }

}