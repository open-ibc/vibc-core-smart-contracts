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

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IbcDispatcher} from "./IbcDispatcher.sol";
import {ChannelOrder, IbcPacket, AckPacket} from "../libs/Ibc.sol";

/**
 * @title IbcChannelReceiver
 * @dev This interface must be implemented by IBC-enabled contracts that act as channel owners and process channel
 * handshake callbacks.
 */
interface IbcChannelReceiver {
    /**
     * @notice Handles the channel open try event (step 2 of the open channel handshake)
     * @dev Make sure to validate that the counterparty version is indeed one supported by the dapp; this callback
     * should
     * revert if not.
     * @param counterpartyVersion The version string provided by the counterparty
     * @return selectedVersion The selected version string
     */
    function onChanOpenTry(
        ChannelOrder order,
        string[] memory connectionHops,
        bytes32 channelId,
        string memory counterpartyPortIdentifier,
        bytes32 counterpartychannelId,
        string memory counterpartyVersion
    ) external returns (string memory selectedVersion);

    /**
     * @notice Handles the channel open acknowledgment event (step 3 of the open channel handshake)
     * @dev Make sure to validate channelId and counterpartyVersion
     * @param channelId The unique identifier of the channel
     * @param counterpartyVersion The version string provided by the counterparty
     */
    function onChanOpenAck(bytes32 channelId, bytes32 counterpartychannelId, string calldata counterpartyVersion)
        external;

    /**
     * @notice Handles the channel close confirmation event
     * @param channelId The unique identifier of the channel
     * @dev Make sure to validate channelId and counterpartyVersion
     * @param counterpartyPortId The unique identifier of the counterparty's port
     * @param counterpartyChannelId The unique identifier of the counterparty's channel
     */
    function onChanCloseConfirm(bytes32 channelId, string calldata counterpartyPortId, bytes32 counterpartyChannelId)
        external;

    /**
     * @notice Handles the channel open confirmation event (step 4 of the open channel handshake)
     * @dev Make sure to validate channelId and counterpartyVersion
     * @param channelId The unique identifier of the channel
     */
    function onChanOpenConfirm(bytes32 channelId) external;
}

/**
 * @title IbcPacketReceiver
 * @notice Packet handler interface must be implemented by a IBC-enabled contract.
 * @dev Packet handling callback methods are invoked by the IBC dispatcher.
 */
interface IbcPacketReceiver {
    /**
     * @notice Callback for receiving a packet; triggered when a counterparty sends an an IBC packet
     * @param packet The IBC packet received
     * @return ackPacket The acknowledgement packet generated in response
     * @dev Make sure to validate packet's source and destiation channels and ports.
     */
    function onRecvPacket(IbcPacket calldata packet) external returns (AckPacket memory ackPacket);

    /**
     * @notice Callback for acknowledging a packet; triggered on reciept of an IBC packet by the counterparty
     * @param packet The IBC packet for which acknowledgement is received
     * @param ack The acknowledgement packet received
     * @dev Make sure to validate packet's source and destiation channels and ports.
     */
    function onAcknowledgementPacket(IbcPacket calldata packet, AckPacket calldata ack) external;

    /**
     * @notice Callback for handling a packet timeout
     * @notice Direct timeouts are currently unimplemented, so this callback is currently unused. Packets can still be
     * indirectly timedout in the recieve callback.
     * @param packet The IBC packet that has timed out
     * @dev Make sure to validate packet's source and destiation channels and ports.
     */
    function onTimeoutPacket(IbcPacket calldata packet) external;
}

/**
 * @title IbcReceiver
 * @author Polymer Labs
 * @notice IBC receiver interface must be implemented by a IBC-enabled contract.
 * The implementer, aka. dApp devs, should implement channel handshake and packet handling methods.
 * @notice : Make sure to integrate carefully. Anyone can open a channel with your dapp. It's important to
 * follow best practices and to note:
 *     - Always validte all given arguments (e.g. channels, ports) in the channel handsakes
 *     - It's recommended to use the onlyIbcDispatcher modifier to restrict access control to only the Dispatcher
 * contract.
 *     - Dispatcher callbacks inherit from nonReentrant, so multiple calls to the Dispatcher cannot be made within the
 * same transaction .
 */
interface IbcReceiver is IbcChannelReceiver, IbcPacketReceiver {}

contract IbcReceiverBase is Ownable {
    IbcDispatcher public dispatcher;

    error notIbcDispatcher();
    error UnsupportedVersion();
    error ChannelNotFound();

    /**
     * @dev Modifier to restrict access to only the IBC dispatcher.
     * Only the address with the IBC_ROLE can execute the function.
     * Should add this modifier to all IBC-related callback functions.
     */
    modifier onlyIbcDispatcher() {
        if (msg.sender != address(dispatcher)) {
            revert notIbcDispatcher();
        }
        _;
    }

    /**
     * @dev Constructor function that takes an IbcDispatcher address and grants the IBC_ROLE to the Polymer IBC
     * Dispatcher.
     * @param _dispatcher The address of the IbcDispatcher contract.
     */
    constructor(IbcDispatcher _dispatcher) Ownable() {
        dispatcher = _dispatcher;
    }

    /// This function is called for plain Ether transfers, i.e. for every call with empty calldata.
    // An empty function body is sufficient to receive packet fee refunds.
    receive() external payable {}
}
