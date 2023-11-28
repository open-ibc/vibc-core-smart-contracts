//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import '@openzeppelin/contracts/access/Ownable.sol';
import './IbcDispatcher.sol';
import './Ibc.sol';

/**
 * @title IbcReceiver
 * @author Polymer Labs
 * @notice IBC receiver interface must be implemented by a IBC-enabled contract.
 * The implementer, aka. dApp devs, should implement channel handshake and packet handling methods.
 */
interface IbcReceiver {
    //
    // Packet handling methods
    //

    function onRecvPacket(IbcPacket calldata packet) external returns (AckPacket memory ackPacket);

    function onAcknowledgementPacket(IbcPacket calldata packet, AckPacket calldata ack) external;

    function onTimeoutPacket(IbcPacket calldata packet) external;

    //
    // Channel handshake methods
    //

    function onOpenIbcChannel(
        string calldata version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        string calldata counterpartyPortId,
        bytes32 counterpartyChannelId,
        string calldata counterpartyVersion
    ) external returns (string memory selectedVersion);

    function onConnectIbcChannel(
        bytes32 channelId,
        bytes32 counterpartyChannelId,
        string calldata counterpartyVersion
    ) external;

    function onCloseIbcChannel(
        bytes32 channelId,
        string calldata counterpartyPortId,
        bytes32 counterpartyChannelId
    ) external;
}

contract IbcReceiverBase is Ownable {

    IbcDispatcher public dispatcher;

    /**
     * @dev Constructor function that takes an IbcDispatcher address and grants the IBC_ROLE to the Polymer IBC Dispatcher.
     * @param _dispatcher The address of the IbcDispatcher contract.
     */
    constructor(IbcDispatcher _dispatcher) Ownable() {
        dispatcher = _dispatcher;
    }


    /**
     * @dev Modifier to restrict access to only the IBC dispatcher.
     * Only the address with the IBC_ROLE can execute the function.
     * Should add this modifier to all IBC-related callback functions.
     */
    modifier onlyIbcDispatcher() {
        require(msg.sender == address(dispatcher), 'only IBC dispatcher');
        _;
    }
}
