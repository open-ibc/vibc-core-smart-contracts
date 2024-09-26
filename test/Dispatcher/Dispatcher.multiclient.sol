// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.15;

import "../../contracts/libs/Ibc.sol";
import {IbcUtils} from "../../contracts/libs/IbcUtils.sol";
import {IBCErrors} from "../../contracts/libs/IbcErrors.sol";
import {Mars} from "../../contracts/examples/Mars.sol";
import {Base} from "../utils/Dispatcher.base.t.sol";
import {DummyLightClient} from "../../contracts/utils/DummyLightClient.sol";
import {OptimisticLightClient} from "../../contracts/core/OptimisticLightClient.sol";
import {ILightClient} from "../../contracts/interfaces/ILightClient.sol";
import "../../contracts/interfaces/IProofVerifier.sol";

contract DispatcherRealProofMultiClient is Base {
    string[] connectionHops0 = ["dummy-connection-1", "dummy-connection-2"];
    string[] connectionHops1 = ["connection-2", "connection-1"];
    string[] connectionHops2 = ["connection-142", "connection-149"];

    string e2ePortPrefix = "polyibc.eth1."; // portprefix that was used to generate proofs for the op light client
    string stagingPortPrefix = "polyibc.base-staging."; // remote end portprefix that was used to generate proofs for
        // sequencer client

    Mars mars;
    Mars marsII;
    address notOwner = vm.addr(1);
    ChannelEnd ch0 =
        ChannelEnd("polyibc.eth1.71C95911E9a5D330f4D621842EC243EE1343292e", IbcUtils.toBytes32("channel-0"), "1.0");
    ChannelEnd ch1 =
        ChannelEnd("polyibc.eth2.71C95911E9a5D330f4D621842EC243EE1343292e", IbcUtils.toBytes32("channel-1"), "1.0");

    ChannelEnd stagingLocalChannel = ChannelEnd(
        "polyibc.optimism-staging.df39685972488FFAFf1f468f332873284Ca0b00b", IbcUtils.toBytes32("channel-730"), "1.0"
    );
    ChannelEnd stagingRemoteChannel = ChannelEnd(
        "polyibc.base-staging.2b2697223C42F63504265ad0B732982C475956ad", IbcUtils.toBytes32("channel-731"), "1.0"
    );

    function setUp() public override {
        opLightClient = new OptimisticLightClient(1, opProofVerifier, l1BlockProvider);
        dummyLightClient = new DummyLightClient();
        (dispatcherProxy, dispatcherImplementation) = deployDispatcherProxyAndImpl(e2ePortPrefix, feeVault);
        dispatcherProxy.setClientForConnection(connectionHops0[0], dummyLightClient);
        dispatcherProxy.setClientForConnection(connectionHops1[0], opLightClient);
        dispatcherProxy.setClientForConnection(connectionHops2[0], sequencerLightClient);
        // We have to inject code at the different mars addresses
        address targetMarsAddress = 0x71C95911E9a5D330f4D621842EC243EE1343292e;
        address targetMarsAddressII = 0x2b2697223C42F63504265ad0B732982C475956ad;
        deployCodeTo("contracts/examples/Mars.sol:Mars", abi.encode(address(dispatcherProxy)), targetMarsAddress);
        deployCodeTo("contracts/examples/Mars.sol:Mars", abi.encode(address(dispatcherProxy)), targetMarsAddressII);
        mars = Mars(payable(targetMarsAddress));
        marsII = Mars(payable(targetMarsAddressII));
    }

    function test_channelOpenTry_multiClient() public {
        Ics23Proof memory dummyProof =
            load_proof("/test/payload/channel_confirm_pending_proof.hex", address(dummyLightClient));

        // We should be able to open any connection through the dummy light client
        // Should be able to open a channel easily through the dummy light client
        vm.expectEmit(true, true, true, true);
        emit ChannelOpenTry(address(mars), "1.0", ChannelOrder.NONE, false, connectionHops0, ch1.portId, ch1.channelId);
        dispatcherProxy.channelOpenTry(ch0, ChannelOrder.NONE, false, connectionHops0, ch1, dummyProof);

        // Having a dummy light client shouldn't impact opLightClient or sequencer client's ability to reject invalid
        // proofs
        vm.expectRevert();
        dispatcherProxy.channelOpenTry(ch1, ChannelOrder.NONE, false, connectionHops1, ch0, dummyProof);

        vm.expectRevert();
        dispatcherProxy.channelOpenTry(
            stagingRemoteChannel, ChannelOrder.NONE, false, connectionHops2, stagingLocalChannel, dummyProof
        );
    }

    function test_sendPacket_multiClient() public {
        // Test that a channel can be confirmed and a packet can be recieved on the dest chain
        Ics23Proof memory openChannelProof =
            load_proof("/test/payload/channel_confirm_pending_proof.hex", address(opLightClient));
        Ics23Proof memory dummyProof =
            load_proof("/test/payload/channel_confirm_pending_proof.hex", address(dummyLightClient)); // Can use try pending
            // even for a confirm pending since dummy proof

        // Set up valid connections from both dummy and real clients
        dispatcherProxy.channelOpenConfirm(ch0, connectionHops0, ChannelOrder.NONE, false, ch1, dummyProof);
        dispatcherProxy.channelOpenConfirm(ch1, connectionHops1, ChannelOrder.NONE, false, ch0, openChannelProof);
        dispatcherProxy.setPortPrefix(stagingPortPrefix); // Have to change the portprefix, since the sequencer
            // proofs were fetched from staging
        Ics23Proof memory openChannelSequencerProof =
            load_proof("/test/payload/channel_confirm_pending_sequencer_proof.hex", address(sequencerLightClient));
        dispatcherProxy.channelOpenConfirm(
            stagingRemoteChannel,
            connectionHops2,
            ChannelOrder.NONE,
            false,
            stagingLocalChannel,
            openChannelSequencerProof
        );
        dispatcherProxy.setPortPrefix(e2ePortPrefix); //Now we set the port prefix

        // Now test that we can send a packet via the op light client
        Ics23Proof memory packetProof = load_proof("/test/payload/packet_ack_proof.hex", address(opLightClient));
        IbcEndpoint memory src = IbcEndpoint(IbcUtils.addressToPortId(e2ePortPrefix, address(mars)), ch1.channelId);
        IbcEndpoint memory dest = IbcEndpoint(ch0.portId, ch0.channelId);

        IbcPacket memory packet = IbcPacket(src, dest, uint64(1), bytes("sendPacket"), Height(0, 0), 0);
        _storeSendPacketCommitment(address(mars), packet.src.channelId, 1, 1);
        bytes memory ackData = bytes.concat(keccak256(hex"22726573756c7422"));

        // OpProofverifier will cause acknowledgement to fail
        vm.expectRevert(abi.encodeWithSelector(IAppStateVerifier.InvalidProofKey.selector));
        dispatcherProxy.acknowledgement(packet, ackData, packetProof);

        // Sequencer client will cause acknowledgement to fail
        dispatcherProxy.setPortPrefix(stagingPortPrefix);
        packet.dest = IbcEndpoint(stagingRemoteChannel.portId, stagingRemoteChannel.channelId);
        packet.src = IbcEndpoint(stagingRemoteChannel.portId, stagingRemoteChannel.channelId);
        vm.expectRevert(abi.encodeWithSelector(IAppStateVerifier.InvalidProofKey.selector));
        dispatcherProxy.acknowledgement(packet, ackData, packetProof);

        // Only the dummy client and the packet should go through since it circumvents the proof
        dispatcherProxy.setPortPrefix(e2ePortPrefix);
        packet.src = IbcEndpoint(ch0.portId, ch0.channelId);
        _storeSendPacketCommitment(address(mars), ch0.channelId, 1, 1);
        dispatcherProxy.acknowledgement(packet, ackData, dummyProof);
    }

    function test_Dispatcher_Prevents_nonOwner_SetConnection() public {
        vm.startPrank(notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        dispatcherProxy.setClientForConnection("malicious-connection-1", opLightClient);
    }

    function test_Dispatcher_Prevents_nonOwner_RemoveConnection() public {
        vm.startPrank(notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        dispatcherProxy.removeConnection(connectionHops0[0]);
    }

    function test_addr0_channels_cannot_be_added() public {
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.invalidAddress.selector));
        dispatcherProxy.setClientForConnection(connectionHops0[0], ILightClient(address(0)));
    }

    function test_Dispatcher_removeConnection() public {
        // Make sure that the connection exists before we delete it
        assertEq(_getConnectiontoClientIdMapping(connectionHops1[0]), address(opLightClient));
        Ics23Proof memory openChannelProof =
            load_proof("/test/payload/channel_confirm_pending_proof.hex", address(opLightClient));

        // Set up valid connection from channel 1
        dispatcherProxy.channelOpenConfirm(ch1, connectionHops1, ChannelOrder.NONE, false, ch0, openChannelProof);
        IbcEndpoint memory src = IbcEndpoint(IbcUtils.addressToPortId(e2ePortPrefix, address(mars)), ch1.channelId);
        IbcEndpoint memory dest = IbcEndpoint(ch0.portId, ch0.channelId);
        IbcPacket memory packet = IbcPacket(src, dest, uint64(1), bytes("sendPacket"), Height(0, 0), 0);

        // Remove connection to ensure packet can't be acked after removing light client
        dispatcherProxy.removeConnection(connectionHops1[0]);
        assertEq(_getConnectiontoClientIdMapping(connectionHops1[0]), address(0));

        vm.expectRevert(abi.encodeWithSelector(IBCErrors.lightClientNotFound.selector, connectionHops1[0]));
        dispatcherProxy.acknowledgement(packet, bytes("ack"), openChannelProof);
    }
}
