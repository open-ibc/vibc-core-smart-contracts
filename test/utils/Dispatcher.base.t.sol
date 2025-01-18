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
import {stdStorage, StdStorage} from "forge-std/Test.sol";
import {FeeVault} from "../../contracts/core/FeeVault.sol";
import {SequencerSoloClient} from "../../contracts/core/SequencerSoloClient.sol";
import {ISignatureVerifier} from "../../contracts/interfaces/ISignatureVerifier.sol";
import {SequencerSignatureVerifier} from "../../contracts/core/SequencerSignatureVerifier.sol";
import {IFeeVault} from "../../contracts/interfaces/IFeeVault.sol";

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
    using stdStorage for StdStorage;

    event SendPacketFeeDeposited(
        bytes32 indexed channelId, uint64 indexed sequence, uint256[2] gasLimits, uint256[2] gasPrices
    );
    event OpenChannelFeeDeposited(
        address sourceAddress,
        string version,
        ChannelOrder ordering,
        string[] connectionHops,
        string counterpartyPortId,
        uint256 fees
    );

    uint32 CONNECTION_TO_CLIENT_ID_STARTING_SLOT = 260;
    uint32 SEND_PACKET_COMMITMENT_STARTING_SLOT = 255;
    uint32 CHANNEL_ID_TO_CONNECTION_STARTING_SLOT = 258;
    uint64 UINT64_MAX = 18_446_744_073_709_551_615;
    bytes32 PEPTIDE_CHAIN_ID = bytes32(uint256(444));

    Height ZERO_HEIGHT = Height(0, 0);
    uint64 maxTimeout = UINT64_MAX;

    ILightClient opLightClient = new OptimisticLightClient(1800, opProofVerifier, l1BlockProvider);

    ISignatureVerifier sequencerSignatureVerifier = new SequencerSignatureVerifier(sequencer, PEPTIDE_CHAIN_ID);
    ILightClient sequencerLightClient = new SequencerSoloClient(sequencerSignatureVerifier, l1BlockProvider);
    ILightClient dummyLightClient = new DummyLightClient();
    IFeeVault feeVault = new FeeVault();

    IDispatcher public dispatcherProxy;
    Dispatcher public dispatcherImplementation;
    string portPrefix = "polyibc.eth.";
    string[] connectionHops = ["connection-1", "connection-2"];

    uint256 BASE_GAS_LIMIT = 1_000_000;
    uint256 BASE_GAS_PRICE = 50 gwei;
    uint256[2] public sendPacketGasLimit = [BASE_GAS_LIMIT, BASE_GAS_LIMIT];
    uint256[2] public sendPacketGasPrice = [BASE_GAS_PRICE, BASE_GAS_PRICE];
    uint256 public totalSendPacketFees = BASE_GAS_LIMIT * BASE_GAS_PRICE * 2;

    uint256 public totalOpenChannelFees = BASE_GAS_LIMIT * BASE_GAS_PRICE * 3; // The msg.value sent in a
        // channelOpenInit call

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
    function channelOpenInit(LocalEnd memory le, ChannelEnd memory re, ChannelHandshakeSetting memory s, bool expPass)
        public
    {
        vm.startPrank(address(le.receiver));
        if (expPass) {
            vm.expectEmit(true, true, true, true);
            emit ChannelOpenInit(
                address(le.receiver), le.versionExpected, s.ordering, s.feeEnabled, le.connectionHops, re.portId
            );
        }
        dispatcherProxy.channelOpenInit(le.versionCall, s.ordering, s.feeEnabled, le.connectionHops, re.portId);
        vm.stopPrank();
    }

    /**
     * @dev Step-1 of the 4-step handshake to open an IBC channel.
     * @param le Local end settings for the channel.
     * @param re Remote end settings for the channel.
     * @param s Channel handshake settings.
     * @param expPass Expected pass status of the operation.
     * If expPass is false, `vm.expectRevert` should be called before this function.
     */
    function channelOpenInitWithFee(
        LocalEnd memory le,
        ChannelEnd memory re,
        ChannelHandshakeSetting memory s,
        bool expPass
    ) public {
        uint256 beforeBalance = address(feeVault).balance;
        vm.deal(address(le.receiver), totalOpenChannelFees);

        vm.startPrank(address(le.receiver));
        if (expPass) {
            vm.expectEmit(true, true, true, true);
            emit ChannelOpenInit(
                address(le.receiver), le.versionExpected, s.ordering, s.feeEnabled, le.connectionHops, re.portId
            );
        }
        dispatcherProxy.channelOpenInit(le.versionCall, s.ordering, s.feeEnabled, le.connectionHops, re.portId);
        if (expPass) {
            vm.expectEmit(true, true, true, true, address(feeVault));
            emit OpenChannelFeeDeposited(
                address(le.receiver), le.versionCall, s.ordering, connectionHops, re.portId, totalOpenChannelFees
            );
        }
        feeVault.depositOpenChannelFee{value: totalOpenChannelFees}(
            address(le.receiver), le.versionCall, s.ordering, le.connectionHops, re.portId
        );

        assertEq(address(feeVault).balance, beforeBalance + totalOpenChannelFees);

        vm.stopPrank();
    }

    /**
     * @dev Step-2 of the 4-step handshake to open an IBC channel.
     * @param le Local end settings for the channel.
     * @param re Remote end settings for the channel.
     * @param s Channel handshake settings.
     * @param expPass Expected pass status of the operation.
     * If expPass is false, `vm.expectRevert` should be called before this function.
     */
    function channelOpenTry(LocalEnd memory le, ChannelEnd memory re, ChannelHandshakeSetting memory s, bool expPass)
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
        ChannelEnd memory cp = ChannelEnd(re.portId, re.channelId, re.version);
        dispatcherProxy.channelOpenTry(
            ChannelEnd(le.portId, le.channelId, le.versionCall),
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
    function channelOpenAck(LocalEnd memory le, ChannelEnd memory re, ChannelHandshakeSetting memory s, bool expPass)
        public
    {
        if (expPass) {
            vm.expectEmit(true, true, true, true);
            emit ChannelOpenAck(address(le.receiver), le.channelId);
        }
        dispatcherProxy.channelOpenAck(
            ChannelEnd(le.portId, le.channelId, le.versionCall),
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
        ChannelEnd memory re,
        ChannelHandshakeSetting memory s,
        bool expPass
    ) public {
        if (expPass) {
            vm.expectEmit(true, true, true, true);
            emit ChannelOpenConfirm(address(le.receiver), le.channelId);
        }
        dispatcherProxy.channelOpenConfirm(
            ChannelEnd(le.portId, le.channelId, le.versionCall),
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

    // genAckPacket generates an ack packet for the given packet sequence
    function genAckPacket(string memory packetSeq) internal pure returns (bytes memory) {
        return ackToBytes(AckPacket(true, bytes(packetSeq)));
    }

    // Store connection in channelid to connection mapping using store
    function _storeChannelidToConnectionMapping(bytes32 channelId, bytes32 connection) internal {
        bytes32 chanIdToConnectionMapping = keccak256(abi.encode(channelId, CHANNEL_ID_TO_CONNECTION_STARTING_SLOT));
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
    function _getConnectiontoClientIdMapping(string memory connection) internal view returns (address clientId) {
        bytes32 clientIdSlot =
            keccak256(abi.encodePacked(bytes(connection), uint256(CONNECTION_TO_CLIENT_ID_STARTING_SLOT)));
        clientId = address(uint160(uint256(vm.load(address(dispatcherProxy), clientIdSlot))));
    }

    function load_proof(string memory filepath, address lightClient) internal returns (Ics23Proof memory) {
        (bytes32 apphash, Ics23Proof memory proof) =
            abi.decode(vm.parseBytes(vm.readFile(string.concat(rootDir, filepath))), (bytes32, Ics23Proof));

        // this loads the app hash we got from the testing data into the consensus state manager internals
        // at the height it's supposed to go. That is, a block less than where the proof was generated from.
        stdstore.target(lightClient).sig("consensusStates(uint256)").with_key(proof.height - 1).checked_write(apphash);
        // trick the fraud time window check
        vm.warp(block.timestamp + 1);

        return proof;
    }
}
