// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import './packet.proto.sol';

/// All hash functions are in sha256, conforming to the ICS
library IbcHash {
    // CommitPacket returns the packet commitment bytes. The commitment consists of:
    // sha256_hash(timeout_timestamp + timeout_height.RevisionNumber + timeout_height.RevisionHeight + sha256_hash(data))
    // from a given packet. This results in a fixed length preimage.
    function hashPacketCommitment(Packet calldata packet) public pure returns (bytes32) {
        bytes memory buf = abi.encodePacked(
            uint64ToBigEndian(packet.timeout_timestamp),
            uint64ToBigEndian(packet.timeout_height.revision_number),
            uint64ToBigEndian(packet.timeout_height.revision_height),
            sha256(packet.data)
        );
        return sha256(buf);
    }

    // convert uint64 to bytes8. EVM encode integers in big-endian natively.
    function uint64ToBigEndian(uint64 num) public pure returns (bytes8) {
        return bytes8(abi.encodePacked(num));
    }
}
