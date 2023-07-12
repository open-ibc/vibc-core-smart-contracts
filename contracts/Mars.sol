//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import '@openzeppelin/contracts/access/Ownable.sol';
import './IbcReceiver.sol';
import './IbcDispatcher.sol';

contract Mars is IbcReceiver, Ownable {
    IbcPacket[] public recvedPackets;
    IbcPacket[] public ackPackets;
    IbcPacket[] public timeoutPackets;
    bytes32[] public connectedChannels;

    string[] supportedVersions = ['1.0', '2.0'];
    bytes32[] supportedVersionsBytes = [bytes32('1.0'), bytes32('2.0')]; // Temp solution, will be removed.

    function onRecvPacket(IbcPacket calldata packet) external returns (AckPacket memory ackPacket) {
        recvedPackets.push(packet);
        return AckPacket(true, abi.encodePacked('{ "account": "account", "reply": "got the message" }'));
    }

    function onAcknowledgementPacket(IbcPacket calldata packet) external {
        ackPackets.push(packet);
    }

    function onTimeoutPacket(IbcPacket calldata packet) external {
        timeoutPackets.push(packet);
    }

    function onOpenIbcChannel(
        string calldata version,
        ChannelOrder ordering,
        string[] calldata connectionHops,
        string calldata counterpartyPortId,
        string calldata counterpartyChannelId,
        string calldata counterpartyVersion
    ) external returns (string memory selectedVersion) {
        require(bytes(counterpartyPortId).length > 8, 'Invalid counterpartyPortId');
        /**
         * Version selection is determined by if the callback is invoked on behalf of ChanOpenInit or ChanOpenTry.
         * ChanOpenInit: self version should be provided whereas the counterparty version is empty.
         * ChanOpenTry: counterparty version should be provided whereas the self version is empty.
         * In both cases, the selected version should be in the supported versions list.
         */
        bool foundVersion = false;
        selectedVersion = keccak256(abi.encodePacked(version)) == keccak256(abi.encodePacked('')) ? counterpartyVersion : version;
        for (uint i = 0; i < supportedVersions.length; i++) {
            if (keccak256(abi.encodePacked(selectedVersion)) == keccak256(abi.encodePacked(supportedVersions[i]))) {
                foundVersion = true;
                break;
            }
        }
        require(foundVersion, 'Unsupported version');

        return selectedVersion;
    }

    function onConnectIbcChannel(
        bytes32 channelId,
        bytes32 counterpartyChannelId,
        bytes32 counterpartyVersion
    ) external {
        // ensure negotiated version is supported
        bool foundVersion = false;
        for (uint i = 0; i < supportedVersions.length; i++) {
            if (counterpartyVersion == supportedVersionsBytes[i]) {
                foundVersion = true;
                break;
            }
        }
        require(foundVersion, 'Unsupported version');
        connectedChannels.push(channelId);
    }

    function onCloseIbcChannel(
        bytes32 channelId,
        string calldata counterpartyPortId,
        bytes32 counterpartyChannelId
    ) external {
        // logic to determin if the channel should be closed
        bool channelFound = false;
        for (uint i = 0; i < connectedChannels.length; i++) {
            if (connectedChannels[i] == channelId) {
                delete connectedChannels[i];
                channelFound = true;
                break;
            }
        }
        require(channelFound, 'Channel not found');
    }

    /**
     * This func triggers channel closure from the dApp.
     * Func args can be arbitary, as long as dispatcher.closeIbcChannel is invoked propperly.
     */
    function triggerChannelClose(bytes32 channelId, IbcDispatcher dispatcher) external onlyOwner {
        dispatcher.closeIbcChannel(channelId);
    }

    function greet(
        IbcDispatcher dispatcher,
        string calldata message,
        bytes32 channelId,
        uint64 timeoutTimestamp,
        uint256 fee
    ) external payable {
        dispatcher.sendPacket{value: fee}(channelId, bytes(message), timeoutTimestamp, fee);
    }
}
