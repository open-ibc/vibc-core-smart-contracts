//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import "./Ibc.sol";
import "./IbcReceiver.sol";
import "./IbcDispatcher.sol";
import "./IbcMiddleware.sol";

contract GeneralMiddleware is IbcMwUser, IbcMiddleware, IbcMwEventsEmitter {
    /**
     * @dev MW_ID is the ID of MW contract on all supported virtual chains.
     * MW_ID must:
     * - be globally unique, ie. no two MWs should have the same MW_ID registered on Polymer chain.
     * - be identical on all supported virtual chains.
     * - be identical on one virtual chain across multiple deployed MW instances. Each MW instance belong exclusively to one MW stack.
     * - be 1 << N, where N is a non-negative integer, and not in conflict with other MWs.
     */
    uint256 public MW_ID;

    /**
     * @param _middleware The middleware contract address this contract sends packets to and receives packets from.
     */

    constructor(uint256 mwId, address _middleware) IbcMwUser(_middleware) {
        MW_ID = mwId;
    }

    function sendUniversalPacket(
        bytes32 channelId,
        address destPortAddr,
        bytes calldata appData,
        uint64 timeoutTimestamp
    ) external override {
        _sendPacket(channelId, msg.sender, destPortAddr, 0, appData, timeoutTimestamp);
    }

    function sendMWPacket(
        bytes32 channelId,
        address srcPortAddr,
        address destPortAddr,
        uint256 srcMwIds,
        bytes calldata appData,
        uint64 timeoutTimestamp
    ) external override {
        _sendPacket(channelId, srcPortAddr, destPortAddr, srcMwIds, appData, timeoutTimestamp);
    }

    function onRecvUniversalPacket(bytes32 channelId, address srcPortId, bytes calldata appData)
        external
        override
        returns (AckPacket memory ackPacket)
    {}

    function onUniversalAcknowledgement(UniversalPacket memory packet, AckPacket calldata ack) external override {}

    function onTimeoutUniversalPacket(UniversalPacket calldata packet) external override {}

    // internal function to send packet to next MW with MW Ids bit flipped
    // param srcMwIds: MW ID bitmap excluding this MW's ID
    function _sendPacket(
        bytes32 channelId,
        address srcPortAddr,
        address destPortAddr,
        uint256 srcMwIds,
        bytes calldata appData,
        uint64 timeoutTimestamp
    ) internal virtual {
        // extra custom logic here to process packet, eg. emit MW events, mutate state, etc.
        // implementer can emit custom data fields suitable for their use cases.
        // Here we use MW_ID as the custom MW data field.
        emit SendMWPacket(
            channelId, srcPortAddr, destPortAddr, MW_ID, appData, timeoutTimestamp, abi.encodePacked(MW_ID)
        );

        // send packet to next MW
        IbcMwPacketSender(mw).sendMWPacket(
            channelId, srcPortAddr, destPortAddr, srcMwIds | MW_ID, appData, timeoutTimestamp
        );
    }
}
