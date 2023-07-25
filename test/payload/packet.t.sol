// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import 'forge-std/Test.sol';
import '@lazyledger/protobuf3-solidity-lib/contracts/ProtobufLib.sol';
import '../../contracts/payload/packet.proto.sol';

contract PacketContract {
    function decode(
        uint64 initial_pos,
        bytes calldata buf,
        uint64 len
    ) public pure returns (bool, uint64, Packet memory) {
        // Message instance
        Packet memory instance;
        // Previous field number
        uint64 previous_field_number = 0;
        // Current position in the buffer
        uint64 pos = initial_pos;

        // Sanity checks
        if (pos + len < pos) {
            return (false, pos, instance);
        }

        while (pos - initial_pos < len) {
            // Decode the key (field number and wire type)
            bool success;
            uint64 field_number;
            ProtobufLib.WireType wire_type;
            (success, pos, field_number, wire_type) = ProtobufLib.decode_key(pos, buf);
            if (!success) {
                return (false, pos, instance);
            }

            // Check that the field number is within bounds
            if (field_number > 8) {
                return (false, pos, instance);
            }

            // Check that the field number of monotonically increasing
            if (field_number <= previous_field_number) {
                return (false, pos, instance);
            }

            // Check that the wire type is correct
            success = check_key(field_number, wire_type);
            if (!success) {
                return (false, pos, instance);
            }

            // Actually decode the field
            (success, pos) = decode_field(pos, buf, len, field_number, instance);
            if (!success) {
                return (false, pos, instance);
            }

            previous_field_number = field_number;
        }

        // Decoding must have consumed len bytes
        if (pos != initial_pos + len) {
            return (false, pos, instance);
        }

        return (true, pos, instance);
    }

    function check_key(uint64 field_number, ProtobufLib.WireType wire_type) internal pure returns (bool) {
        if (field_number == 1) {
            return wire_type == ProtobufLib.WireType.Varint;
        }

        if (field_number == 2) {
            return wire_type == ProtobufLib.WireType.LengthDelimited;
        }

        if (field_number == 3) {
            return wire_type == ProtobufLib.WireType.LengthDelimited;
        }

        if (field_number == 4) {
            return wire_type == ProtobufLib.WireType.LengthDelimited;
        }

        if (field_number == 5) {
            return wire_type == ProtobufLib.WireType.LengthDelimited;
        }

        if (field_number == 6) {
            return wire_type == ProtobufLib.WireType.LengthDelimited;
        }

        if (field_number == 7) {
            return wire_type == ProtobufLib.WireType.LengthDelimited;
        }

        if (field_number == 8) {
            return wire_type == ProtobufLib.WireType.Varint;
        }

        return false;
    }

    function decode_field(
        uint64 initial_pos,
        bytes memory buf,
        uint64 len,
        uint64 field_number,
        Packet memory instance
    ) internal pure returns (bool, uint64) {
        uint64 pos = initial_pos;

        if (field_number == 1) {
            bool success;
            (success, pos) = decode_1(pos, buf, instance);
            if (!success) {
                return (false, pos);
            }

            return (true, pos);
        }

        if (field_number == 2) {
            bool success;
            (success, pos) = decode_2(pos, buf, instance);
            if (!success) {
                return (false, pos);
            }

            return (true, pos);
        }

        if (field_number == 3) {
            bool success;
            (success, pos) = decode_3(pos, buf, instance);
            if (!success) {
                return (false, pos);
            }

            return (true, pos);
        }

        if (field_number == 4) {
            bool success;
            (success, pos) = decode_4(pos, buf, instance);
            if (!success) {
                return (false, pos);
            }

            return (true, pos);
        }

        if (field_number == 5) {
            bool success;
            (success, pos) = decode_5(pos, buf, instance);
            if (!success) {
                return (false, pos);
            }

            return (true, pos);
        }

        if (field_number == 6) {
            bool success;
            (success, pos) = decode_6(pos, buf, instance);
            if (!success) {
                return (false, pos);
            }

            return (true, pos);
        }

        if (field_number == 7) {
            bool success;
            (success, pos) = decode_7(pos, buf, instance);
            if (!success) {
                return (false, pos);
            }

            return (true, pos);
        }

        if (field_number == 8) {
            bool success;
            (success, pos) = decode_8(pos, buf, instance);
            if (!success) {
                return (false, pos);
            }

            return (true, pos);
        }

        return (false, pos);
    }

    // Packet.sequence
    function decode_1(uint64 pos, bytes memory buf, Packet memory instance) internal pure returns (bool, uint64) {
        bool success;

        uint64 v;
        (success, pos, v) = ProtobufLib.decode_uint64(pos, buf);
        if (!success) {
            return (false, pos);
        }

        // Default value must be omitted
        if (v == 0) {
            return (false, pos);
        }

        instance.sequence = v;

        return (true, pos);
    }

    // Packet.source_port
    function decode_2(uint64 pos, bytes memory buf, Packet memory instance) internal pure returns (bool, uint64) {
        bool success;

        string memory v;
        (success, pos, v) = ProtobufLib.decode_string(pos, buf);
        if (!success) {
            return (false, pos);
        }

        // Default value must be omitted
        if (bytes(v).length == 0) {
            return (false, pos);
        }

        instance.source_port = v;

        return (true, pos);
    }

    // Packet.source_channel
    function decode_3(uint64 pos, bytes memory buf, Packet memory instance) internal pure returns (bool, uint64) {
        bool success;

        string memory v;
        (success, pos, v) = ProtobufLib.decode_string(pos, buf);
        if (!success) {
            return (false, pos);
        }

        // Default value must be omitted
        if (bytes(v).length == 0) {
            return (false, pos);
        }

        instance.source_channel = v;

        return (true, pos);
    }

    // Packet.destination_port
    function decode_4(uint64 pos, bytes memory buf, Packet memory instance) internal pure returns (bool, uint64) {
        bool success;

        string memory v;
        (success, pos, v) = ProtobufLib.decode_string(pos, buf);
        if (!success) {
            return (false, pos);
        }

        // Default value must be omitted
        if (bytes(v).length == 0) {
            return (false, pos);
        }

        instance.destination_port = v;

        return (true, pos);
    }

    // Packet.destination_channel
    function decode_5(uint64 pos, bytes memory buf, Packet memory instance) internal pure returns (bool, uint64) {
        bool success;

        string memory v;
        (success, pos, v) = ProtobufLib.decode_string(pos, buf);
        if (!success) {
            return (false, pos);
        }

        // Default value must be omitted
        if (bytes(v).length == 0) {
            return (false, pos);
        }

        instance.destination_channel = v;

        return (true, pos);
    }

    // Packet.data
    function decode_6(uint64 pos, bytes memory buf, Packet memory instance) internal pure returns (bool, uint64) {
        bool success;

        uint64 len;
        (success, pos, len) = ProtobufLib.decode_bytes(pos, buf);
        if (!success) {
            return (false, pos);
        }

        // Default value must be omitted
        if (len == 0) {
            return (false, pos);
        }

        instance.data = new bytes(len);
        for (uint64 i = 0; i < len; i++) {
            instance.data[i] = buf[pos + i];
        }

        pos = pos + len;

        return (true, pos);
    }

    // Packet.timeout_height
    function decode_7(uint64 pos, bytes memory buf, Packet memory instance) internal pure returns (bool, uint64) {
        bool success;

        uint64 len;
        (success, pos, len) = ProtobufLib.decode_embedded_message(pos, buf);
        if (!success) {
            return (false, pos);
        }

        // Default value must be omitted
        if (len == 0) {
            return (false, pos);
        }

        Height memory nestedInstance;
        (success, pos, nestedInstance) = HeightCodec.decode(pos, buf, len);
        if (!success) {
            return (false, pos);
        }

        instance.timeout_height = nestedInstance;

        return (true, pos);
    }

    // Packet.timeout_timestamp
    function decode_8(uint64 pos, bytes memory buf, Packet memory instance) internal pure returns (bool, uint64) {
        bool success;

        uint64 v;
        (success, pos, v) = ProtobufLib.decode_uint64(pos, buf);
        if (!success) {
            return (false, pos);
        }

        // Default value must be omitted
        if (v == 0) {
            return (false, pos);
        }

        instance.timeout_timestamp = v;

        return (true, pos);
    }
}

contract PacketTest is Test {
    PacketContract packetContract = new PacketContract();

    bytes PacketPayload =
        hex'0801120b736f757263652d706f72741a0e736f757263652d6368616e6e656c221064657374696e6174696f6e2d706f72742a1364657374696e6174696f6e2d6368616e6e656c3204646174613a04080210034004';

    // decode from memory bytes should succeed
    function testDecodeMemory() public {
        Packet memory packet = decodePacket(PacketPayload);
    }

    // decode from calldata bytes should succeed
    function testDecodeCalldata() public {
        (bool ok, uint64 _endPos, Packet memory packet) = packetContract.decode(
            0,
            PacketPayload,
            uint64(PacketPayload.length)
        );
    }

    // assert that the packet is decoded correctly and returns the parsed packet
    function decodePacket(bytes memory data) public returns (Packet memory packet) {
        (bool ok, uint64 _endPos, Packet memory _packet) = PacketCodec.decode(0, data, uint64(data.length));
        assert(ok);
        return _packet;
    }
}
