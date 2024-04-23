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

import {Height, ChannelEnd, ChannelOrder, AckPacket} from "../libs/Ibc.sol";
import {IbcChannelReceiver} from "./IbcReceiver.sol";
import {Ics23Proof} from "./ProofVerifier.sol";

/**
 * @title IbcPacketSender
 * @author Polymer Labs
 * @dev IBC packet sender interface.
 */
interface IbcPacketSender {
    function sendPacket(bytes32 channelId, bytes calldata payload, uint64 timeoutTimestamp) external;
}

/**
 * @title IbcDispatcher
 * @author Polymer Labs
 * @notice IBC dispatcher interface is the Polymer Core Smart Contract that implements the core IBC protocol.
 * @dev IBC-compatible contracts depend on this interface to actively participate in the IBC protocol.
 *         Other features are implemented as callback methods in the IbcReceiver interface.
 */
interface IbcDispatcher is IbcPacketSender {
    function channelOpenInit(
        string calldata version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        string calldata counterpartyPortId
    ) external;

    function channelCloseConfirm(address portAddress, bytes32 channelId, Ics23Proof calldata proof) external;
    function channelCloseInit(bytes32 channelId) external;

    function portPrefix() external view returns (string memory portPrefix);
}

/**
 * @title IbcEventsEmitter
 * @notice IBC CoreSC events interface.
 */
interface IbcEventsEmitter {
    //
    // channel events
    //
    event ChannelOpenInit(
        address indexed recevier,
        string version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] connectionHops,
        string counterpartyPortId
    );
    event ChannelOpenInitError(address indexed receiver, bytes error);

    event ChannelOpenTry(
        address indexed receiver,
        string version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] connectionHops,
        string counterpartyPortId,
        bytes32 counterpartyChannelId
    );
    event ChannelOpenTryError(address indexed receiver, bytes error);

    event ChannelOpenAck(address indexed receiver, bytes32 channelId);
    event ChannelOpenAckError(address indexed receiver, bytes error);

    event ChannelOpenConfirm(address indexed receiver, bytes32 channelId);
    event ChannelOpenConfirmError(address indexed receiver, bytes error);

    event CloseIbcChannel(address indexed portAddress, bytes32 indexed channelId);

    event CloseIbcChannelError(address indexed receiver, bytes error);
    event AcknowledgementError(address indexed receiver, bytes error);
    event TimeoutError(address indexed receiver, bytes error);

    //
    // packet events
    //
    event SendPacket(
        address indexed sourcePortAddress,
        bytes32 indexed sourceChannelId,
        bytes packet,
        uint64 sequence,
        // timeoutTimestamp is in UNIX nano seconds; packet will be rejected if
        // delivered after this timestamp on the receiving chain.
        // Timeout semantics is compliant to IBC spec and ibc-go implementation
        uint64 timeoutTimestamp
    );

    event Acknowledgement(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 sequence);

    event Timeout(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 indexed sequence);

    event RecvPacket(address indexed destPortAddress, bytes32 indexed destChannelId, uint64 sequence);

    event WriteAckPacket(
        address indexed writerPortAddress, bytes32 indexed writerChannelId, uint64 sequence, AckPacket ackPacket
    );

    event WriteTimeoutPacket(
        address indexed writerPortAddress,
        bytes32 indexed writerChannelId,
        uint64 sequence,
        Height timeoutHeight,
        uint64 timeoutTimestamp
    );
}
