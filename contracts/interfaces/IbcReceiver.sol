//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IbcDispatcher} from "./IbcDispatcher.sol";
import {ChannelOrder, ChannelEnd, IbcPacket, AckPacket} from "../libs/Ibc.sol";

/**
 * @title IbcChannelReceiver
 * @dev This interface must be implemented by IBC-enabled contracts that act as channel owners and process channel
 * handshake callbacks.
 */
interface IbcChannelReceiver {
    function onChanOpenInit(
        ChannelOrder order,
        string[] calldata connectionHops,
        string calldata counterpartyPortIdentifier,
        string calldata version
    ) external returns (string memory selectedVersion);

    function onChanOpenTry(
        ChannelOrder order,
        string[] memory connectionHops,
        bytes32 channelId,
        string memory counterpartyPortIdentifier,
        bytes32 counterpartychannelId,
        string memory counterpartyVersion
    ) external returns (string memory selectedVersion);

    function onChanOpenAck(bytes32 channelId, bytes32 counterpartychannelId, string calldata counterpartyVersion)
        external;

    /**
     * @notice Handles the channel close init event
     * @param channelId The unique identifier of the channel
     * @dev Make sure to validate channelId and counterpartyVersion
     * @param counterpartyPortId The unique identifier of the counterparty's port
     * @param counterpartyChannelId The unique identifier of the counterparty's channel
     */
    function onChanCloseInit(bytes32 channelId, string calldata counterpartyPortId, bytes32 counterpartyChannelId)
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

    function onChanOpenConfirm(bytes32 channelId) external;
}

/**
 * @title IbcPacketReceiver
 * @notice Packet handler interface must be implemented by a IBC-enabled contract.
 * @dev Packet handling callback methods are invoked by the IBC dispatcher.
 */
interface IbcPacketReceiver {
    function onRecvPacket(IbcPacket calldata packet) external returns (AckPacket memory ackPacket);

    function onAcknowledgementPacket(IbcPacket calldata packet, AckPacket calldata ack) external;

    function onTimeoutPacket(IbcPacket calldata packet) external;
}

/**
 * @title IbcReceiver
 * @author Polymer Labs
 * @notice IBC receiver interface must be implemented by a IBC-enabled contract.
 * The implementer, aka. dApp devs, should implement channel handshake and packet handling methods.
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
