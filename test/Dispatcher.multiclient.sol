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
    CounterParty ch0 =
        CounterParty("polyibc.eth1.71C95911E9a5D330f4D621842EC243EE1343292e", IbcUtils.toBytes32("channel-0"), "1.0");
    CounterParty ch1 =
        CounterParty("polyibc.eth2.71C95911E9a5D330f4D621842EC243EE1343292e", IbcUtils.toBytes32("channel-1"), "1.0");

    function setUp() public override {
        opLightClient = new OptimisticLightClient(1, opProofVerifier, l1BlockProvider);
        dummyLightClient = new DummyLightClient();
        (dispatcherProxy, dispatcherImplementation) = deployDispatcherProxyAndImpl("polyibc.eth1.");
        dispatcherProxy.setNewConnection(connectionHops0[0], dummyLightClient);
        dispatcherProxy.setNewConnection(connectionHops1[0], opLightClient);
        mars = new Mars(dispatcherProxy);
    }

    function test_channelOpenTry_multiClient() public {
        Ics23Proof memory dummyProof = load_proof("/test/payload/channel_confirm_pending_proof.hex");

        // We should be able to open any connection through the dummy light client
        // Should be able to open a channel easily through the dummy light client
        vm.expectEmit(true, true, true, true);
        emit ChannelOpenTry(address(mars), "1.0", ChannelOrder.NONE, false, connectionHops0, ch1.portId, ch1.channelId);
        dispatcherProxy.channelOpenTry(mars, ch0, ChannelOrder.NONE, false, connectionHops0, ch1, dummyProof);

        // Having a dummy light client shouldn't impact opLightClient's ability to reject invalid proofs
        vm.expectRevert(abi.encodeWithSelector(ProofVerifier.InvalidProofValue.selector));
        dispatcherProxy.channelOpenTry(mars, ch1, ChannelOrder.NONE, false, connectionHops1, ch0, dummyProof);
    }

    function test_sendPacket_multiClient() public {
        Ics23Proof memory openChannelProof = load_proof("/test/payload/channel_confirm_pending_proof.hex");
        Ics23Proof memory invalidProof = load_proof("/test/payload/packet_ack_proof.hex");

        // Set up valid connections from both dummy and real clients
        dispatcherProxy.channelOpenConfirm(mars, ch0, connectionHops0, ChannelOrder.NONE, false, ch1, invalidProof);
        dispatcherProxy.channelOpenConfirm(mars, ch1, connectionHops1, ChannelOrder.NONE, false, ch0, openChannelProof);

        // dispatcherproxy.sendpacket should revert using channel 1
        IbcEndpoint memory src = IbcEndpoint(IbcUtils.addressToPortId("polyibc.eth1.", address(mars)), ch1.channelId);
        IbcEndpoint memory dest = IbcEndpoint(ch0.portId, ch0.channelId);

        IbcPacket memory packet = IbcPacket(src, dest, uint64(1), bytes("sendPacket"), Height(0, 0), 0);

        // OpProofverifier will cause acknowledgement to fail
        vm.expectRevert(abi.encodeWithSelector(ProofVerifier.InvalidProofKey.selector));
        dispatcherProxy.acknowledgement(mars, packet, bytes("ack"), invalidProof);
    }

    function test_Dispatcher_Prevenrts_nonOwner_AddConnection() public {
        vm.startPrank(notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        dispatcherProxy.setNewConnection("malicious-connection-1", opLightClient);
    }

    function test_addr0_channels_cannot_be_added() public {
        vm.expectRevert(abi.encodeWithSelector(IBCErrors.invalidAddress.selector));
        dispatcherProxy.setNewConnection(connectionHops0[0], LightClient(address(0)));
    }
}
