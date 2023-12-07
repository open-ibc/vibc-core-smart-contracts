// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import "forge-std/Test.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "../contracts/Ibc.sol";
import "../contracts/Dispatcher.sol";
import "../contracts/Verifier.sol";
import {UniversalChannelHandler} from "../contracts/UniversalChannelHandler.sol";
import {Mars} from "../contracts/Mars.sol";
import {Earth} from "../contracts/Earth.sol";
import "../contracts/DummyConsensusStateManager.sol";

struct ChannelSetting {
    ChannelOrder ordering;
    string version;
    bool feeEnabled;
    Proof proof;
}

// A test contract that keeps two types of dApps, 1. regular IBC-enabled dApp, 2. universal channel dApp
contract VirtualChain is Test, IbcEventsEmitter {
    Dispatcher public dispatcher;
    UniversalChannelHandler public ucHandler;

    Mars public mars;
    Earth public earth;
    mapping(address => mapping(address => bytes32)) public channelIds;
    mapping(address => mapping(bytes32 => Channel)) public channels;
    mapping(address => string) public portIds;

    string[] public connectionHops;
    uint256 _seed;

    // seed is used to generate unique channelIds and connectionIds;
    // it should be unique for each VirtualChain instance
    // Ports are initialized with a prefix, e.g. "polyibc.eth1.", which are only used for counterparty chains
    // ChannelIds are not initialized until channel handshake is started
    constructor(uint256 seed, string memory portPrefix) {
        _seed = seed;
        dispatcher = new Dispatcher(new Verifier(), portPrefix, new DummyConsensusStateManager());
        ucHandler = new UniversalChannelHandler(dispatcher);
        dispatcher.setUniversalChannelHandler(ucHandler);

        mars = new Mars(dispatcher);
        earth = new Earth(dispatcher);
        // initialize portIds for counterparty chains
        address[3] memory portContracts = [address(ucHandler), address(mars), address(earth)];
        for (uint256 i = 0; i < portContracts.length; i++) {
            portIds[portContracts[i]] = string(abi.encodePacked(portPrefix, toHexStr(portContracts[i])));
        }
        connectionHops = new string[](2);
        connectionHops[0] = newConnectionId();
        connectionHops[1] = newConnectionId();
    }

    // expectedChannel returns a Channel struct with expected values
    // this is used to verify on its counterpart chain
    function expectedChannel(
        address localAddr,
        bytes32 localChanId,
        string[] memory counterpartyConnectionHops,
        ChannelSetting memory setting
    ) public view returns (Channel memory) {
        return Channel(
            setting.version,
            setting.ordering,
            setting.feeEnabled,
            counterpartyConnectionHops,
            portIds[localAddr],
            localChanId
        );
    }

    function getConnectionHops() external view returns (string[] memory) {
        return connectionHops;
    }

    // Assign new channelIds to both ends of the channel
    function assignChanelIds(IbcChannelHandler localEnd, IbcChannelHandler remoteEnd, VirtualChain remoteChain)
        external
    {
        bytes32 localChannelId = this.newChannelId();
        bytes32 remoteChannelId = remoteChain.newChannelId();
        // save channelIds on each virtual chain
        this.setChannelId(localEnd, remoteEnd, localChannelId);
        remoteChain.setChannelId(remoteEnd, localEnd, remoteChannelId);
    }

    // finishHandshake is a helper function to finish the 4-step handshake
    // @arg localEnd: the local end of the channel on this virtual chain
    // @arg remoteChain: the remote virtual chain
    // @arg remoteEnd: the remote end of the channel on the other virtual chain
    // @arg setting: the channel handshake setting
    // @dev: Successfully created channelIds and channels will be set on both virtual chains's channelIds and channels
    function finishHandshake(
        IbcChannelHandler localEnd,
        VirtualChain remoteChain,
        IbcChannelHandler remoteEnd,
        ChannelSetting memory setting
    ) external {
        this.assignChanelIds(localEnd, remoteEnd, remoteChain);

        // localEnd initiates the first step in 4-step handshake
        vm.prank(address(this));
        this.channelOpenInit(localEnd, remoteChain, remoteEnd, setting, true); // step-1

        vm.prank(address(remoteChain));
        remoteChain.channelOpenTry(remoteEnd, this, localEnd, setting, true); // step-2

        vm.prank(address(this));
        this.channelOpenAckOrConfirm(localEnd, remoteChain, remoteEnd, setting, true); // step-3

        vm.prank(address(remoteChain));
        remoteChain.channelOpenAckOrConfirm(remoteEnd, this, localEnd, setting, true); // step-4
    }

    function channelOpenInit(
        IbcChannelHandler localEnd,
        VirtualChain remoteChain,
        IbcChannelHandler remoteEnd,
        ChannelSetting memory setting,
        bool expPass
    ) external {
        string memory cpPortId = remoteChain.portIds(address(remoteEnd));
        require(bytes(cpPortId).length > 0, "channelOpenTry: portId does not exist");

        // set dispatcher's msg.sender to this function's msg.sender
        vm.prank(msg.sender);

        if (expPass) {
            vm.expectEmit(true, true, true, true);
            emit OpenIbcChannel(
                address(localEnd),
                setting.version,
                setting.ordering,
                setting.feeEnabled,
                connectionHops,
                remoteChain.portIds(address(remoteEnd)),
                bytes32(0)
            );
        }
        dispatcher.openIbcChannel(
            localEnd,
            setting.version,
            setting.ordering,
            setting.feeEnabled,
            connectionHops,
            // counterparty channelId and version are not known at this point
            CounterParty(cpPortId, bytes32(0), ""),
            setting.proof
        );
    }

    function channelOpenTry(
        IbcChannelHandler localEnd,
        VirtualChain remoteChain,
        IbcChannelHandler remoteEnd,
        ChannelSetting memory setting,
        bool expPass
    ) external {
        bytes32 cpChanId = remoteChain.channelIds(address(remoteEnd), address(localEnd));
        require(cpChanId != bytes32(0), "channelOpenTry: channel does not exist");

        string memory cpPortId = remoteChain.portIds(address(remoteEnd));
        require(bytes(cpPortId).length > 0, "channelOpenTry: portId does not exist");

        // set dispatcher's msg.sender to this function's msg.sender
        vm.prank(msg.sender);

        if (expPass) {
            vm.expectEmit(true, true, true, true);
            emit OpenIbcChannel(
                address(localEnd),
                setting.version,
                setting.ordering,
                setting.feeEnabled,
                connectionHops,
                cpPortId,
                cpChanId
            );
        }
        dispatcher.openIbcChannel(
            localEnd,
            setting.version,
            setting.ordering,
            setting.feeEnabled,
            connectionHops,
            CounterParty(cpPortId, cpChanId, setting.version),
            setting.proof
        );
    }

    function channelOpenAckOrConfirm(
        IbcChannelHandler localEnd,
        VirtualChain remoteChain,
        IbcChannelHandler remoteEnd,
        ChannelSetting memory setting,
        bool expPass
    ) external {
        // local channel Id must be known
        bytes32 chanId = channelIds[address(localEnd)][address(remoteEnd)];
        require(chanId != bytes32(0), "channelOpenAckOrConfirm: channel does not exist");

        bytes32 cpChanId = remoteChain.channelIds(address(remoteEnd), address(localEnd));
        require(cpChanId != bytes32(0), "channelOpenAckOrConfirm: channel does not exist");

        string memory cpPortId = remoteChain.portIds(address(remoteEnd));
        require(bytes(cpPortId).length > 0, "channelOpenAckOrConfirm: counterparty portId does not exist");

        // set dispatcher's msg.sender to this function's msg.sender
        vm.prank(msg.sender);

        if (expPass) {
            vm.expectEmit(true, true, true, true);
            emit ConnectIbcChannel(address(localEnd), chanId);
        }
        dispatcher.connectIbcChannel(
            localEnd,
            chanId,
            connectionHops,
            setting.ordering,
            setting.feeEnabled,
            cpPortId,
            cpChanId,
            setting.version,
            setting.proof
        );
    }

    // Converts a local dApp address on this virtual chain to a Counterparty struct for a remote chain
    function localEndToCounterparty(address localEnd) external view returns (CounterParty memory) {
        return CounterParty(portIds[localEnd], channelIds[localEnd][address(this)], "");
    }

    function setChannelId(IbcChannelHandler localEnd, IbcChannelHandler remoteEnd, bytes32 channelId) external {
        channelIds[address(localEnd)][address(remoteEnd)] = channelId;
    }

    function newChannelId() external returns (bytes32) {
        bytes memory channelId = abi.encodePacked("channel-", Strings.toString(_seed));
        _seed += 1;
        return bytes32(channelId);
    }

    function newConnectionId() internal returns (string memory) {
        string memory connectionId = string(abi.encodePacked("connection-", Strings.toString(_seed)));
        _seed += 1;
        return connectionId;
    }

    // convert an address to its hex string, but without 0x prefix
    function toHexStr(address addr) internal pure returns (bytes memory) {
        bytes memory addrWithPrefix = abi.encodePacked(Strings.toHexString(addr));
        bytes memory addrWithoutPrefix = new bytes(addrWithPrefix.length - 2);
        for (uint256 i = 0; i < addrWithoutPrefix.length; i++) {
            addrWithoutPrefix[i] = addrWithPrefix[i + 2];
        }
        return addrWithoutPrefix;
    }
}
