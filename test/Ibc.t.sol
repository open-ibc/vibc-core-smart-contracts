// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../contracts/libs/Ibc.sol";
import "../contracts/libs/IbcUtils.sol";
import "forge-std/Test.sol";
import {IbcChannelReceiver} from "../contracts/interfaces/IbcReceiver.sol";

contract IbcTest is Test {
    function test_packet_commitment_proof_key() public {
        IbcPacket memory packet =
            IbcPacket(IbcEndpoint("portid", hex"6368616e6e656c2d30"), IbcEndpoint("", 0), 12, hex"", Height(0, 0), 0);
        assertEq("commitments/ports/portid/channels/channel-0/sequences/12", Ibc.packetCommitmentProofKey(packet));
    }

    function test_packet_ack_proof_key() public {
        IbcPacket memory packet =
            IbcPacket(IbcEndpoint("", 0), IbcEndpoint("portid", hex"6368616e6e656c2d30"), 12, hex"", Height(0, 0), 0);
        assertEq("acks/ports/portid/channels/channel-0/sequences/12", Ibc.ackProofKey(packet));
    }

    function test_channel_proof_key() public {
        bytes memory key =
            bytes("channelEnds/ports/polyibc.eth1.71C95911E9a5D330f4D621842EC243EE1343292e/channels/channel-0");

        assertEq(
            key,
            Ibc.channelProofKey(
                "polyibc.eth1.71C95911E9a5D330f4D621842EC243EE1343292e", IbcUtils.toBytes32("channel-0")
            )
        );
    }

    function test_string_to_bytes32() public {
        assertEq(bytes32(hex"6368616e6e656c2d30"), IbcUtils.toBytes32("channel-0"));
    }

    function test_bytes32_to_string() public {
        assertEq("channel-0", Ibc.toStr(bytes32(hex"6368616e6e656c2d30")));
    }

    function test_uint256_to_string() public {
        assertEq("1", Ibc.toStr(1));
        assertEq("112233445566", Ibc.toStr(112_233_445_566));
        assertEq("16", Ibc.toStr(16));
    }

    function test_parse_ack() public {
        // this data is taken from the write_acknowledgement event emitted by polymer
        bytes memory ack =
            bytes('{"result":"eyAiYWNjb3VudCI6ICJhY2NvdW50IiwgInJlcGx5IjogImdvdCB0aGUgbWVzc2FnZSIgfQ=="}');

        AckPacket memory parsed = Ibc.parseAckData(ack);
        assertTrue(parsed.success);
        assertEq(bytes('{ "account": "account", "reply": "got the message" }'), parsed.data);

        bytes memory error = bytes('{"error":"this is an error message"}');
        AckPacket memory parsederr = Ibc.parseAckData(error);
        assertFalse(parsederr.success);
        assertEq(bytes("this is an error message"), parsederr.data);
    }

    function assert_Equal_Before_and_After_Encoding(address a1, uint256 mwBitmap, address a2, bytes memory appData)
        internal
    {
        bytes memory afterEncoding = IbcUtils.toUniversalPacketBytes(
            UniversalPacket(IbcUtils.toBytes32(a1), mwBitmap, IbcUtils.toBytes32(a2), appData)
        );

        UniversalPacket memory afterDecoding = IbcUtils.fromUniversalPacketBytes(afterEncoding);
        assertEq(a1, IbcUtils.toAddress(afterDecoding.srcPortAddr));
        assertEq(mwBitmap, afterDecoding.mwBitmap);
        assertEq(a2, IbcUtils.toAddress(afterDecoding.destPortAddr));
        assertEq(appData, afterDecoding.appData);
    }

    function test_To_From_UniversalPacketBytes() public {
        // Standard
        assert_Equal_Before_and_After_Encoding(vm.addr(1), uint256(123), vm.addr(2), "helloooo decode this for me ");
        // empty string
        assert_Equal_Before_and_After_Encoding(vm.addr(1), uint256(123), vm.addr(2), ""); // empty string

        // Really long string:
        assert_Equal_Before_and_After_Encoding(
            vm.addr(1),
            uint256(123),
            vm.addr(2),
            "daf;lkdsajflkasjdv;lkjzdljga;lkgfjda;iocjvz;lkjval;dsjkf;alkdj;zlkjv;lkjaeg;ijafd;lkjzvc,mnb.kahgd;ajkfaj;dgoij;zlckjv;lzkjv;kaldfjg;alkjgf;lkzvjcx;lkvja;lkjg;aslgjdz;adf;kasjg;lkjwaea;lkjg;io1j;4kjrda;lkfjaleot8ywp89yz;dvhlsdkj"
        );
    }

    function test_InvalidHexStr_To_Address() public {
        string memory validString = "2B5AD5c4795c026514f8317c7a215E218DcCD6cF"; // Can't start with g
        assertEq(IbcUtils.hexStrToAddress(validString), vm.addr(2));

        string memory invalidHexStr1 = "123"; // Too Short
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.invalidHexStringLength.selector));
        IbcUtils.hexStrToAddress(invalidHexStr1);

        string memory invalidHexStr2 = "2B5AD5c4795c026514f8317c7a215E218DcCD6cFa"; // Too long
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.invalidHexStringLength.selector));
        IbcUtils.hexStrToAddress(invalidHexStr2);

        string memory invalidHexStr3 = "2G5AD5c4795c026514f8317c7a215E218DcCD6cF"; // Can't have G
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.invalidCharacter.selector));
        IbcUtils.hexStrToAddress(invalidHexStr3);
    }

    function test_To_From_addr_hexStr(address addr) public {
        bytes memory hexStr = IbcUtils.toHexStr(addr);
        assertEq(addr, IbcUtils.hexStrToAddress(string(hexStr)));
    }
}
