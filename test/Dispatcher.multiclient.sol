// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.15;

import "../contracts/libs/Ibc.sol";
import {Mars} from "../contracts/examples/Mars.sol";
import {DispatcherProofTestUtils} from "./Dispatcher.proof.t.sol";
import {DummyLightClient} from "../contracts/utils/DummyLightClient.sol";
import {OptimisticLightClient} from "../contracts/core/OpLightClient.sol";
import {LightClient} from "../contracts/interfaces/LightClient.sol";
import "../contracts/interfaces/ProofVerifier.sol";

contract DispatcherRealProofMultiClient is DispatcherProofTestUtils {
    string[] connectionHops0 = ["dummy-connection-1", "dummy-connection-2"];
    string[] connectionHops1 = ["connection-2", "connection-1"];
    // string[] connectionHops1 = ["op-connection-1", "op-connection-2"];
    Mars mars;
    address notOwner = vm.addr(1);
    ChannelEnd ch0 =
        ChannelEnd("polyibc.eth1.71C95911E9a5D330f4D621842EC243EE1343292e", IbcUtils.toBytes32("channel-0"), "1.0");
    ChannelEnd ch1 =
        ChannelEnd("polyibc.eth2.71C95911E9a5D330f4D621842EC243EE1343292e", IbcUtils.toBytes32("channel-1"), "1.0");

    function setUp() public override {
        opLightClient = new OptimisticLightClient(1, opProofVerifier, l1BlockProvider);
        dummyLightClient = new DummyLightClient();
        (dispatcherProxy, dispatcherImplementation) = deployDispatcherProxyAndImpl("polyibc.eth1.");
        dispatcherProxy.setClientForConnection(connectionHops0[0], dummyLightClient);
        dispatcherProxy.setClientForConnection(connectionHops1[0], opLightClient);
        address targetMarsAddress = 0x71C95911E9a5D330f4D621842EC243EE1343292e;
        deployCodeTo("Mars.sol", abi.encode(address(dispatcherProxy)), targetMarsAddress);
        mars = Mars(payable(targetMarsAddress));
    }

    function test_channelOpenTry_multiClient() public {
        Ics23Proof memory dummyProof = load_proof("/test/payload/channel_confirm_pending_proof.hex");

        // We should be able to open any connection through the dummy light client
        // Should be able to open a channel easily through the dummy light client
        vm.expectEmit(true, true, true, true);
        emit ChannelOpenTry(address(mars), "1.0", ChannelOrder.NONE, false, connectionHops0, ch1.portId, ch1.channelId);
        dispatcherProxy.channelOpenTry(ch0, ChannelOrder.NONE, false, connectionHops0, ch1, dummyProof);

        // Having a dummy light client shouldn't impact opLightClient's ability to reject invalid proofs
        vm.expectRevert(abi.encodeWithSelector(ProofVerifier.InvalidProofValue.selector));
        dispatcherProxy.channelOpenTry(ch1, ChannelOrder.NONE, false, connectionHops1, ch0, dummyProof);
    }

    function test_sendPacket_multiClient() public {
        Ics23Proof memory openChannelProof = load_proof("/test/payload/channel_confirm_pending_proof.hex");
        Ics23Proof memory invalidProof = load_proof("/test/payload/packet_ack_proof.hex");

        // Set up valid connections from both dummy and real clients
        dispatcherProxy.channelOpenConfirm(ch0, connectionHops0, ChannelOrder.NONE, false, ch1, invalidProof);
        dispatcherProxy.channelOpenConfirm(ch1, connectionHops1, ChannelOrder.NONE, false, ch0, openChannelProof);

        // dispatcherproxy.sendpacket should revert using channel 1
        IbcEndpoint memory src = IbcEndpoint(IbcUtils.addressToPortId("polyibc.eth1.", address(mars)), ch1.channelId);
        IbcEndpoint memory dest = IbcEndpoint(ch0.portId, ch0.channelId);

        IbcPacket memory packet = IbcPacket(src, dest, uint64(1), bytes("sendPacket"), Height(0, 0), 0);
        _storeSendPacketCommitment(address(mars), packet.src.channelId, 1, 1);
        bytes memory ackData = bytes.concat(keccak256(hex"22726573756c7422"));

        // OpProofverifier will cause acknowledgement to fail
        vm.expectRevert(abi.encodeWithSelector(ProofVerifier.InvalidProofKey.selector));
        dispatcherProxy.acknowledgement(packet, ackData, invalidProof);

        // Now we change the client to the dummy client and the packet should go through since it circumvents the proof
        dispatcherProxy.setClientForConnection(connectionHops1[0], dummyLightClient);
        dispatcherProxy.acknowledgement(packet, ackData, invalidProof);
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
        dispatcherProxy.setClientForConnection(connectionHops0[0], LightClient(address(0)));
    }

    function test_Dispatcher_removeConnection() public {
        Ics23Proof memory openChannelProof = load_proof("/test/payload/channel_confirm_pending_proof.hex");

        // Set up valid connection from channel 1
        dispatcherProxy.channelOpenConfirm(ch1, connectionHops1, ChannelOrder.NONE, false, ch0, openChannelProof);
        IbcEndpoint memory src = IbcEndpoint(IbcUtils.addressToPortId("polyibc.eth1.", address(mars)), ch1.channelId);
        IbcEndpoint memory dest = IbcEndpoint(ch0.portId, ch0.channelId);
        IbcPacket memory packet = IbcPacket(src, dest, uint64(1), bytes("sendPacket"), Height(0, 0), 0);

        // Remove connection to ensure packet can't be acked after removing light client
        dispatcherProxy.removeConnection(connectionHops1[0]);
        assertEq(_getConnectiontoClientIdMapping(connectionHops1[0]), 0);

        vm.expectRevert(abi.encodeWithSelector(IBCErrors.lightClientNotFound.selector, connectionHops1[0]));
        dispatcherProxy.acknowledgement(packet, bytes("ack"), openChannelProof);
    }
}
