// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import {ProofBase} from "./Proof.base.t.sol";
import "../../contracts/libs/Ibc.sol";
import {Dispatcher} from "../../contracts/core/Dispatcher.sol";
import {IDispatcher} from "../../contracts/interfaces/IDispatcher.sol";
import {IbcEventsEmitter} from "../../contracts/interfaces/IbcDispatcher.sol";
import {IbcChannelReceiver} from "../../contracts/interfaces/IbcReceiver.sol";
import "../../contracts/examples/Mars.sol";
import "../../contracts/core/OptimisticLightClient.sol";
import "../../contracts/utils/DummyLightClient.sol";
import "../../contracts/core/OptimisticProofVerifier.sol";
import {TestUtilsTest} from "./TestUtils.t.sol";

struct LocalEnd {
    IbcChannelReceiver receiver;
    string portId;
    // channelId is only used in connectIbcChannel
    bytes32 channelId;
    string[] connectionHops;
    string versionCall;
    string versionExpected;
}

struct ChannelHandshakeSetting {
    ChannelOrder ordering;
    bool feeEnabled;
    bool localInitiate;
    Ics23Proof proof;
}

// Base contract for testing Dispatcher
contract Base is IbcEventsEmitter, ProofBase, TestUtilsTest {
    uint32 CONNECTION_TO_CLIENT_ID_STARTING_SLOT = 161;
    uint32 SEND_PACKET_COMMITMENT_STARTING_SLOT = 156;
    uint64 UINT64_MAX = 18_446_744_073_709_551_615;

    Height ZERO_HEIGHT = Height(0, 0);
    uint64 maxTimeout = UINT64_MAX;

    ILightClient opLightClient = new OptimisticLightClient(1800, opProofVerifier, l1BlockProvider);
    ILightClient dummyLightClient = new DummyLightClient();

    IDispatcher public dispatcherProxy;
    Dispatcher public dispatcherImplementation;
    string portPrefix = "polyibc.eth.";
    string[] connectionHops = ["connection-1", "connection-2"];

    // ⬇️ Utility functions for testing

    // getHexStr converts an address to a hex string without the leading 0x
    function getHexBytes(address addr) internal pure returns (bytes memory) {
        bytes memory addrWithPrefix = abi.encodePacked(vm.toString(addr));
        bytes memory addrWithoutPrefix = new bytes(addrWithPrefix.length - 2);
        for (uint256 i = 0; i < addrWithoutPrefix.length; i++) {
            addrWithoutPrefix[i] = addrWithPrefix[i + 2];
        }
        return addrWithoutPrefix;
    }

    // deriveAddress derives an address from a given string deterministically for testing
    function deriveAddress(string memory str) internal pure returns (address) {
        return vm.addr(uint256(keccak256(abi.encodePacked(str))));
    }

    // ⬇️ IBC functions for testing

    /**
     * @dev Step-1 of the 4-step handshake to open an IBC channel.
     * @param le Local end settings for the channel.
     * @param re Remote end settings for the channel.
     * @param s Channel handshake settings.
     * @param expPass Expected pass status of the operation.
     * If expPass is false, `vm.expectRevert` should be called before this function.
     */
    function channelOpenInit(LocalEnd memory le, CounterParty memory re, ChannelHandshakeSetting memory s, bool expPass)
        public
    {
        if (expPass) {
            vm.expectEmit(true, true, true, true);
            emit ChannelOpenInit(
                address(le.receiver), le.versionExpected, s.ordering, s.feeEnabled, le.connectionHops, re.portId
            );
        }
        dispatcherProxy.channelOpenInit(
            le.receiver, le.versionCall, s.ordering, s.feeEnabled, le.connectionHops, re.portId
        );
    }

    /**
     * @dev Step-2 of the 4-step handshake to open an IBC channel.
     * @param le Local end settings for the channel.
     * @param re Remote end settings for the channel.
     * @param s Channel handshake settings.
     * @param expPass Expected pass status of the operation.
     * If expPass is false, `vm.expectRevert` should be called before this function.
     */
    function channelOpenTry(LocalEnd memory le, CounterParty memory re, ChannelHandshakeSetting memory s, bool expPass)
        public
    {
        if (expPass) {
            vm.expectEmit(true, true, true, true);
            emit ChannelOpenTry(
                address(le.receiver),
                le.versionExpected,
                s.ordering,
                s.feeEnabled,
                le.connectionHops,
                re.portId,
                re.channelId
            );
        }
        CounterParty memory cp = CounterParty(re.portId, re.channelId, re.version);
        dispatcherProxy.channelOpenTry(
            le.receiver,
            CounterParty(le.portId, le.channelId, le.versionCall),
            s.ordering,
            s.feeEnabled,
            le.connectionHops,
            cp,
            s.proof
        );
    }

    /**
     * @dev Step-3 of the 4-step handshake to open an IBC channel.
     * @param le Local end settings for the channel.
     * @param re Remote end settings for the channel.
     * @param s Channel handshake settings.
     * @param expPass Expected pass status of the operation.
     * If expPass is false, `vm.expectRevert` should be called before this function.
     */
    function channelOpenAck(LocalEnd memory le, CounterParty memory re, ChannelHandshakeSetting memory s, bool expPass)
        public
    {
        if (expPass) {
            vm.expectEmit(true, true, true, true);
            emit ChannelOpenAck(address(le.receiver), le.channelId);
        }
        dispatcherProxy.channelOpenAck(
            le.receiver,
            CounterParty(le.portId, le.channelId, le.versionCall),
            le.connectionHops,
            s.ordering,
            s.feeEnabled,
            re,
            s.proof
        );
    }

    /**
     * @dev Step-4 of the 4-step handshake to open an IBC channel.
     * @param le Local end settings for the channel.
     * @param re Remote end settings for the channel.
     * @param s Channel handshake settings.
     * @param expPass Expected pass status of the operation.
     * If expPass is false, `vm.expectRevert` should be called before this function.
     */
    function channelOpenConfirm(
        LocalEnd memory le,
        CounterParty memory re,
        ChannelHandshakeSetting memory s,
        bool expPass
    ) public {
        if (expPass) {
            vm.expectEmit(true, true, true, true);
            emit ChannelOpenConfirm(address(le.receiver), le.channelId);
        }
        dispatcherProxy.channelOpenConfirm(
            le.receiver,
            CounterParty(le.portId, le.channelId, le.versionCall),
            le.connectionHops,
            s.ordering,
            s.feeEnabled,
            re,
            s.proof
        );
    }

    // A helper function to expect revert with message when contract is called with non-owner.
    // Error msg is defined by OpenZeppelin Ownable contract.
    function expectRevertNonOwner() internal {
        vm.expectRevert("Ownable: caller is not the owner");
    }

    function ackToBytes(AckPacket memory ack) public pure returns (bytes memory) {
        return ack.success
            ? abi.encodePacked('{"result":"', Base64.encode(ack.data), '"}')
            : abi.encodePacked('{"error":"', ack.data, '"}');
    }

    // Store connection in channelid to connection mapping using store
    function _storeChannelidToConnectionMapping(bytes32 channelId, bytes32 connection) internal {
        bytes32 chanIdToConnectionMapping = keccak256(abi.encode(channelId, uint32(160)));
        vm.store(address(dispatcherProxy), chanIdToConnectionMapping, connection);
    }

    // Plant a fake packet commitment so the ack checks go through
    // Stdstore doesn't work for proxies so we have to use store
    // use "forge inspect Dispatcher storage" to find the nested mapping slot
    function _storeSendPacketCommitment(address receiver, bytes32 channelId, uint64 sequence, uint64 commitment)
        internal
    {
        bytes32 slot1 = keccak256(abi.encode(receiver, uint32(SEND_PACKET_COMMITMENT_STARTING_SLOT)));
        bytes32 slot2 = keccak256(abi.encode(channelId, slot1));
        bytes32 slot3 = keccak256(abi.encode(sequence, slot2));
        vm.store(address(dispatcherProxy), slot3, bytes32(uint256(commitment)));
    }

    // Store connection in channelid to connection mapping using store
    function _getConnectiontoClientIdMapping(string memory connection) internal view returns (uint256 clientId) {
        bytes32 clientIdSlot = keccak256(abi.encode(connection, CONNECTION_TO_CLIENT_ID_STARTING_SLOT));
        clientId = uint256(vm.load(address(dispatcherProxy), clientIdSlot));
    }
}
