// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.15;

import "../contracts/libs/Ibc.sol";
import {Base} from "./Dispatcher.base.t.sol";
import {Dispatcher} from "../contracts/core/Dispatcher.sol";
import {IDispatcher} from "../contracts/interfaces/IDispatcher.sol";
import {Mars} from "../contracts/examples/Mars.sol";
import {IbcDispatcher, IbcEventsEmitter} from "../contracts/interfaces/IbcDispatcher.sol";
import "../contracts/core/OpLightClient.sol";
import "./Proof.base.t.sol";
import {stdStorage, StdStorage} from "forge-std/Test.sol";
import {ChannelHandshakeSetting} from "./Dispatcher.base.t.sol";
import {DummyLightClient} from "../contracts/utils/DummyLightClient.sol";
import {DispatcherSendPacketTestSuite, ChannelOpenTestBaseSetup} from "./Dispatcher.t.sol";

contract DispatcherProofTestUtils is Base {
    using stdStorage for StdStorage;

    function load_proof(string memory filepath) internal returns (Ics23Proof memory) {
        (bytes32 apphash, Ics23Proof memory proof) =
            abi.decode(vm.parseBytes(vm.readFile(string.concat(rootDir, filepath))), (bytes32, Ics23Proof));

        // this loads the app hash we got from the testing data into the consensus state manager internals
        // at the height it's supposed to go. That is, a block less than where the proof was generated from.
        stdstore.target(address(opLightClient)).sig("consensusStates(uint256)").with_key(proof.height - 1).checked_write(
            apphash
        );
        // trick the fraud time window check
        vm.warp(block.timestamp + 1);

        return proof;
    }
}

abstract contract DispatcherIbcWithRealProofsSuite is IbcEventsEmitter, DispatcherProofTestUtils {
    Mars mars;

    ChannelEnd ch0;
    ChannelEnd ch1;
    string[] connectionHops0 = ["connection-0", "connection-3"];
    string[] connectionHops1 = ["connection-2", "connection-1"];

    function test_ibc_channel_open_init() public {
        vm.expectEmit(true, true, true, true);
        emit ChannelOpenInit(address(mars), "1.0", ChannelOrder.NONE, false, connectionHops1, ch1.portId);

        // since this is open chann init, the proof is not used. so use an invalid one
        vm.prank(address(mars));
        dispatcherProxy.channelOpenInit(ch1.version, ChannelOrder.NONE, false, connectionHops1, ch1.portId);
    }

    function test_ibc_channel_open_try() public {
        Ics23Proof memory proof = load_proof("/test/payload/channel_try_pending_proof.hex");

        vm.expectEmit(true, true, true, true);
        emit ChannelOpenTry(address(mars), "1.0", ChannelOrder.NONE, false, connectionHops1, ch0.portId, ch0.channelId);

        dispatcherProxy.channelOpenTry(ch1, ChannelOrder.NONE, false, connectionHops1, ch0, proof);
    }

    function test_ibc_channel_ack_123_a() public {
        Ics23Proof memory proof = load_proof("/test/payload/channel_ack_pending_proof.hex");

        vm.expectEmit(true, true, true, true);
        emit ChannelOpenAck(address(mars), ch0.channelId);

        dispatcherProxy.channelOpenAck(ch0, connectionHops0, ChannelOrder.NONE, false, ch1, proof);
    }

    function test_ibc_channel_confirm() public {
        Ics23Proof memory proof = load_proof("/test/payload/channel_confirm_pending_proof.hex");

        vm.expectEmit(true, true, true, true);
        emit ChannelOpenConfirm(address(mars), ch1.channelId);

        dispatcherProxy.channelOpenConfirm(ch1, connectionHops1, ChannelOrder.NONE, false, ch0, proof);
    }

    function test_ack_packet() public {
        Ics23Proof memory proof = load_proof("/test/payload/packet_ack_proof.hex");

        _storeSendPacketCommitment(address(mars), ch0.channelId, 1, 1);

        // store connection in channelid to connection
        bytes32 connectionStr = bytes32(0x636f6e6e656374696f6e2d300000000000000000000000000000000000000018); // Connection-0
        _storeChannelidToConnectionMapping(ch0.channelId, connectionStr);

        IbcPacket memory packet;
        packet.data = bytes("packet-1");
        packet.timeoutTimestamp = 15_566_401_733_896_437_760;
        packet.src.channelId = ch0.channelId;
        packet.src.portId = string(abi.encodePacked("polyibc.eth1.", IbcUtils.toHexStr(address(mars))));
        packet.dest.portId = ch1.portId;
        packet.dest.channelId = ch1.channelId;
        packet.sequence = 1;

        // this data is taken from the write_acknowledgement event emitted by polymer
        bytes memory ack =
            bytes('{"result":"eyAiYWNjb3VudCI6ICJhY2NvdW50IiwgInJlcGx5IjogImdvdCB0aGUgbWVzc2FnZSIgfQ=="}');

        vm.expectEmit(true, true, true, true);
        emit Acknowledgement(address(mars), packet.src.channelId, packet.sequence);

        dispatcherProxy.acknowledgement(packet, ack, proof);
    }

    function test_recv_packet() public {
        Ics23Proof memory proof = load_proof("/test/payload/packet_commitment_proof.hex");
        // store connection in channelid to connection
        bytes32 connectionStr = bytes32(0x636f6e6e656374696f6e2d300000000000000000000000000000000000000018); // Connection-0
        _storeChannelidToConnectionMapping(ch1.channelId, connectionStr);

        // this data is taken from polymerase/tests/e2e/tests/evm.events.test.ts MarsDappPair.createSentPacket()
        IbcPacket memory packet;
        packet.data = bytes("packet-1");
        packet.timeoutTimestamp = 15_566_401_733_896_437_760;
        packet.dest.channelId = ch1.channelId;
        packet.dest.portId = string(abi.encodePacked("polyibc.eth1.", IbcUtils.toHexStr(address(mars))));
        packet.src.portId = ch0.portId;
        packet.src.channelId = ch0.channelId;
        packet.sequence = 1;

        vm.expectEmit(true, true, true, true);
        emit WriteAckPacket(
            address(mars),
            packet.dest.channelId,
            packet.sequence,
            AckPacket(true, abi.encodePacked('{ "account": "account", "reply": "got the message" }'))
        );
        dispatcherProxy.recvPacket(packet, proof);
    }

    function test_timeout_packet_revert() public {
        bytes32 connectionStr = bytes32(0x636f6e6e656374696f6e2d310000000000000000000000000000000000000018); // Connection-1
        _storeChannelidToConnectionMapping(ch1.channelId, connectionStr);
        // Timeout reverts since it is not yet implemented
        Ics23Proof memory proof = load_proof("/test/payload/packet_commitment_proof.hex");
        IbcPacket memory packet;
        packet.data = bytes("packet-1");
        packet.timeoutTimestamp = 15_566_401_733_896_437_760;
        packet.dest.channelId = ch1.channelId;
        packet.dest.portId = string(abi.encodePacked("polyibc.eth1.", IbcUtils.toHexStr(address(mars))));
        packet.src.portId = string(abi.encodePacked("polyibc.eth1.", IbcUtils.toHexStr(address(mars))));
        packet.src.channelId = ch1.channelId;
        packet.sequence = 1;

        vm.expectRevert(abi.encodeWithSelector(ProofVerifier.MethodNotImplemented.selector));
        dispatcherProxy.timeout(packet, proof);
    }
}

contract DispatcherIbcWithRealProofs is DispatcherIbcWithRealProofsSuite {
    function setUp() public override {
        super.setUp();

        string memory portPrefix1 = "polyibc.eth1.";
        string memory portPrefix2 = "polyibc.eth2.";
        opLightClient = new OptimisticLightClient(1, opProofVerifier, l1BlockProvider);
        (dispatcherProxy, dispatcherImplementation) = deployDispatcherProxyAndImpl(portPrefix1);
        dispatcherProxy.setNewConnection(connectionHops0[0], opLightClient);
        dispatcherProxy.setNewConnection(connectionHops0[1], opLightClient);
        dispatcherProxy.setNewConnection(connectionHops1[0], opLightClient);
        dispatcherProxy.setNewConnection(connectionHops1[1], opLightClient);
        address targetMarsAddress = 0x71C95911E9a5D330f4D621842EC243EE1343292e;
        deployCodeTo("contracts/examples/Mars.sol:Mars", abi.encode(address(dispatcherProxy)), targetMarsAddress);
        mars = Mars(payable(targetMarsAddress));
        string memory portId1 = "polyibc.eth1.71C95911E9a5D330f4D621842EC243EE1343292e";
        string memory portId2 = "polyibc.eth2.71C95911E9a5D330f4D621842EC243EE1343292e";
        ch0 = ChannelEnd(portId1, IbcUtils.toBytes32("channel-0"), "1.0");
        ch1 = ChannelEnd(portId2, IbcUtils.toBytes32("channel-1"), "1.0");
    }
}
