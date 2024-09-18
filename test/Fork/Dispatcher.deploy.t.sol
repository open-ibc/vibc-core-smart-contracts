// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {DispatcherUpdateClientTestSuite} from "../Dispatcher/Dispatcher.client.t.sol";
import {DispatcherIbcWithRealProofsSuite} from "../Dispatcher/Dispatcher.proof.t.sol";
import {Mars} from "../../contracts/examples/Mars.sol";
import "../../contracts/core/OptimisticLightClient.sol";
import {ChannelHandshakeTestSuite, ChannelHandshakeTest, ChannelHandshakeUtils} from "../Dispatcher/Dispatcher.t.sol";
import {LocalEnd} from "../utils/Dispatcher.base.t.sol";
import {Base, ChannelHandshakeSetting} from "../utils/Dispatcher.base.t.sol";
import {
    ChannelEnd,
    ChannelOrder,
    ChannelState,
    IbcEndpoint,
    IbcPacket,
    AckPacket,
    Ibc,
    Height
} from "../../contracts/libs/Ibc.sol";
import {IbcUtils} from "../../contracts/libs/IbcUtils.sol";
import {IbcReceiver, IbcChannelReceiver} from "../../contracts/interfaces/IbcReceiver.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts/proxy/utils/UUPSUpgradeable.sol";
import {OptimisticLightClient} from "../../contracts/core/OptimisticLightClient.sol";
import {IProofVerifier} from "../../contracts/interfaces/IProofVerifier.sol";
import {Dispatcher} from "../../contracts/core/Dispatcher.sol";

import {IDispatcher} from "../../contracts/interfaces/IDispatcher.sol";
import {UniversalChannelHandler} from "../../contracts/core/UniversalChannelHandler.sol";
import {IUniversalChannelHandler} from "../../contracts/interfaces/IUniversalChannelHandler.sol";
import {UniversalChannelHandlerV2} from "../upgradeableProxy/upgrades/UCHV2.sol";
import {DispatcherV2Initializable} from "../upgradeableProxy/upgrades/DispatcherV2Initializable.sol";
import {DispatcherV2} from "../upgradeableProxy/upgrades/DispatcherV2.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";
import {ChannelHandShakeUpgradeUtil, UpgradeTestUtils} from "../upgradeableProxy/UpgradeUtils.t.sol";
import "forge-std/Test.sol";

struct ChainAddresses {
    IDispatcher dispatcherProxy;
    IUniversalChannelHandler uch;
    ILightClient optimisticLightClient;
    address owner; // Owner Address of dispatcher
}

contract DispatcherDeployTest is ChannelHandShakeUpgradeUtil, UpgradeTestUtils {
    address targetMarsAddress = 0x71C95911E9a5D330f4D621842EC243EE1343292e; // Address that proof was generated for
    string portPrefix1 = "polyibc.eth1.";
    string portId1 = "polyibc.eth1.71C95911E9a5D330f4D621842EC243EE1343292e";
    string portId2 = "polyibc.eth2.71C95911E9a5D330f4D621842EC243EE1343292e";

    function setUp() public override {
        ChainAddresses memory addresses = ChainAddresses(
            IDispatcher(vm.envAddress("DispatcherProxy")),
            IUniversalChannelHandler(vm.envAddress("UCProxy")),
            ILightClient(vm.envAddress("OptimisticLightClient")),
            vm.envAddress("OwnerAddress")
        );

        opLightClient = addresses.optimisticLightClient; // Need to set this so that when we call load_proof, it loads
            // the proof to the right address

        mars = Mars(payable(targetMarsAddress));

        dispatcherProxy = addresses.dispatcherProxy;
        vm.prank(addresses.owner);

        // For now, we need to change the portPrefix to that of the one which was used to generate the proof. We also
        // have to set that for the connectionHop to light client mapping.
        // Ideally, we can move to eventually automatically generating the proof by querying on chain.
        dispatcherProxy.setPortPrefix(portPrefix1);
        // Use legacy mars implementation to test upgrade compatibility
        deployCodeTo("contracts/examples/Mars.sol:Mars", abi.encode(address(dispatcherProxy)), targetMarsAddress);

        // We have to set the connection hops to hard-coded values since these will be checked in the proof
        connectionHops0 = ["connection-0", "connection-3"];
        connectionHops1 = ["connection-2", "connection-1"];

        vm.startPrank(addresses.owner); // Only sender should have permission
        dispatcherProxy.setClientForConnection("connection-0", opLightClient);

        dispatcherProxy.setClientForConnection("connection-2", opLightClient);
        vm.stopPrank();

        _local = LocalEnd(IbcChannelReceiver(targetMarsAddress), portId1, "channel-0", connectionHops0, "1.0", "1.0");
        _remote = ChannelEnd(portId2, "channel-1", "1.0");

        // Do handshake as if Mars was the sending localparty
        doSrcProofChannelHandshake(_local, _remote);

        // Do handshake as if Mars was the receiving counterparty
        doDestProofChannelHandshake(
            _remote,
            ChannelEnd(_local.portId, _local.channelId, _local.versionCall),
            connectionHops1,
            _local.versionExpected,
            mars
        );

        // State should be initialized correctly & not impacted after any upgrades; used to check for malformatted state
        uint64 nextSequenceRecvValue =
            uint64(uint256(vm.load(address(dispatcherProxy), findNextSequenceRecv(address(mars), _local.channelId))));
        uint64 nextSequenceAckValue =
            uint64(uint256(vm.load(address(dispatcherProxy), findNextSequenceAck(address(mars), _local.channelId))));
        assertEq(1, nextSequenceRecvValue);
        assertEq(1, nextSequenceAckValue);
    }

    function test_Deployment_SentPacketState_Conserved() public {
        // Send a few packets as if mars was sending party
        payload = "packet-1"; // Overwrite these values for inheriting class
        packet = abi.encodePacked(payload);
        sendOnePacket(_local.channelId, 1, Mars(payable(targetMarsAddress)));
        sendOnePacket(_local.channelId, 2, Mars(payable(targetMarsAddress)));
        sendOnePacket(_local.channelId, 3, Mars(payable(targetMarsAddress)));

        uint64 nextSequenceSendValue = uint64(
            uint256(vm.load(address(dispatcherProxy), findNextSequenceSendSlot(address(mars), _local.channelId)))
        );

        assertEq(4, nextSequenceSendValue);

        // packets
        assert(vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(mars), _local.channelId, 1)) > 0);
        assert(vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(mars), _local.channelId, 2)) > 0);
        assert(vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(mars), _local.channelId, 3)) > 0);

        // Recv a packet as if it had been acked on dest chain. Note: the src chain is oblivious to the actual state of
        // dest chain, aside from what is proved from the light client. So we can get away with only mocking the
        // recvpacket proofs for testing dest chain
        Ics23Proof memory proof = load_proof("/test/payload/packet_ack_proof.hex");
        IbcPacket memory packet;
        packet.data = bytes("packet-1");
        packet.timeoutTimestamp = 15_566_401_733_896_437_760;
        packet.src.channelId = _local.channelId;
        packet.src.portId = string(abi.encodePacked("polyibc.eth1.", IbcUtils.toHexStr(address(mars))));
        packet.dest.portId = portId2;
        packet.dest.channelId = _remote.channelId;
        packet.sequence = 1;

        // This data is taken from the write_acknowledgement event emitted by polymer; this needs to be hardcoded since
        // the proof is hardcoded
        bytes memory ack =
            bytes('{"result":"eyAiYWNjb3VudCI6ICJhY2NvdW50IiwgInJlcGx5IjogImdvdCB0aGUgbWVzc2FnZSIgfQ=="}');

        vm.expectEmit(true, true, true, true);
        emit Acknowledgement(address(mars), packet.src.channelId, 1);
        dispatcherProxy.acknowledgement(packet, ack, proof);

        // We can send more packets (but we neglect testing them for now since it would require generating more proofs)
        sendOnePacket(_local.channelId, 4, Mars(payable(targetMarsAddress)));
        assert(vm.load(address(dispatcherProxy), findSendPacketCommitmentSlot(address(mars), _local.channelId, 4)) > 0);
        uint64 nextSequenceSendAfterSending = uint64(
            uint256(vm.load(address(dispatcherProxy), findNextSequenceSendSlot(address(mars), _local.channelId)))
        );
        assertEq(5, nextSequenceSendAfterSending);
    }
}
