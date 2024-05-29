// SPDX-License-Identifier: Apache-2.0
/*
 * Copyright 2024, Polymer Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
pragma solidity 0.8.15;

import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";
import {UniversalPacket} from "./Ibc.sol";
import {IBCErrors} from "./IbcErrors.sol";

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
    function hexStrToAddress(string memory hexStr) external pure returns (address addr) {
        bytes memory hexBytes = bytes(hexStr);
        if (bytes(hexStr).length != 40) {
            revert IBCErrors.invalidHexStringLength(); // Addresses must always be 20 bytes long; equal to 40 nibbles
        }

        uint256 total = 0;
        uint256 base = 1;
        uint256 i = 40;
        while (i > 0) {
            i--;
            uint256 digit;
            // Convert ASCII to integer value
            if (uint8(hexBytes[i]) >= 48 && uint8(hexBytes[i]) <= 57) {
                /**
                 * This triggers if hexBytes[i] is equal to '0' to '9'
                 * '0' - '9' are 48-57 in ascii, and we want to map this into 0-9, so we subtract 48
                 */
                digit = uint160(uint8(hexBytes[i]) - 48);
            } else if (uint8(hexBytes[i]) >= 65 && uint8(hexBytes[i]) <= 70) {
                /**
                 * This triggers if hexBytes[i] is equal to 'A' to 'F'
                 * 'A' - 'F' are 65-70 in ascii, and we want to map this into 0xa-0xf (equal to 10-15), so we
                 * sutract 55
                 */
                digit = uint160(uint8(hexBytes[i]) - 55);
            } else if (uint8(hexBytes[i]) >= 97 && uint8(hexBytes[i]) <= 102) {
                /**
                 * This triggers if hexBytes[i] is equal to 'a' to 'f'
                 * 'a' to 'f' are 97-102 in ascii, and we want to amp this into 0xa-0xf (equal to 10-15), so we
                 * sutract 87
                 */
                digit = uint160(uint8(hexBytes[i]) - 87);
            } else {
                revert IBCErrors.invalidCharacter();
            }
            total += digit * base;
            base *= 16;
        }

        addr = address(uint160(total));
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
