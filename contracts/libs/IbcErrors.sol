// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.15;

library IBCErrors {
    // Misc errors
    error invalidCounterParty();
    error invalidCounterPartyPortId();
    error invalidHexStringLength();
    error invalidRelayerAddress();
    error consensusStateVerificationFailed();
    error packetNotTimedOut();
    error invalidAddress();
    error invalidPortPrefix();
    error notEnoughGas();
    error invalidCharacter();

    // packet sequence related errors.
    error invalidPacketSequence();
    error unexpectedPacketSequence();

    // channel related errors.
    error channelNotOwnedBySender();
    error channelNotOwnedByPortAddress();

    // client related errors.
    error clientAlreadyCreated();
    error clientNotCreated();

    // packet commitment related errors.
    error packetCommitmentNotFound();
    error ackPacketCommitmentAlreadyExists();
    error packetReceiptAlreadyExists();

    // receiver related errors.
    error receiverNotIntendedPacketDestination();
    error receiverNotOriginPacketSender();

    error invalidChannelType(string channelType);

    // related to clients
    error lightClientNotFound(string connection);
    error channelIdNotFound(bytes32 channelId);
    error invalidConnection(string connection);
}
