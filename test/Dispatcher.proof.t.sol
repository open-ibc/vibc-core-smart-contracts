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

abstract contract DispatcherIbcWithRealProofsSuite is IbcEventsEmitter, Base {
    using stdStorage for StdStorage;

    string portPrefix1;
    string portPrefix2;
    string portId1;
    string portId2;

    Mars mars;
    OptimisticLightClient consensusStateManager;

    CounterParty ch0;
    CounterParty ch1;
    string[] connectionHops0 = ["connection-0", "connection-3"];
    string[] connectionHops1 = ["connection-2", "connection-1"];

    function test_ibc_channel_open_init() public {
        vm.expectEmit(true, true, true, true);
        emit ChannelOpenInit(address(mars), "1.0", ChannelOrder.NONE, false, connectionHops1, ch1.portId);

        // since this is open chann init, the proof is not used. so use an invalid one
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

        // plant a fake packet commitment so the ack checks go through
        // Stdstore doesn't work for proxies so we have to use store
        // use "forge inspect --storage" to find the nested mapping slot
        bytes32 slot1 = keccak256(abi.encode(address(mars), uint32(156))); // current nested mapping slot: 157
        bytes32 slot2 = keccak256(abi.encode(ch0.channelId, slot1));
        bytes32 slot3 = keccak256(abi.encode(uint256(1), slot2));
        vm.store(address(dispatcherProxy), slot3, bytes32(uint256(1)));

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

    function test_timeout_packet() public {
        vm.skip(true); // not implemented
    }

    function load_proof(string memory filepath) internal returns (Ics23Proof memory) {
        (bytes32 apphash, Ics23Proof memory proof) =
            abi.decode(vm.parseBytes(vm.readFile(string.concat(rootDir, filepath))), (bytes32, Ics23Proof));

        // this loads the app hash we got from the testing data into the consensus state manager internals
        // at the height it's supposed to go. That is, a block less than where the proof was generated from.
        stdstore.target(address(consensusStateManager)).sig("consensusStates(uint256)").with_key(proof.height - 1)
            .checked_write(apphash);
        // trick the fraud time window check
        vm.warp(block.timestamp + 1);

        return proof;
    }
}

contract DispatcherIbcWithRealProofs is DispatcherIbcWithRealProofsSuite {
    function setUp() public override {
        super.setUp();

        portPrefix1 = "polyibc.eth1.";
        portPrefix2 = "polyibc.eth2.";
        consensusStateManager = new OptimisticLightClient(1, opProofVerifier, l1BlockProvider);
        (dispatcherProxy, dispatcherImplementation) = deployDispatcherProxyAndImpl(portPrefix1, consensusStateManager);

        address targetMarsAddress = 0x71C95911E9a5D330f4D621842EC243EE1343292e;
        Mars marsTemplate = new Mars(dispatcherProxy);
        deployCodeTo("Mars.sol", abi.encode(address(dispatcherProxy)), targetMarsAddress);
        // vm.etch(targetMarsAddress, address(marsTemplate).code);
        // vm.store(targetMarsAddress, bytes32(uint256(1)), bytes32(uint256(uint160(address(dispatcherProxy)))));
        // vm.store(targetMarsAddress, bytes32(uint256(0)), bytes32(uint256(uint160(address(dispatcherProxy)))));

        Mars mars = Mars(payable(targetMarsAddress));
        // mars = new Mars(dispatcherProxy);
        // mars = Mars()

        // CounterParty ch0 = CounterParty(
        //     "polyibc.eth1.71C95911E9a5D330f4D621842EC243EE1343292e", IbcUtils.toBytes32("channel-0"), "1.0"
        // );
        // CounterParty ch1 = CounterParty(
        //     "polyibc.eth2.71C95911E9a5D330f4D621842EC243EE1343292e", IbcUtils.toBytes32("channel-1"), "1.0"
        // );

        // portId1 = IbcUtils.addressToPortId(portPrefix1, address(mars));
        // portId2 = IbcUtils.addressToPortId(portPrefix2, address(mars));
        portId1 = "polyibc.eth1.71C95911E9a5D330f4D621842EC243EE1343292e";
        portId2 = "polyibc.eth2.71C95911E9a5D330f4D621842EC243EE1343292e";
        console2.log("portId1 ", portId1);
        console2.log("portId2 ", portId2);
        console2.log("owner", address(mars.dispatcher()), "dispatcher", address(dispatcherProxy));
        console2.logBytes32(vm.load(address(mars), bytes32(uint256(0))));
        console2.logBytes32(vm.load(address(marsTemplate), bytes32(uint256(0))));
        console2.logBytes32(vm.load(address(mars), bytes32(uint256(1))));
        console2.logBytes32(vm.load(address(marsTemplate), bytes32(uint256(1))));
        ch0 = CounterParty(portId1, IbcUtils.toBytes32("channel-0"), "1.0");
        ch1 = CounterParty(portId2, IbcUtils.toBytes32("channel-1"), "1.0");
    }
}
