// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import 'forge-std/Test.sol';

import '../../contracts/Ibc.sol';

// import '../../contracts/payload/hash.sol';

contract AckContract {
    // getAckPacketBytes returns the bytes representation of the AckPacket
    function getAckPacketBytes(AckPacket calldata ack) external pure returns (bytes memory) {
        return abi.encodePacked(ack.success, ack.data);
    }

    // getIncentivizedAckPacketBytes returns the bytes representation of the IncentivizedAckPacket
    function getIncentivizedAckPacketBytes(IncentivizedAckPacket calldata ack) external pure returns (bytes memory) {
        bytes memory relayerBytes = bytes(ack.relayer);
        return abi.encodePacked(ack.success, uint32(relayerBytes.length), relayerBytes, ack.data);
    }
}

contract AckTest is Test {
    AckContract ackContract = new AckContract();

    bytes ackPayload1 = hex'0161636b2d64617461';
    AckPacket ackPacket1 = AckPacket({success: true, data: bytes('ack-data')});
    bytes32 ackPacket1Hash = hex'863b46bf0508a780d8954b8f22b37e91a76c6427a2dcdff98076f977adda55b7';

    bytes ackPayload2 =
        hex'00000000283030303030303030303030303030303030303030303030303030303030303030303030303030303061636b2d64617461';
    IncentivizedAckPacket ackPacket2 =
        IncentivizedAckPacket({
            success: false,
            relayer: '0000000000000000000000000000000000000000',
            data: bytes('ack-data')
        });
    bytes32 ackPacket2Hash = hex'7d0c7d92ef6e1eb1a3c6689239871e8683b6d7a24a630bb8288227316d12afca';

    function setUp() public {}

    // testAckPacketBytes tests the bytes representation of AckPacket
    function testAckPacketBytes() public {
        bytes memory ackPacketBytes = ackContract.getAckPacketBytes(ackPacket1);
        assertEq(ackPacketBytes, ackPayload1, 'ackPacketBytes should be equal to ackPayload1');
        assertEq(ackPacket1Hash, sha256(ackPacketBytes), 'ackPacket1Hash should be equal to keccak256(ackPacketBytes)');
    }

    // testIncentivizedAckPacketBytes tests the bytes representation of IncentivizedAckPacket
    function testIncentivizedAckPacketBytes() public {
        bytes memory ackPacketBytes = ackContract.getIncentivizedAckPacketBytes(ackPacket2);
        assertEq(ackPacketBytes, ackPayload2, 'ackPacketBytes should be equal to ackPayload2');
        assertEq(ackPacket2Hash, sha256(ackPacketBytes), 'ackPacket2Hash should be equal to keccak256(ackPacketBytes)');
    }
}
