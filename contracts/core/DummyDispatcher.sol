//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import {IbcEventsEmitter} from "../interfaces/IbcDispatcher.sol";
import {CounterParty, ChannelOrder, Height, AckPacket} from "../libs/Ibc.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";

/**
 * @title DummyDispatcher
 * @author Polymer Labs
 * @notice
 *     Contract callers call this contract to send IBC-like msg,
 *     which can be relayed to a rollup module on the Polymerase chain
 */
contract DummyDispatcher is IbcEventsEmitter {
    constructor() {}

    function openIbcChannel(
        address portAddress,
        string calldata version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        CounterParty calldata counterparty
    ) external {
        emit OpenIbcChannel(
            portAddress,
            version,
            ordering,
            feeEnabled,
            connectionHops,
            counterparty.portId,
            counterparty.channelId
        );
    }

    function connectIbcChannel(
        address portAddress,
        bytes32 channelId
    ) external {
        emit ConnectIbcChannel(portAddress, channelId);
    }

    function closeIbcChannel(bytes32 channelId) external {
        emit CloseIbcChannel(msg.sender, channelId);
    }

    function sendPacket(
        bytes32 channelId,
	bytes calldata packet,
	uint64 timeoutTimestamp,
	uint64 sequence
    ) external {
        emit SendPacket(msg.sender, channelId, packet, sequence, timeoutTimestamp);
    }

    function acknowledgement(
        address receiver,
        bytes32 channelId,
        uint64 sequence
    ) external {
        emit Acknowledgement(receiver, channelId, sequence);
    }

    function timeout(
        address receiver,
        bytes32 channelId,
        uint64 sequence
    ) external {
        emit Timeout(receiver, channelId, sequence);
    }

    function recvPacket(
        address receiver,
        bytes32 channelId,
        uint64 sequence
    ) external {
        emit RecvPacket(receiver, channelId, sequence);
    }

    function ackPacket(
        address receiver,
        bytes32 channelId,
        uint64 sequence,
        AckPacket calldata ack
    ) external {
        emit WriteAckPacket(receiver, channelId, sequence, ack);
    }

    function writeTimeoutPacket(
        address receiver,
        bytes32 channelId,
        uint64 sequence,
        Height calldata timeoutHeight,
        uint64 timeoutTimestamp
    ) external {
        emit WriteTimeoutPacket(
            receiver, channelId, sequence, timeoutHeight, timeoutTimestamp
        );
    }
}
