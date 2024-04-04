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

    function triggerChannelInit(
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        string calldata counterpartyPortId
    ) external onlyOwner {
        dispatcher.channelOpenInit(supportedVersions[0], ordering, feeEnabled, connectionHops, counterpartyPortId);
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

    function onChanOpenAck(bytes32 channelId, string calldata counterpartyVersion) external virtual onlyIbcDispatcher {
        _connectChannel(channelId, counterpartyVersion);
    }

    function onChanOpenConfirm(bytes32 channelId, string calldata counterpartyVersion) external onlyIbcDispatcher {
        _connectChannel(channelId, counterpartyVersion);
    }

    function onChanOpenInit(string calldata version)
        external
        view
        virtual
        onlyIbcDispatcher
        returns (string memory selectedVersion)
    {
        return _openChannel(version);
    }

    function onChanOpenTry(string calldata counterpartyVersion)
        external
        view
        virtual
        onlyIbcDispatcher
        returns (string memory selectedVersion)
    {
        return _openChannel(counterpartyVersion);
    }

    function _connectChannel(bytes32 channelId, string calldata counterpartyVersion) private {
        // ensure negotiated version is supported
        for (uint256 i = 0; i < supportedVersions.length; i++) {
            if (keccak256(abi.encodePacked(counterpartyVersion)) == keccak256(abi.encodePacked(supportedVersions[i]))) {
                connectedChannels.push(channelId);
                return;
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
}

/*
 * These contracts are the exact same as Mars, but they revert/panick in different ways.
 * they are used to test that transport error is seperated from errors thrown in onRecvPacket
 */
contract RevertingStringMars is Mars {
    constructor(IbcDispatcher _dispatcher) Mars(_dispatcher) {}

    // solhint-disable-next-line
    function onChanOpenInit(string calldata) external view override onlyIbcDispatcher returns (string memory) {
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
    function onChanOpenAck(bytes32, string calldata) external view override onlyIbcDispatcher {
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
