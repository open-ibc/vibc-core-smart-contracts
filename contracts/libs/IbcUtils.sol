//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import {UniversalPacket, CounterParty, Strings, IBCErrors} from "./Ibc.sol";

// define a library of Ibc utility functions
library IbcUtils {
    error StringTooLong();

    // fromUniversalPacketBytes converts UniversalPacketDataBytes to UniversalPacketData, per how its packed into bytes
    function fromUniversalPacketBytes(bytes calldata data)
        external
        pure
        returns (UniversalPacket memory universalPacketData)
    {
        bytes32 srcPortAddr;
        uint256 mwBitmap;
        bytes32 destPortAddr;
        assembly {
            // Keep reusing 0x0 to move from calldata to return vars
            calldatacopy(0x0, data.offset, 32)
            srcPortAddr := mload(0x0)
            calldatacopy(0x0, add(data.offset, 32), 32)
            mwBitmap := mload(0x0)
            calldatacopy(0x0, add(data.offset, 64), 32)
            destPortAddr := mload(0x0)
        }
        universalPacketData = UniversalPacket(srcPortAddr, uint256(mwBitmap), destPortAddr, data[96:data.length]);
    }

    /**
     * Convert a non-0x-prefixed hex string to an address
     * @param hexStr hex string to convert to address. Note that the hex string must not include a 0x prefix.
     * hexStr is case-insensitive.
     */
    function hexStrToAddress(string memory hexStr) public pure returns (address addr) {
        if (bytes(hexStr).length != 40) {
            revert IBCErrors.invalidHexStringLength();
        }

        bytes memory strBytes = bytes(hexStr);
        bytes memory addrBytes = new bytes(20);

        for (uint256 i = 0; i < 20; i++) {
            uint8 high = uint8(strBytes[i * 2]);
            uint8 low = uint8(strBytes[1 + i * 2]);
            // Convert to lowercase if the character is in uppercase
            if (high >= 65 && high <= 90) {
                high += 32;
            }
            if (low >= 65 && low <= 90) {
                low += 32;
            }
            uint8 digit = (high - (high >= 97 ? 87 : 48)) * 16 + (low - (low >= 97 ? 87 : 48));
            addrBytes[i] = bytes1(digit);
        }

        assembly {
            addr := mload(add(addrBytes, 20))
        }

        return addr;
    }

    // For XXXX => vIBC direction, SC needs to verify the proof of membership of TRY_PENDING
    // For vIBC initiated channel, SC doesn't need to verify any proof, and these should be all empty
    function isChannelOpenTry(CounterParty calldata counterparty) public pure returns (bool open) {
        if (counterparty.channelId == bytes32(0) && bytes(counterparty.version).length == 0) {
            return false;
            // ChanOpenInit with unknow conterparty
        } else if (counterparty.channelId != bytes32(0) && bytes(counterparty.version).length != 0) {
            // this is the ChanOpenTry; counterparty must not be zero-value
            return true;
        } else {
            revert IBCErrors.invalidCounterParty();
        }
    }

    function toUniversalPacketBytes(UniversalPacket memory data) internal pure returns (bytes memory packetBytes) {
        packetBytes = bytes.concat(data.srcPortAddr, bytes32(data.mwBitmap), data.destPortAddr, data.appData);
    }

    // addressToPortId converts an address to a port ID
    function addressToPortId(string memory portPrefix, address addr) internal pure returns (string memory portId) {
        portId = string(abi.encodePacked(portPrefix, toHexStr(addr)));
    }

    // convert an address to its hex string, but without 0x prefix
    function toHexStr(address addr) internal pure returns (bytes memory hexStr) {
        bytes memory addrWithPrefix = abi.encodePacked(Strings.toHexString(addr));
        bytes memory addrWithoutPrefix = new bytes(addrWithPrefix.length - 2);
        for (uint256 i = 0; i < addrWithoutPrefix.length; i++) {
            addrWithoutPrefix[i] = addrWithPrefix[i + 2];
        }
        hexStr = addrWithoutPrefix;
    }

    // toAddress converts a bytes32 to an address
    function toAddress(bytes32 b) internal pure returns (address out) {
        out = address(uint160(uint256(b)));
    }

    // toBytes32 converts an address to a bytes32
    function toBytes32(address a) internal pure returns (bytes32 out) {
        out = bytes32(uint256(uint160(a)));
    }

    function toBytes32(string memory s) internal pure returns (bytes32 result) {
        bytes memory b = bytes(s);
        if (b.length > 32) revert StringTooLong();

        assembly {
            result := mload(add(b, 32))
        }
    }
}
