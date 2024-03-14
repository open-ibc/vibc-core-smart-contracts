//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import {IBCErrors, AckPacket, ChannelOrder, CounterParty} from "../libs/Ibc.sol";
import {IbcReceiverBase, IbcReceiver, IbcPacket} from "../interfaces/IbcReceiver.sol";
import {IbcDispatcher} from "../interfaces/IbcDispatcher.sol";

contract Mars is IbcReceiverBase, IbcReceiver {
    // received packet as chain B
    IbcPacket[] public recvedPackets;
    // received ack packet as chain A
    AckPacket[] public ackPackets;
    // received timeout packet as chain A
    IbcPacket[] public timeoutPackets;
    bytes32[] public connectedChannels;

    string[] public supportedVersions = ["1.0", "2.0"];

    constructor(IbcDispatcher _dispatcher) IbcReceiverBase(_dispatcher) {}

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

    function onConnectIbcChannel(bytes32 channelId, bytes32, string calldata counterpartyVersion)
        external
        virtual
        onlyIbcDispatcher
    {
        // ensure negotiated version is supported
        bool foundVersion = false;
        for (uint256 i = 0; i < supportedVersions.length; i++) {
            if (keccak256(abi.encodePacked(counterpartyVersion)) == keccak256(abi.encodePacked(supportedVersions[i]))) {
                foundVersion = true;
                break;
            }
        }
        if (!foundVersion) revert UnsupportedVersion();
        connectedChannels.push(channelId);
    }

    function onCloseIbcChannel(bytes32 channelId, string calldata, bytes32) external virtual onlyIbcDispatcher {
        // logic to determin if the channel should be closed
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
     * Func args can be arbitary, as long as dispatcher.closeIbcChannel is invoked propperly.
     */
    function triggerChannelClose(bytes32 channelId) external onlyOwner {
        dispatcher.closeIbcChannel(channelId);
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

    function onOpenIbcChannel(
        string calldata version,
        ChannelOrder,
        bool,
        string[] calldata,
        CounterParty calldata counterparty
    ) external view virtual onlyIbcDispatcher returns (string memory selectedVersion) {
        if (bytes(counterparty.portId).length <= 8) {
            revert IBCErrors.invalidCounterPartyPortId();
        }
        /**
         * Version selection is determined by if the callback is invoked on behalf of ChanOpenInit or ChanOpenTry.
         * ChanOpenInit: self version should be provided whereas the counterparty version is empty.
         * ChanOpenTry: counterparty version should be provided whereas the self version is empty.
         * In both cases, the selected version should be in the supported versions list.
         */
        bool foundVersion = false;
        selectedVersion =
            keccak256(abi.encodePacked(version)) == keccak256(abi.encodePacked("")) ? counterparty.version : version;
        for (uint256 i = 0; i < supportedVersions.length; i++) {
            if (keccak256(abi.encodePacked(selectedVersion)) == keccak256(abi.encodePacked(supportedVersions[i]))) {
                foundVersion = true;
                break;
            }
        }
        if (!foundVersion) revert UnsupportedVersion();
        // if counterpartyVersion is not empty, then it must be the same foundVersion
        if (keccak256(abi.encodePacked(counterparty.version)) != keccak256(abi.encodePacked(""))) {
            if (keccak256(abi.encodePacked(counterparty.version)) != keccak256(abi.encodePacked(selectedVersion))) {
                revert VersionMismatch();
            }
        }

        return selectedVersion;
    }
}

/*
 * These contracts are the exact same as Mars, but they revert/panick in different ways.
 * they are used to test that transport error is seperated from errors thrown in onRecvPacket
 */
contract RevertingStringMars is Mars {
    constructor(IbcDispatcher _dispatcher) Mars(_dispatcher) {}

    // solhint-disable-next-line
    function onOpenIbcChannel(string calldata, ChannelOrder, bool, string[] calldata, CounterParty calldata)
        external
        view
        override
        onlyIbcDispatcher
        returns (string memory)
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
    function onConnectIbcChannel(bytes32, bytes32, string calldata) external view override onlyIbcDispatcher {
        // solhint-disable-next-line
        require(false, "connect ibc channel is reverting");
    }

    // solhint-disable-next-line
    function onCloseIbcChannel(bytes32, string calldata, bytes32) external view override onlyIbcDispatcher {
        // solhint-disable-next-line
        require(false, "close ibc channel is reverting");
    }

    // solhint-disable-next-line
    function onAcknowledgementPacket(IbcPacket calldata, AckPacket calldata) external view override onlyIbcDispatcher {
        // solhint-disable-next-line
        require(false, "acknowledgement packet is reverting");
    }
}

contract RevertingBytesMars is Mars {
    error OnRecvPacketRevert();
    error OnTimeoutPacket();

    constructor(IbcDispatcher _dispatcher) Mars(_dispatcher) {}

    function onRecvPacket(IbcPacket memory) external view override onlyIbcDispatcher returns (AckPacket memory ack) {
        ack = AckPacket(false, "");
        revert OnRecvPacketRevert();
    }

    function onTimeoutPacket(IbcPacket calldata) external view override onlyIbcDispatcher {
        // solhint-disable-next-line
        revert OnTimeoutPacket();
    }
}

contract RevertingEmptyMars is Mars {
    constructor(IbcDispatcher _dispatcher) Mars(_dispatcher) {}

    function onRecvPacket(IbcPacket memory) external view override onlyIbcDispatcher returns (AckPacket memory ack) {
        // solhint-disable-next-line
        require(false);
        ack = AckPacket(false, "");
    }
}

contract PanickingMars is Mars {
    constructor(IbcDispatcher _dispatcher) Mars(_dispatcher) {}

    function onRecvPacket(IbcPacket memory) external view override onlyIbcDispatcher returns (AckPacket memory ack) {
        assert(false);
        ack = AckPacket(false, "");
    }
}
