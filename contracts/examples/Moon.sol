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

/**
 * @title Moon
 * @notice Moon is a simple IBC receiver contract that receives packets and sends acks. 
 * For now, it is a copy of Mars.sol, but may be extended in the future.
 * @dev This contract is used for only testing IBC functionality and as an example for dapp developers 
 *  on how to integrate with the vibc protocol. 
 */
contract Moon is IbcReceiverBase, IbcReceiver {
    // received packet as chain B
    IbcPacket[] public recvedPackets;
    // received ack packet as chain A
    AckPacket[] public ackPackets;
    // received timeout packet as chain A
    IbcPacket[] public timeoutPackets;
    bytes32[] public connectedChannels;

    string[] public supportedVersions = ["1.0", "2.0"];

    constructor(IbcDispatcher _dispatcher) IbcReceiverBase(_dispatcher) {}

    function triggerChannelInit(
        string calldata version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        string calldata counterpartyPortId
    ) external onlyOwner {
        dispatcher.channelOpenInit(version, ordering, feeEnabled, connectionHops, counterpartyPortId);
    }

    function onRecvPacket(IbcPacket memory packet)
        external
        virtual
        onlyIbcDispatcher
        returns (AckPacket memory ackPacket)
    {
        recvedPackets.push(packet);

        // solhint-disable-next-line quotes
        return AckPacket(true, abi.encodePacked('{ "account": "account", "reply": "got the message" }'));
    }

    function onAcknowledgementPacket(IbcPacket calldata, AckPacket calldata ack) external virtual onlyIbcDispatcher {
        ackPackets.push(ack);
    }

    function onTimeoutPacket(IbcPacket calldata packet) external virtual onlyIbcDispatcher {
        timeoutPackets.push(packet);
    }

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
     * This func triggers channel closure from the dApp.
     * Func args can be arbitary, as long as dispatcher.channelCloseInit is invoked propperly.
     */
    function triggerChannelClose(bytes32 channelId) external onlyOwner {
        dispatcher.channelCloseInit(channelId);
    }

    /**
     * @dev Sends a packet with a greeting message over a specified channel.
     * @param message The greeting message to be sent.
     * @param channelId The ID of the channel to send the packet to.
     * @param timeoutTimestamp The timestamp at which the packet will expire if not received.
     */
    function greet(string calldata message, bytes32 channelId, uint64 timeoutTimestamp) external {
        dispatcher.sendPacket(channelId, bytes(message), timeoutTimestamp);
    }

    function onChanOpenInit(ChannelOrder, string[] calldata, string calldata, string calldata version)
        external
        view
        virtual
        onlyIbcDispatcher
        returns (string memory selectedVersion)
    {
        return _openChannel(version);
    }

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

    function onChanOpenAck(bytes32 channelId, bytes32, string calldata counterpartyVersion)
        external
        virtual
        onlyIbcDispatcher
    {
        _connectChannel(channelId, counterpartyVersion);
    }

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

    function _openChannel(string calldata version) private view returns (string memory selectedVersion) {
        for (uint256 i = 0; i < supportedVersions.length; i++) {
            if (keccak256(abi.encodePacked(version)) == keccak256(abi.encodePacked(supportedVersions[i]))) {
                return version;
            }
        }
        revert UnsupportedVersion();
    }

    function onChanCloseInit(bytes32 channelId, string calldata, bytes32) external virtual onlyIbcDispatcher {}
}

/*
 * These contracts are the exact same as Moon, but they revert/panick in different ways.
 * they are used to test that transport error is seperated from errors thrown in onRecvPacket
 */
contract RevertingStringMoon is Moon {
    constructor(IbcDispatcher _dispatcher) Moon(_dispatcher) {}

    // solhint-disable-next-line
    function onChanOpenInit(ChannelOrder, string[] calldata, string calldata, string calldata)
        external
        view
        virtual
        override
        onlyIbcDispatcher
        returns (string memory selectedVersion)
    {
        // solhint-disable-next-line
        require(false, "open ibc channel is reverting");
        return "";
    }

    // solhint-disable-next-line
    function onRecvPacket(IbcPacket memory) external view override onlyIbcDispatcher returns (AckPacket memory ack) {
        // solhint-disable-next-line
        require(false, "on recv packet is reverting");
        ack = AckPacket(false, "");
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

/// Used to only test reverts for close channel (seperate from RevertingStringMoon to avoid reverts on setting up the
/// channel)
contract RevertingStringCloseChannelMoon is Moon {
    constructor(IbcDispatcher _dispatcher) Moon(_dispatcher) {}
    // solhint-disable-next-line

    function onChanCloseInit(bytes32, string calldata, bytes32) external view override onlyIbcDispatcher {
        // solhint-disable-next-line
        require(false, "close ibc channel is reverting");
    }

    function onChanCloseConfirm(bytes32, string calldata, bytes32) external view override onlyIbcDispatcher {
        // solhint-disable-next-line
        require(false, "close ibc channel is reverting");
    }
}

contract RevertingBytesMoon is Moon {
    error OnRecvPacketRevert();
    error OnTimeoutPacket();

    constructor(IbcDispatcher _dispatcher) Moon(_dispatcher) {}

    function onRecvPacket(IbcPacket memory) external view override onlyIbcDispatcher returns (AckPacket memory ack) {
        ack = AckPacket(false, "");
        revert OnRecvPacketRevert();
    }

    function onTimeoutPacket(IbcPacket calldata) external view override onlyIbcDispatcher {
        // solhint-disable-next-line
        revert OnTimeoutPacket();
    }
}

contract RevertingEmptyMoon is Moon {
    constructor(IbcDispatcher _dispatcher) Moon(_dispatcher) {}

    function onRecvPacket(IbcPacket memory) external view override onlyIbcDispatcher returns (AckPacket memory ack) {
        // solhint-disable-next-line
        require(false);
        ack = AckPacket(false, "");
    }
}

contract PanickingMoon is Moon {
    constructor(IbcDispatcher _dispatcher) Moon(_dispatcher) {}

    function onRecvPacket(IbcPacket memory) external view override onlyIbcDispatcher returns (AckPacket memory ack) {
        assert(false);
        ack = AckPacket(false, "");
    }
}
