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
pragma solidity ^0.8.0;

import {IbcDispatcher, IbcEventsEmitter} from "./IbcDispatcher.sol";

import {L1Header, OpL2StateProof, Ics23Proof} from "./ProofVerifier.sol";
import {IbcUniversalChannelMW} from "./IbcMiddleware.sol";
import {
    Channel,
    ChannelEnd,
    ChannelOrder,
    IbcPacket,
    ChannelState,
    AckPacket,
    IBCErrors,
    IbcUtils,
    Ibc
} from "../libs/Ibc.sol";

interface IUniversalChannelHandler is IbcUniversalChannelMW {
    function registerMwStack(uint256 mwBitmap, address[] calldata mwAddrs) external;
    function openChannel(
        string calldata version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        string calldata counterpartyPortIdentifier
    ) external;
    function closeChannel(bytes32 channelId) external;
    function setDispatcher(IbcDispatcher dispatcher) external;
    function dispatcher() external returns (IbcDispatcher dispatcher);
    function connectedChannels(uint256) external view returns (bytes32 channel);
}
