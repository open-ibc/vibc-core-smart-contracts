//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import "./Ibc.sol";
import "./IbcDispatcher.sol";
import "./IbcReceiver.sol";

/**
 * @title IbcUniversalPacketSender
 * @dev IbcUniversalPacketSender is called by end-users of IBC middleware contracts to send a packet over a MW stack.
 */
interface IbcUniversalPacketSender {
    function sendUniversalPacket(bytes32 channelId, address destPortId, bytes calldata appData, uint64 timeoutTimestamp)
        external;
}

interface IbcUniversalPacketReceiver {
    function onRecvUniversalPacket(bytes32 channelId, address srcPortId, bytes calldata appData)
        external
        returns (AckPacket memory ackPacket);

    function onUniversalAcknowledgement(UniversalPacket memory packet, AckPacket calldata ack) external;

    function onTimeoutUniversalPacket(UniversalPacket calldata packet) external;
}

/**
 * @title IbcMiddleware
 * @notice IBC middleware interface must be implemented by a IBC-middleware contract, similar to ICS-30
 *  https://github.com/cosmos/ibc/tree/main/spec/app/ics-030-middleware.
 */
interface IbcMiddleware is IbcUniversalPacketSender, IbcPacketReceiver {
    // sendMWPacket must be called by another authorized IBC middleware contract.
    // NEVER by a dApp contract.
    // The middleware contract may choose to send the packet by calling IbcCoreSC or another IBC MW.
    function sendMWPacket(
        bytes32 channelId,
        // original source address of the packet
        address srcPortAddr,
        address destPortAddr,
        // source middleware ids bit AND
        uint256 srcMwIds,
        bytes calldata appData,
        uint64 timeoutTimestamp
    ) external;
}

/**
 * @title IbcUniversalChannelMW
 * @notice This interface must be implemented by IBC universal channel middleware contracts.
 */
interface IbcUniversalChannelMW is IbcMiddleware, IbcChannelReceiver {}

contract IbcMwReceiverBase is Ownable {
    IbcMiddleware public mw;

    /**
     * @dev Constructor function that takes an IbcMiddleware address and grants the IBC_ROLE to the Polymer IBC Dispatcher.
     * @param _middleware The address of the IbcMiddleware contract.
     */
    constructor(IbcMiddleware _middleware) Ownable() {
        mw = _middleware;
    }

    /// This function is called for plain Ether transfers, i.e. for every call with empty calldata.
    // An empty function body is sufficient to receive packet fee refunds.
    receive() external payable {}

    /**
     * @dev Modifier to restrict access to only the IBC middleware.
     * Only the address with the IBC_ROLE can execute the function.
     * Should add this modifier to all IBC-related callback functions.
     */
    modifier onlyIbcMw() {
        require(msg.sender == address(mw), "only IBC middleware");
        _;
    }
}
