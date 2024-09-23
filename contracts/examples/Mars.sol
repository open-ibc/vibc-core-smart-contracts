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

import {AckPacket, ChannelOrder} from "../libs/Ibc.sol";
import {IbcReceiverBase, IbcReceiver, IbcPacket} from "../interfaces/IbcReceiver.sol";
import {IbcDispatcher} from "../interfaces/IbcDispatcher.sol";
import {FeeSender} from "../implementation_templates/FeeSender.sol";

/**
 * @title Mars
 * @notice Mars is a simple IBC receiver contract that receives packets and sends acks.
 * @dev This contract is used for only testing IBC functionality and as an example for dapp developers on how to
 * integrate with the vibc protocol.
 */
contract Mars is IbcReceiverBase, IbcReceiver, FeeSender {
    // received packet as chain B
    IbcPacket[] public recvedPackets;
    // received ack packet as chain A
    AckPacket[] public ackPackets;
    // received timeout packet as chain A
    IbcPacket[] public timeoutPackets;
    bytes32[] public connectedChannels;

    string[] public supportedVersions = ["1.0", "2.0"];

    constructor(IbcDispatcher _dispatcher) IbcReceiverBase(_dispatcher) {}

    /**
     * @notice trigger a channelInit in the dispatcher with no relayer fees.
     * @notice If you want polymer to relay txs for you, use triggerChannelInitWithFee instead.
     */
    function triggerChannelInit(
        string calldata version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        string calldata counterpartyPortId
    ) external onlyOwner {
        dispatcher.channelOpenInit(version, ordering, feeEnabled, connectionHops, counterpartyPortId);
    }

    /**
     * @notice trigger a channelInit in the dispatcher with an additional call to deposit a fee into the FeeVault
     * @notice This should not be used if you are relaying your own txs, and triggerChannelInit should instead be
     * called.
     */
    function triggerChannelInitWithFee(
        string calldata version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        string calldata counterpartyPortId
    ) external payable onlyOwner {
        IbcDispatcher _dispatcher = dispatcher; // cache for gas savings to avoid 2 SLOADS
        _dispatcher.channelOpenInit(version, ordering, feeEnabled, connectionHops, counterpartyPortId);
        _depositOpenChannelFee(_dispatcher, version, ordering, connectionHops, counterpartyPortId);
    }

    /**
     * @notice Callback for receiving a packet; triggered when a counterparty sends an an IBC packet
     * @param packet The IBC packet received
     * @return ackPacket The acknowledgement packet generated in response
     * @return skipAck Whether to skip the writeAck event.
     * @dev Make sure to validate packet's source and destiation channels and ports.
     */
    function onRecvPacket(IbcPacket memory packet)
        external
        virtual
        onlyIbcDispatcher
        returns (AckPacket memory ackPacket, bool skipAck)
    {
        recvedPackets.push(packet);

        // solhint-disable-next-line quotes
        return (AckPacket(true, abi.encodePacked('{ "account": "account", "reply": "got the message" }')), false);
    }

    /**
     * @notice Callback for acknowledging a packet; triggered on reciept of an IBC packet by the counterparty
     * @dev Make sure to validate packet's source and destiation channels and ports.
     */
    function onAcknowledgementPacket(IbcPacket calldata, AckPacket calldata ack) external virtual onlyIbcDispatcher {
        ackPackets.push(ack);
    }

    /**
     * @notice Callback for handling a packet timeout
     * @notice Direct timeouts are currently unimplemented, so this callback is currently unused. Packets can still be
     * indirectly timedout in the recieve callback.
     * @param packet The IBC packet that has timed out
     * @dev Make sure to validate packet's source and destiation channels and ports.
     */
    function onTimeoutPacket(IbcPacket calldata packet) external virtual onlyIbcDispatcher {
        timeoutPackets.push(packet);
    }

    function onChanCloseInit(bytes32 channelId, string calldata, bytes32) external virtual onlyIbcDispatcher {}

    /**
     * @notice Handles channel close callback on the dest chain
     * @param channelId The unique identifier of the channel
     * @dev Make sure to validate channelId and counterpartyVersion
     */
    function onChanCloseConfirm(bytes32 channelId, string calldata, bytes32) external virtual onlyIbcDispatcher {
        // logic to determine if the channel should be closed
        bool channelFound = false;
        for (uint256 i = 0; i < connectedChannels.length; i++) {
            if (connectedChannels[i] == channelId) {
                delete connectedChannels[i];
                channelFound = true;
                break;
            }
        }
        if (!channelFound) revert ChannelNotFound();
    }

    /**
     * Trigger a channel close handhake
     */
    function triggerChannelClose(bytes32 channelId) external onlyOwner {
        dispatcher.channelCloseInit(channelId);
    }

    /**
     * @dev Sends a packet with a greeting message over a specified channel, without depositing any sendPacket relaying
     * fees.
     * @notice Use greetWithFee for sending packets with fees.
     * @param message The greeting message to be sent.
     * @param channelId The ID of the channel to send the packet to.
     * @param timeoutTimestamp The timestamp at which the packet will expire if not received.
     * @dev This method also returns sequence from the dispatcher for easy testing
     */
    function greet(string calldata message, bytes32 channelId, uint64 timeoutTimestamp)
        external
        returns (uint64 sequence)
    {
        sequence = dispatcher.sendPacket(channelId, bytes(message), timeoutTimestamp);
    }

    /**
     * @dev Sends a packet with a greeting message over a specified channel, and deposits a fee for relaying the packet
     * @param message The greeting message to be sent.
     * @param channelId The ID of the channel to send the packet to.
     * @param timeoutTimestamp The timestamp at which the packet will expire if not received.
     * @notice If you are relaying your own packets, you should not call this method, and instead call greet.
     * @param gasLimits An array containing two gas limit values:
     *                  - gasLimits[0] for `recvPacket` fees
     *                  - gasLimits[1] for `ackPacket` fees.
     * @param gasPrices An array containing two gas price values:
     *                  - gasPrices[0] for `recvPacket` fees, for the dest chain
     *                  - gasPrices[1] for `ackPacket` fees, for the src chain
     * @notice The total fees sent in the msg.value should be equal to the total gasLimits[0] * gasPrices[0] +
     * @notice Use the Polymer fee estimation api to get the required fees to ensure that enough fees are sent.
     * gasLimits[1] * gasPrices[1]. The transaction will revert if a higher or lower value is sent
     */
    function greetWithFee(
        string calldata message,
        bytes32 channelId,
        uint64 timeoutTimestamp,
        uint256[2] memory gasLimits,
        uint256[2] memory gasPrices
    ) external payable returns (uint64 sequence) {
        sequence = dispatcher.sendPacket(channelId, bytes(message), timeoutTimestamp);
        _depositSendPacketFee(dispatcher, channelId, sequence, gasLimits, gasPrices);
    }

    /**
     * @notice Handles the channel open try event (step 2 of the open channel handshake)
     * @dev Make sure to validate that the counterparty version is indeed one supported by the dapp; this callback
     * should
     * revert if not.
     * @param counterpartyVersion The version string provided by the counterparty
     * @return selectedVersion The selected version string
     */
    // solhint-disable-next-line ordering
    function onChanOpenTry(
        ChannelOrder,
        string[] memory,
        bytes32 channelId,
        string memory,
        bytes32,
        string calldata counterpartyVersion
    ) external virtual onlyIbcDispatcher returns (string memory selectedVersion) {
        return _connectChannel(channelId, counterpartyVersion);
    }

    /**
     * @notice Handles the channel open acknowledgment event (step 3 of the open channel handshake)
     * @dev Make sure to validate channelId and counterpartyVersion
     * @param channelId The unique identifier of the channel
     * @param counterpartyVersion The version string provided by the counterparty
     */
    function onChanOpenAck(bytes32 channelId, bytes32, string calldata counterpartyVersion)
        external
        virtual
        onlyIbcDispatcher
    {
        _connectChannel(channelId, counterpartyVersion);
    }

    /**
     * @notice Handles the channel open confirmation event (step 4 of the open channel handshake)
     * @dev Make sure to validate channelId and counterpartyVersion
     * @param channelId The unique identifier of the channel
     */
    function onChanOpenConfirm(bytes32 channelId) external onlyIbcDispatcher {}

    function _connectChannel(bytes32 channelId, string calldata counterpartyVersion)
        private
        returns (string memory version)
    {
        // ensure negotiated version is supported
        for (uint256 i = 0; i < supportedVersions.length; i++) {
            if (keccak256(abi.encodePacked(counterpartyVersion)) == keccak256(abi.encodePacked(supportedVersions[i]))) {
                connectedChannels.push(channelId);
                return counterpartyVersion;
            }
        }
        revert UnsupportedVersion();
    }
}

/*
 * These contracts are the exact same as Mars, but they revert/panick in different ways.
 * they are used to test that transport error is seperated from errors thrown in onRecvPacket
 */
contract RevertingStringMars is Mars {
    constructor(IbcDispatcher _dispatcher) Mars(_dispatcher) {}

    // solhint-disable-next-line
    function onRecvPacket(IbcPacket memory)
        external
        view
        override
        onlyIbcDispatcher
        returns (AckPacket memory ack, bool skipAck)
    {
        // solhint-disable-next-line
        require(false, "on recv packet is reverting");
        ack = AckPacket(false, "");
        skipAck = false;
    }

    // solhint-disable-next-line
    function onChanOpenAck(bytes32, bytes32, string calldata) external view override onlyIbcDispatcher {
        // solhint-disable-next-line
        require(false, "connect ibc channel is reverting");
    }

    // solhint-disable-next-line
    function onAcknowledgementPacket(IbcPacket calldata, AckPacket calldata) external view override onlyIbcDispatcher {
        // solhint-disable-next-line
        require(false, "acknowledgement packet is reverting");
    }
}

/// Used to only test reverts for close channel (seperate from RevertingStringMars to avoid reverts on setting up the
/// channel)
contract RevertingStringCloseChannelMars is Mars {
    constructor(IbcDispatcher _dispatcher) Mars(_dispatcher) {}
    // solhint-disable-next-line

    function onChanCloseConfirm(bytes32, string calldata, bytes32) external view override onlyIbcDispatcher {
        // solhint-disable-next-line
        require(false, "close ibc channel is reverting");
    }
}

contract RevertingBytesMars is Mars {
    error OnRecvPacketRevert();
    error OnTimeoutPacket();

    constructor(IbcDispatcher _dispatcher) Mars(_dispatcher) {}

    function onRecvPacket(IbcPacket memory)
        external
        view
        override
        onlyIbcDispatcher
        returns (AckPacket memory ack, bool skipAck)
    {
        ack = AckPacket(false, "");
        skipAck = false;
        revert OnRecvPacketRevert();
    }

    function onTimeoutPacket(IbcPacket calldata) external view override onlyIbcDispatcher {
        // solhint-disable-next-line
        revert OnTimeoutPacket();
    }
}

contract RevertingEmptyMars is Mars {
    constructor(IbcDispatcher _dispatcher) Mars(_dispatcher) {}

    function onRecvPacket(IbcPacket memory)
        external
        view
        override
        onlyIbcDispatcher
        returns (AckPacket memory ack, bool skipAck)
    {
        // solhint-disable-next-line
        require(false);
        // Set values to avoid compiler warning
        ack = AckPacket(false, "");
        skipAck = false;
    }
}

contract PanickingMars is Mars {
    constructor(IbcDispatcher _dispatcher) Mars(_dispatcher) {}

    function onRecvPacket(IbcPacket memory)
        external
        view
        override
        onlyIbcDispatcher
        returns (AckPacket memory ack, bool skipAck)
    {
        assert(false);
        ack = AckPacket(false, "");
        skipAck = false;
    }
}
