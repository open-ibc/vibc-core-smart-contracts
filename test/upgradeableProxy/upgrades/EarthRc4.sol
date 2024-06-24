//SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.6.0 ^0.8.0 ^0.8.15 ^0.8.9;

// contracts/interfaces/ProofVerifier.sol

struct OpIcs23ProofPath {
    bytes prefix;
    bytes suffix;
}

struct OpIcs23Proof {
    OpIcs23ProofPath[] path;
    bytes key;
    bytes value;
    bytes prefix;
}

// the Ics23 proof related structs are used to do membership verification. These are not the actual Ics23
// format but a "solidity friendly" version of it - data is the same just packaged differently
struct Ics23Proof {
    OpIcs23Proof[] proof;
    uint256 height;
}

// This is the proof we use to verify the apphash (state) updates.
struct OpL2StateProof {
    bytes[] accountProof;
    bytes[] outputRootProof;
    bytes32 l2OutputProposalKey;
    bytes32 l2BlockHash;
}

// The `header` field is a list of RLP encoded L1 header fields. Both stateRoot and number are not
// encoded for easy usage. They must match with their RLP encoded counterparty versions.
struct L1Header {
    bytes[] header;
    bytes32 stateRoot;
    uint64 number;
}

interface ProofVerifier {
    error InvalidL1BlockNumber();
    error InvalidL1BlockHash();
    error InvalidRLPEncodedL1BlockNumber();
    error InvalidRLPEncodedL1StateRoot();
    error InvalidAppHash();
    error InvalidProofKey();
    error InvalidProofValue();
    error InvalidPacketProof();
    error InvalidIbcStateProof();
    error MethodNotImplemented();

    /**
     * @dev verifies if a state update (apphash) is valid, given the provided proofs.
     *      Reverts in case of failure.
     *
     * @param l1header RLP "encoded" version of the L1 header that matches with the trusted hash and number
     * @param proof l2 state proof. It includes the keys, hashes and storage proofs required to verify the app hash
     * @param appHash  l2 app hash (state root) to be verified
     * @param trustedL1BlockHash trusted L1 block hash. Provided L1 header must match with it.
     * @param trustedL1BlockNumber trusted L1 block number. Provided L1 header must match with it.
     */
    function verifyStateUpdate(
        L1Header calldata l1header,
        OpL2StateProof calldata proof,
        bytes32 appHash,
        bytes32 trustedL1BlockHash,
        uint64 trustedL1BlockNumber
    ) external view;

    /**
     * @dev verifies the provided ICS23 proof given the trusted app hash. Reverts in case of failure.
     *
     * @param appHash trusted l2 app hash (state root)
     * @param key key to be proven
     * @param value value to be proven
     * @param proof ICS23 membership proof
     */
    function verifyMembership(bytes32 appHash, bytes calldata key, bytes calldata value, Ics23Proof calldata proof)
        external
        pure;

    /**
     * @dev verifies the provided ICS23 proof given the trusted app hash. Reverts in case of failure.
     *
     * @param appHash trusted l2 app hash (state root)
     * @param key key to be proven non-existing
     * @param proof ICS23 non-membership proof
     */
    function verifyNonMembership(bytes32 appHash, bytes calldata key, Ics23Proof calldata proof) external pure;
}

// lib/base64/base64.sol

/// @title Base64
/// @author Brecht Devos - <brecht@loopring.org>
/// @notice Provides functions for encoding/decoding base64
library Base64 {
    string internal constant TABLE_ENCODE = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";
    bytes internal constant TABLE_DECODE = hex"0000000000000000000000000000000000000000000000000000000000000000"
        hex"00000000000000000000003e0000003f3435363738393a3b3c3d000000000000"
        hex"00000102030405060708090a0b0c0d0e0f101112131415161718190000000000"
        hex"001a1b1c1d1e1f202122232425262728292a2b2c2d2e2f303132330000000000";

    function encode(bytes memory data) internal pure returns (string memory) {
        if (data.length == 0) return "";

        // load the table into memory
        string memory table = TABLE_ENCODE;

        // multiply by 4/3 rounded up
        uint256 encodedLen = 4 * ((data.length + 2) / 3);

        // add some extra buffer at the end required for the writing
        string memory result = new string(encodedLen + 32);

        assembly {
            // set the actual output length
            mstore(result, encodedLen)

            // prepare the lookup table
            let tablePtr := add(table, 1)

            // input ptr
            let dataPtr := data
            let endPtr := add(dataPtr, mload(data))

            // result ptr, jump over length
            let resultPtr := add(result, 32)

            // run over the input, 3 bytes at a time
            for {} lt(dataPtr, endPtr) {} {
                // read 3 bytes
                dataPtr := add(dataPtr, 3)
                let input := mload(dataPtr)

                // write 4 characters
                mstore8(resultPtr, mload(add(tablePtr, and(shr(18, input), 0x3F))))
                resultPtr := add(resultPtr, 1)
                mstore8(resultPtr, mload(add(tablePtr, and(shr(12, input), 0x3F))))
                resultPtr := add(resultPtr, 1)
                mstore8(resultPtr, mload(add(tablePtr, and(shr(6, input), 0x3F))))
                resultPtr := add(resultPtr, 1)
                mstore8(resultPtr, mload(add(tablePtr, and(input, 0x3F))))
                resultPtr := add(resultPtr, 1)
            }

            // padding with '='
            switch mod(mload(data), 3)
            case 1 { mstore(sub(resultPtr, 2), shl(240, 0x3d3d)) }
            case 2 { mstore(sub(resultPtr, 1), shl(248, 0x3d)) }
        }

        return result;
    }

    function decode(string memory _data) internal pure returns (bytes memory) {
        bytes memory data = bytes(_data);

        if (data.length == 0) return new bytes(0);
        require(data.length % 4 == 0, "invalid base64 decoder input");

        // load the table into memory
        bytes memory table = TABLE_DECODE;

        // every 4 characters represent 3 bytes
        uint256 decodedLen = (data.length / 4) * 3;

        // add some extra buffer at the end required for the writing
        bytes memory result = new bytes(decodedLen + 32);

        assembly {
            // padding with '='
            let lastBytes := mload(add(data, mload(data)))
            if eq(and(lastBytes, 0xFF), 0x3d) {
                decodedLen := sub(decodedLen, 1)
                if eq(and(lastBytes, 0xFFFF), 0x3d3d) { decodedLen := sub(decodedLen, 1) }
            }

            // set the actual output length
            mstore(result, decodedLen)

            // prepare the lookup table
            let tablePtr := add(table, 1)

            // input ptr
            let dataPtr := data
            let endPtr := add(dataPtr, mload(data))

            // result ptr, jump over length
            let resultPtr := add(result, 32)

            // run over the input, 4 characters at a time
            for {} lt(dataPtr, endPtr) {} {
                // read 4 characters
                dataPtr := add(dataPtr, 4)
                let input := mload(dataPtr)

                // write 3 bytes
                let output :=
                    add(
                        add(
                            shl(18, and(mload(add(tablePtr, and(shr(24, input), 0xFF))), 0xFF)),
                            shl(12, and(mload(add(tablePtr, and(shr(16, input), 0xFF))), 0xFF))
                        ),
                        add(
                            shl(6, and(mload(add(tablePtr, and(shr(8, input), 0xFF))), 0xFF)),
                            and(mload(add(tablePtr, and(input, 0xFF))), 0xFF)
                        )
                    )
                mstore(resultPtr, shl(232, output))
                resultPtr := add(resultPtr, 3)
            }
        }

        return result;
    }
}

// lib/openzeppelin-contracts/contracts/utils/Context.sol

// OpenZeppelin Contracts v4.4.1 (utils/Context.sol)

/**
 * @dev Provides information about the current execution context, including the
 * sender of the transaction and its data. While these are generally available
 * via msg.sender and msg.data, they should not be accessed in such a direct
 * manner, since when dealing with meta-transactions the account sending and
 * paying for execution may not be the actual sender (as far as an application
 * is concerned).
 *
 * This contract is only required for intermediate, library-like contracts.
 */
abstract contract Context {
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }

    function _msgData() internal view virtual returns (bytes calldata) {
        return msg.data;
    }
}

// lib/openzeppelin-contracts/contracts/utils/math/Math.sol

// OpenZeppelin Contracts (last updated v4.9.0) (utils/math/Math.sol)

/**
 * @dev Standard math utilities missing in the Solidity language.
 */
library Math {
    enum Rounding {
        Down, // Toward negative infinity
        Up, // Toward infinity
        Zero // Toward zero

    }

    /**
     * @dev Returns the largest of two numbers.
     */
    function max(uint256 a, uint256 b) internal pure returns (uint256) {
        return a > b ? a : b;
    }

    /**
     * @dev Returns the smallest of two numbers.
     */
    function min(uint256 a, uint256 b) internal pure returns (uint256) {
        return a < b ? a : b;
    }

    /**
     * @dev Returns the average of two numbers. The result is rounded towards
     * zero.
     */
    function average(uint256 a, uint256 b) internal pure returns (uint256) {
        // (a + b) / 2 can overflow.
        return (a & b) + (a ^ b) / 2;
    }

    /**
     * @dev Returns the ceiling of the division of two numbers.
     *
     * This differs from standard division with `/` in that it rounds up instead
     * of rounding down.
     */
    function ceilDiv(uint256 a, uint256 b) internal pure returns (uint256) {
        // (a + b - 1) / b can overflow on addition, so we distribute.
        return a == 0 ? 0 : (a - 1) / b + 1;
    }

    /**
     * @notice Calculates floor(x * y / denominator) with full precision. Throws if result overflows a uint256 or
     * denominator == 0
     * @dev Original credit to Remco Bloemen under MIT license (https://xn--2-umb.com/21/muldiv)
     * with further edits by Uniswap Labs also under MIT license.
     */
    function mulDiv(uint256 x, uint256 y, uint256 denominator) internal pure returns (uint256 result) {
        unchecked {
            // 512-bit multiply [prod1 prod0] = x * y. Compute the product mod 2^256 and mod 2^256 - 1, then use
            // use the Chinese Remainder Theorem to reconstruct the 512 bit result. The result is stored in two 256
            // variables such that product = prod1 * 2^256 + prod0.
            uint256 prod0; // Least significant 256 bits of the product
            uint256 prod1; // Most significant 256 bits of the product
            assembly {
                let mm := mulmod(x, y, not(0))
                prod0 := mul(x, y)
                prod1 := sub(sub(mm, prod0), lt(mm, prod0))
            }

            // Handle non-overflow cases, 256 by 256 division.
            if (prod1 == 0) {
                // Solidity will revert if denominator == 0, unlike the div opcode on its own.
                // The surrounding unchecked block does not change this fact.
                // See https://docs.soliditylang.org/en/latest/control-structures.html#checked-or-unchecked-arithmetic.
                return prod0 / denominator;
            }

            // Make sure the result is less than 2^256. Also prevents denominator == 0.
            require(denominator > prod1, "Math: mulDiv overflow");

            ///////////////////////////////////////////////
            // 512 by 256 division.
            ///////////////////////////////////////////////

            // Make division exact by subtracting the remainder from [prod1 prod0].
            uint256 remainder;
            assembly {
                // Compute remainder using mulmod.
                remainder := mulmod(x, y, denominator)

                // Subtract 256 bit number from 512 bit number.
                prod1 := sub(prod1, gt(remainder, prod0))
                prod0 := sub(prod0, remainder)
            }

            // Factor powers of two out of denominator and compute largest power of two divisor of denominator. Always
            // >= 1.
            // See https://cs.stackexchange.com/q/138556/92363.

            // Does not overflow because the denominator cannot be zero at this stage in the function.
            uint256 twos = denominator & (~denominator + 1);
            assembly {
                // Divide denominator by twos.
                denominator := div(denominator, twos)

                // Divide [prod1 prod0] by twos.
                prod0 := div(prod0, twos)

                // Flip twos such that it is 2^256 / twos. If twos is zero, then it becomes one.
                twos := add(div(sub(0, twos), twos), 1)
            }

            // Shift in bits from prod1 into prod0.
            prod0 |= prod1 * twos;

            // Invert denominator mod 2^256. Now that denominator is an odd number, it has an inverse modulo 2^256 such
            // that denominator * inv = 1 mod 2^256. Compute the inverse by starting with a seed that is correct for
            // four bits. That is, denominator * inv = 1 mod 2^4.
            uint256 inverse = (3 * denominator) ^ 2;

            // Use the Newton-Raphson iteration to improve the precision. Thanks to Hensel's lifting lemma, this also
            // works
            // in modular arithmetic, doubling the correct bits in each step.
            inverse *= 2 - denominator * inverse; // inverse mod 2^8
            inverse *= 2 - denominator * inverse; // inverse mod 2^16
            inverse *= 2 - denominator * inverse; // inverse mod 2^32
            inverse *= 2 - denominator * inverse; // inverse mod 2^64
            inverse *= 2 - denominator * inverse; // inverse mod 2^128
            inverse *= 2 - denominator * inverse; // inverse mod 2^256

            // Because the division is now exact we can divide by multiplying with the modular inverse of denominator.
            // This will give us the correct result modulo 2^256. Since the preconditions guarantee that the outcome is
            // less than 2^256, this is the final result. We don't need to compute the high bits of the result and prod1
            // is no longer required.
            result = prod0 * inverse;
            return result;
        }
    }

    /**
     * @notice Calculates x * y / denominator with full precision, following the selected rounding direction.
     */
    function mulDiv(uint256 x, uint256 y, uint256 denominator, Rounding rounding) internal pure returns (uint256) {
        uint256 result = mulDiv(x, y, denominator);
        if (rounding == Rounding.Up && mulmod(x, y, denominator) > 0) {
            result += 1;
        }
        return result;
    }

    /**
     * @dev Returns the square root of a number. If the number is not a perfect square, the value is rounded down.
     *
     * Inspired by Henry S. Warren, Jr.'s "Hacker's Delight" (Chapter 11).
     */
    function sqrt(uint256 a) internal pure returns (uint256) {
        if (a == 0) {
            return 0;
        }

        // For our first guess, we get the biggest power of 2 which is smaller than the square root of the target.
        //
        // We know that the "msb" (most significant bit) of our target number `a` is a power of 2 such that we have
        // `msb(a) <= a < 2*msb(a)`. This value can be written `msb(a)=2**k` with `k=log2(a)`.
        //
        // This can be rewritten `2**log2(a) <= a < 2**(log2(a) + 1)`
        // → `sqrt(2**k) <= sqrt(a) < sqrt(2**(k+1))`
        // → `2**(k/2) <= sqrt(a) < 2**((k+1)/2) <= 2**(k/2 + 1)`
        //
        // Consequently, `2**(log2(a) / 2)` is a good first approximation of `sqrt(a)` with at least 1 correct bit.
        uint256 result = 1 << (log2(a) >> 1);

        // At this point `result` is an estimation with one bit of precision. We know the true value is a uint128,
        // since it is the square root of a uint256. Newton's method converges quadratically (precision doubles at
        // every iteration). We thus need at most 7 iteration to turn our partial result with one bit of precision
        // into the expected uint128 result.
        unchecked {
            result = (result + a / result) >> 1;
            result = (result + a / result) >> 1;
            result = (result + a / result) >> 1;
            result = (result + a / result) >> 1;
            result = (result + a / result) >> 1;
            result = (result + a / result) >> 1;
            result = (result + a / result) >> 1;
            return min(result, a / result);
        }
    }

    /**
     * @notice Calculates sqrt(a), following the selected rounding direction.
     */
    function sqrt(uint256 a, Rounding rounding) internal pure returns (uint256) {
        unchecked {
            uint256 result = sqrt(a);
            return result + (rounding == Rounding.Up && result * result < a ? 1 : 0);
        }
    }

    /**
     * @dev Return the log in base 2, rounded down, of a positive value.
     * Returns 0 if given 0.
     */
    function log2(uint256 value) internal pure returns (uint256) {
        uint256 result = 0;
        unchecked {
            if (value >> 128 > 0) {
                value >>= 128;
                result += 128;
            }
            if (value >> 64 > 0) {
                value >>= 64;
                result += 64;
            }
            if (value >> 32 > 0) {
                value >>= 32;
                result += 32;
            }
            if (value >> 16 > 0) {
                value >>= 16;
                result += 16;
            }
            if (value >> 8 > 0) {
                value >>= 8;
                result += 8;
            }
            if (value >> 4 > 0) {
                value >>= 4;
                result += 4;
            }
            if (value >> 2 > 0) {
                value >>= 2;
                result += 2;
            }
            if (value >> 1 > 0) {
                result += 1;
            }
        }
        return result;
    }

    /**
     * @dev Return the log in base 2, following the selected rounding direction, of a positive value.
     * Returns 0 if given 0.
     */
    function log2(uint256 value, Rounding rounding) internal pure returns (uint256) {
        unchecked {
            uint256 result = log2(value);
            return result + (rounding == Rounding.Up && 1 << result < value ? 1 : 0);
        }
    }

    /**
     * @dev Return the log in base 10, rounded down, of a positive value.
     * Returns 0 if given 0.
     */
    function log10(uint256 value) internal pure returns (uint256) {
        uint256 result = 0;
        unchecked {
            if (value >= 10 ** 64) {
                value /= 10 ** 64;
                result += 64;
            }
            if (value >= 10 ** 32) {
                value /= 10 ** 32;
                result += 32;
            }
            if (value >= 10 ** 16) {
                value /= 10 ** 16;
                result += 16;
            }
            if (value >= 10 ** 8) {
                value /= 10 ** 8;
                result += 8;
            }
            if (value >= 10 ** 4) {
                value /= 10 ** 4;
                result += 4;
            }
            if (value >= 10 ** 2) {
                value /= 10 ** 2;
                result += 2;
            }
            if (value >= 10 ** 1) {
                result += 1;
            }
        }
        return result;
    }

    /**
     * @dev Return the log in base 10, following the selected rounding direction, of a positive value.
     * Returns 0 if given 0.
     */
    function log10(uint256 value, Rounding rounding) internal pure returns (uint256) {
        unchecked {
            uint256 result = log10(value);
            return result + (rounding == Rounding.Up && 10 ** result < value ? 1 : 0);
        }
    }

    /**
     * @dev Return the log in base 256, rounded down, of a positive value.
     * Returns 0 if given 0.
     *
     * Adding one to the result gives the number of pairs of hex symbols needed to represent `value` as a hex string.
     */
    function log256(uint256 value) internal pure returns (uint256) {
        uint256 result = 0;
        unchecked {
            if (value >> 128 > 0) {
                value >>= 128;
                result += 16;
            }
            if (value >> 64 > 0) {
                value >>= 64;
                result += 8;
            }
            if (value >> 32 > 0) {
                value >>= 32;
                result += 4;
            }
            if (value >> 16 > 0) {
                value >>= 16;
                result += 2;
            }
            if (value >> 8 > 0) {
                result += 1;
            }
        }
        return result;
    }

    /**
     * @dev Return the log in base 256, following the selected rounding direction, of a positive value.
     * Returns 0 if given 0.
     */
    function log256(uint256 value, Rounding rounding) internal pure returns (uint256) {
        unchecked {
            uint256 result = log256(value);
            return result + (rounding == Rounding.Up && 1 << (result << 3) < value ? 1 : 0);
        }
    }
}

// lib/openzeppelin-contracts/contracts/utils/math/SignedMath.sol

// OpenZeppelin Contracts (last updated v4.8.0) (utils/math/SignedMath.sol)

/**
 * @dev Standard signed math utilities missing in the Solidity language.
 */
library SignedMath {
    /**
     * @dev Returns the largest of two signed numbers.
     */
    function max(int256 a, int256 b) internal pure returns (int256) {
        return a > b ? a : b;
    }

    /**
     * @dev Returns the smallest of two signed numbers.
     */
    function min(int256 a, int256 b) internal pure returns (int256) {
        return a < b ? a : b;
    }

    /**
     * @dev Returns the average of two signed numbers without overflow.
     * The result is rounded towards zero.
     */
    function average(int256 a, int256 b) internal pure returns (int256) {
        // Formula from the book "Hacker's Delight"
        int256 x = (a & b) + ((a ^ b) >> 1);
        return x + (int256(uint256(x) >> 255) & (a ^ b));
    }

    /**
     * @dev Returns the absolute unsigned value of a signed value.
     */
    function abs(int256 n) internal pure returns (uint256) {
        unchecked {
            // must be unchecked in order to support `n = type(int256).min`
            return uint256(n >= 0 ? n : -n);
        }
    }
}

// lib/proto/ProtoBufRuntime.sol

/**
 * @title Runtime library for ProtoBuf serialization and/or deserialization.
 * All ProtoBuf generated code will use this library.
 */
library ProtoBufRuntime {
    // Types defined in ProtoBuf
    enum WireType {
        Varint,
        Fixed64,
        LengthDelim,
        StartGroup,
        EndGroup,
        Fixed32
    }
    // Constants for bytes calculation

    uint256 constant WORD_LENGTH = 32;
    uint256 constant HEADER_SIZE_LENGTH_IN_BYTES = 4;
    uint256 constant BYTE_SIZE = 8;
    uint256 constant REMAINING_LENGTH = WORD_LENGTH - HEADER_SIZE_LENGTH_IN_BYTES;
    string constant OVERFLOW_MESSAGE = "length overflow";

    //Storages
    /**
     * @dev Encode to storage location using assembly to save storage space.
     * @param location The location of storage
     * @param encoded The encoded ProtoBuf bytes
     */
    function encodeStorage(bytes storage location, bytes memory encoded) internal {
        /**
         * This code use the first four bytes as size,
         * and then put the rest of `encoded` bytes.
         */
        uint256 length = encoded.length;
        uint256 firstWord;
        uint256 wordLength = WORD_LENGTH;
        uint256 remainingLength = REMAINING_LENGTH;

        assembly {
            firstWord := mload(add(encoded, wordLength))
        }
        firstWord =
            (firstWord >> (BYTE_SIZE * HEADER_SIZE_LENGTH_IN_BYTES)) | (length << (BYTE_SIZE * REMAINING_LENGTH));

        assembly {
            sstore(location.slot, firstWord)
        }

        if (length > REMAINING_LENGTH) {
            length -= REMAINING_LENGTH;
            for (uint256 i = 0; i < ceil(length, WORD_LENGTH); i++) {
                assembly {
                    let offset := add(mul(i, wordLength), remainingLength)
                    let slotIndex := add(i, 1)
                    sstore(add(location.slot, slotIndex), mload(add(add(encoded, wordLength), offset)))
                }
            }
        }
    }

    /**
     * @dev Decode storage location using assembly using the format in `encodeStorage`.
     * @param location The location of storage
     * @return The encoded bytes
     */
    function decodeStorage(bytes storage location) internal view returns (bytes memory) {
        /**
         * This code is to decode the first four bytes as size,
         * and then decode the rest using the decoded size.
         */
        uint256 firstWord;
        uint256 remainingLength = REMAINING_LENGTH;
        uint256 wordLength = WORD_LENGTH;

        assembly {
            firstWord := sload(location.slot)
        }

        uint256 length = firstWord >> (BYTE_SIZE * REMAINING_LENGTH);
        bytes memory encoded = new bytes(length);

        assembly {
            mstore(add(encoded, remainingLength), firstWord)
        }

        if (length > REMAINING_LENGTH) {
            length -= REMAINING_LENGTH;
            for (uint256 i = 0; i < ceil(length, WORD_LENGTH); i++) {
                assembly {
                    let offset := add(mul(i, wordLength), remainingLength)
                    let slotIndex := add(i, 1)
                    mstore(add(add(encoded, wordLength), offset), sload(add(location.slot, slotIndex)))
                }
            }
        }
        return encoded;
    }

    /**
     * @dev Fast memory copy of bytes using assembly.
     * @param src The source memory address
     * @param dest The destination memory address
     * @param len The length of bytes to copy
     */
    function copyBytes(uint256 src, uint256 dest, uint256 len) internal pure {
        if (len == 0) {
            return;
        }

        // Copy word-length chunks while possible
        for (; len > WORD_LENGTH; len -= WORD_LENGTH) {
            assembly {
                mstore(dest, mload(src))
            }
            dest += WORD_LENGTH;
            src += WORD_LENGTH;
        }

        // Copy remaining bytes
        uint256 mask = 256 ** (WORD_LENGTH - len) - 1;
        assembly {
            let srcpart := and(mload(src), not(mask))
            let destpart := and(mload(dest), mask)
            mstore(dest, or(destpart, srcpart))
        }
    }

    /**
     * @dev Use assembly to get memory address.
     * @param r The in-memory bytes array
     * @return The memory address of `r`
     */
    function getMemoryAddress(bytes memory r) internal pure returns (uint256) {
        uint256 addr;
        assembly {
            addr := r
        }
        return addr;
    }

    /**
     * @dev Implement Math function of ceil
     * @param a The denominator
     * @param m The numerator
     * @return r The result of ceil(a/m)
     */
    function ceil(uint256 a, uint256 m) internal pure returns (uint256 r) {
        return (a + m - 1) / m;
    }

    // Decoders
    /**
     * This section of code `_decode_(u)int(32|64)`, `_decode_enum` and `_decode_bool`
     * is to decode ProtoBuf native integers,
     * using the `varint` encoding.
     */

    /**
     * @dev Decode integers
     * @param p The memory offset of `bs`
     * @param bs The bytes array to be decoded
     * @return The decoded integer
     * @return The length of `bs` used to get decoded
     */
    function _decode_uint32(uint256 p, bytes memory bs) internal pure returns (uint32, uint256) {
        (uint256 varint, uint256 sz) = _decode_varint(p, bs);
        return (uint32(varint), sz);
    }

    /**
     * @dev Decode integers
     * @param p The memory offset of `bs`
     * @param bs The bytes array to be decoded
     * @return The decoded integer
     * @return The length of `bs` used to get decoded
     */
    function _decode_uint64(uint256 p, bytes memory bs) internal pure returns (uint64, uint256) {
        (uint256 varint, uint256 sz) = _decode_varint(p, bs);
        return (uint64(varint), sz);
    }

    /**
     * @dev Decode integers
     * @param p The memory offset of `bs`
     * @param bs The bytes array to be decoded
     * @return The decoded integer
     * @return The length of `bs` used to get decoded
     */
    function _decode_int32(uint256 p, bytes memory bs) internal pure returns (int32, uint256) {
        (uint256 varint, uint256 sz) = _decode_varint(p, bs);
        int32 r;
        assembly {
            r := varint
        }
        return (r, sz);
    }

    /**
     * @dev Decode integers
     * @param p The memory offset of `bs`
     * @param bs The bytes array to be decoded
     * @return The decoded integer
     * @return The length of `bs` used to get decoded
     */
    function _decode_int64(uint256 p, bytes memory bs) internal pure returns (int64, uint256) {
        (uint256 varint, uint256 sz) = _decode_varint(p, bs);
        int64 r;
        assembly {
            r := varint
        }
        return (r, sz);
    }

    /**
     * @dev Decode enum
     * @param p The memory offset of `bs`
     * @param bs The bytes array to be decoded
     * @return The decoded enum's integer
     * @return The length of `bs` used to get decoded
     */
    function _decode_enum(uint256 p, bytes memory bs) internal pure returns (int64, uint256) {
        return _decode_int64(p, bs);
    }

    /**
     * @dev Decode enum
     * @param p The memory offset of `bs`
     * @param bs The bytes array to be decoded
     * @return The decoded boolean
     * @return The length of `bs` used to get decoded
     */
    function _decode_bool(uint256 p, bytes memory bs) internal pure returns (bool, uint256) {
        (uint256 varint, uint256 sz) = _decode_varint(p, bs);
        if (varint == 0) {
            return (false, sz);
        }
        return (true, sz);
    }

    /**
     * This section of code `_decode_sint(32|64)`
     * is to decode ProtoBuf native signed integers,
     * using the `zig-zag` encoding.
     */

    /**
     * @dev Decode signed integers
     * @param p The memory offset of `bs`
     * @param bs The bytes array to be decoded
     * @return The decoded integer
     * @return The length of `bs` used to get decoded
     */
    function _decode_sint32(uint256 p, bytes memory bs) internal pure returns (int32, uint256) {
        (int256 varint, uint256 sz) = _decode_varints(p, bs);
        return (int32(varint), sz);
    }

    /**
     * @dev Decode signed integers
     * @param p The memory offset of `bs`
     * @param bs The bytes array to be decoded
     * @return The decoded integer
     * @return The length of `bs` used to get decoded
     */
    function _decode_sint64(uint256 p, bytes memory bs) internal pure returns (int64, uint256) {
        (int256 varint, uint256 sz) = _decode_varints(p, bs);
        return (int64(varint), sz);
    }

    /**
     * @dev Decode string
     * @param p The memory offset of `bs`
     * @param bs The bytes array to be decoded
     * @return The decoded string
     * @return The length of `bs` used to get decoded
     */
    function _decode_string(uint256 p, bytes memory bs) internal pure returns (string memory, uint256) {
        (bytes memory x, uint256 sz) = _decode_lendelim(p, bs);
        return (string(x), sz);
    }

    /**
     * @dev Decode bytes array
     * @param p The memory offset of `bs`
     * @param bs The bytes array to be decoded
     * @return The decoded bytes array
     * @return The length of `bs` used to get decoded
     */
    function _decode_bytes(uint256 p, bytes memory bs) internal pure returns (bytes memory, uint256) {
        return _decode_lendelim(p, bs);
    }

    /**
     * @dev Decode ProtoBuf key
     * @param p The memory offset of `bs`
     * @param bs The bytes array to be decoded
     * @return The decoded field ID
     * @return The decoded WireType specified in ProtoBuf
     * @return The length of `bs` used to get decoded
     */
    function _decode_key(uint256 p, bytes memory bs) internal pure returns (uint256, WireType, uint256) {
        (uint256 x, uint256 n) = _decode_varint(p, bs);
        WireType typeId = WireType(x & 7);
        uint256 fieldId = x / 8;
        return (fieldId, typeId, n);
    }

    /**
     * @dev Decode ProtoBuf varint
     * @param p The memory offset of `bs`
     * @param bs The bytes array to be decoded
     * @return The decoded unsigned integer
     * @return The length of `bs` used to get decoded
     */
    function _decode_varint(uint256 p, bytes memory bs) internal pure returns (uint256, uint256) {
        /**
         * Read a byte.
         * Use the lower 7 bits and shift it to the left,
         * until the most significant bit is 0.
         * Refer to https://developers.google.com/protocol-buffers/docs/encoding
         */
        uint256 x = 0;
        uint256 sz = 0;
        uint256 length = bs.length + WORD_LENGTH;
        assembly {
            let b := 0x80
            p := add(bs, p)
            for {} eq(0x80, and(b, 0x80)) {} {
                if eq(lt(sub(p, bs), length), 0) {
                    mstore(0, 0x08c379a000000000000000000000000000000000000000000000000000000000) //error function
                        // selector
                    mstore(4, 32)
                    mstore(36, 15)
                    mstore(68, 0x6c656e677468206f766572666c6f770000000000000000000000000000000000) // length overflow in
                        // hex
                    revert(0, 83)
                }
                let tmp := mload(p)
                let pos := 0
                for {} and(eq(0x80, and(b, 0x80)), lt(pos, 32)) {} {
                    if eq(lt(sub(p, bs), length), 0) {
                        mstore(0, 0x08c379a000000000000000000000000000000000000000000000000000000000) //error function
                            // selector
                        mstore(4, 32)
                        mstore(36, 15)
                        mstore(68, 0x6c656e677468206f766572666c6f770000000000000000000000000000000000) // length
                            // overflow in hex
                        revert(0, 83)
                    }
                    b := byte(pos, tmp)
                    x := or(x, shl(mul(7, sz), and(0x7f, b)))
                    sz := add(sz, 1)
                    pos := add(pos, 1)
                    p := add(p, 0x01)
                }
            }
        }
        return (x, sz);
    }

    /**
     * @dev Decode ProtoBuf zig-zag encoding
     * @param p The memory offset of `bs`
     * @param bs The bytes array to be decoded
     * @return The decoded signed integer
     * @return The length of `bs` used to get decoded
     */
    function _decode_varints(uint256 p, bytes memory bs) internal pure returns (int256, uint256) {
        /**
         * Refer to https://developers.google.com/protocol-buffers/docs/encoding
         */
        (uint256 u, uint256 sz) = _decode_varint(p, bs);
        int256 s;
        assembly {
            s := xor(shr(1, u), add(not(and(u, 1)), 1))
        }
        return (s, sz);
    }

    /**
     * @dev Decode ProtoBuf fixed-length encoding
     * @param p The memory offset of `bs`
     * @param bs The bytes array to be decoded
     * @return The decoded unsigned integer
     * @return The length of `bs` used to get decoded
     */
    function _decode_uintf(uint256 p, bytes memory bs, uint256 sz) internal pure returns (uint256, uint256) {
        /**
         * Refer to https://developers.google.com/protocol-buffers/docs/encoding
         */
        uint256 x = 0;
        uint256 length = bs.length + WORD_LENGTH;
        assert(p + sz <= length);
        assembly {
            let i := 0
            p := add(bs, p)
            let tmp := mload(p)
            for {} lt(i, sz) {} {
                x := or(x, shl(mul(8, i), byte(i, tmp)))
                p := add(p, 0x01)
                i := add(i, 1)
            }
        }
        return (x, sz);
    }

    /**
     * `_decode_(s)fixed(32|64)` is the concrete implementation of `_decode_uintf`
     */
    function _decode_fixed32(uint256 p, bytes memory bs) internal pure returns (uint32, uint256) {
        (uint256 x, uint256 sz) = _decode_uintf(p, bs, 4);
        return (uint32(x), sz);
    }

    function _decode_fixed64(uint256 p, bytes memory bs) internal pure returns (uint64, uint256) {
        (uint256 x, uint256 sz) = _decode_uintf(p, bs, 8);
        return (uint64(x), sz);
    }

    function _decode_sfixed32(uint256 p, bytes memory bs) internal pure returns (int32, uint256) {
        (uint256 x, uint256 sz) = _decode_uintf(p, bs, 4);
        int256 r;
        assembly {
            r := x
        }
        return (int32(r), sz);
    }

    function _decode_sfixed64(uint256 p, bytes memory bs) internal pure returns (int64, uint256) {
        (uint256 x, uint256 sz) = _decode_uintf(p, bs, 8);
        int256 r;
        assembly {
            r := x
        }
        return (int64(r), sz);
    }

    /**
     * @dev Decode bytes array
     * @param p The memory offset of `bs`
     * @param bs The bytes array to be decoded
     * @return The decoded bytes array
     * @return The length of `bs` used to get decoded
     */
    function _decode_lendelim(uint256 p, bytes memory bs) internal pure returns (bytes memory, uint256) {
        /**
         * First read the size encoded in `varint`, then use the size to read bytes.
         */
        (uint256 len, uint256 sz) = _decode_varint(p, bs);
        bytes memory b = new bytes(len);
        uint256 length = bs.length + WORD_LENGTH;
        assert(p + sz + len <= length);
        uint256 sourcePtr;
        uint256 destPtr;
        assembly {
            destPtr := add(b, 32)
            sourcePtr := add(add(bs, p), sz)
        }
        copyBytes(sourcePtr, destPtr, len);
        return (b, sz + len);
    }

    /**
     * @dev Skip the decoding of a single field
     * @param wt The WireType of the field
     * @param p The memory offset of `bs`
     * @param bs The bytes array to be decoded
     * @return The length of `bs` to skipped
     */
    function _skip_field_decode(WireType wt, uint256 p, bytes memory bs) internal pure returns (uint256) {
        if (wt == ProtoBufRuntime.WireType.Fixed64) {
            return 8;
        } else if (wt == ProtoBufRuntime.WireType.Fixed32) {
            return 4;
        } else if (wt == ProtoBufRuntime.WireType.Varint) {
            (, uint256 size) = ProtoBufRuntime._decode_varint(p, bs);
            return size;
        } else {
            require(wt == ProtoBufRuntime.WireType.LengthDelim);
            (uint256 len, uint256 size) = ProtoBufRuntime._decode_varint(p, bs);
            return size + len;
        }
    }

    // Encoders
    /**
     * @dev Encode ProtoBuf key
     * @param x The field ID
     * @param wt The WireType specified in ProtoBuf
     * @param p The offset of bytes array `bs`
     * @param bs The bytes array to encode
     * @return The length of encoded bytes
     */
    function _encode_key(uint256 x, WireType wt, uint256 p, bytes memory bs) internal pure returns (uint256) {
        uint256 i;
        assembly {
            i := or(mul(x, 8), mod(wt, 8))
        }
        return _encode_varint(i, p, bs);
    }

    /**
     * @dev Encode ProtoBuf varint
     * @param x The unsigned integer to be encoded
     * @param p The offset of bytes array `bs`
     * @param bs The bytes array to encode
     * @return The length of encoded bytes
     */
    function _encode_varint(uint256 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        /**
         * Refer to https://developers.google.com/protocol-buffers/docs/encoding
         */
        uint256 sz = 0;
        assembly {
            let bsptr := add(bs, p)
            let byt := and(x, 0x7f)
            for {} gt(shr(7, x), 0) {} {
                mstore8(bsptr, or(0x80, byt))
                bsptr := add(bsptr, 1)
                sz := add(sz, 1)
                x := shr(7, x)
                byt := and(x, 0x7f)
            }
            mstore8(bsptr, byt)
            sz := add(sz, 1)
        }
        return sz;
    }

    /**
     * @dev Encode ProtoBuf zig-zag encoding
     * @param x The signed integer to be encoded
     * @param p The offset of bytes array `bs`
     * @param bs The bytes array to encode
     * @return The length of encoded bytes
     */
    function _encode_varints(int256 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        /**
         * Refer to https://developers.google.com/protocol-buffers/docs/encoding
         */
        uint256 encodedInt = _encode_zigzag(x);
        return _encode_varint(encodedInt, p, bs);
    }

    /**
     * @dev Encode ProtoBuf bytes
     * @param xs The bytes array to be encoded
     * @param p The offset of bytes array `bs`
     * @param bs The bytes array to encode
     * @return The length of encoded bytes
     */
    function _encode_bytes(bytes memory xs, uint256 p, bytes memory bs) internal pure returns (uint256) {
        uint256 xsLength = xs.length;
        uint256 sz = _encode_varint(xsLength, p, bs);
        uint256 count = 0;
        assembly {
            let bsptr := add(bs, add(p, sz))
            let xsptr := add(xs, 32)
            for {} lt(count, xsLength) {} {
                mstore8(bsptr, byte(0, mload(xsptr)))
                bsptr := add(bsptr, 1)
                xsptr := add(xsptr, 1)
                count := add(count, 1)
            }
        }
        return sz + count;
    }

    /**
     * @dev Encode ProtoBuf string
     * @param xs The string to be encoded
     * @param p The offset of bytes array `bs`
     * @param bs The bytes array to encode
     * @return The length of encoded bytes
     */
    function _encode_string(string memory xs, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_bytes(bytes(xs), p, bs);
    }

    /**
     * `_encode_(u)int(32|64)`, `_encode_enum` and `_encode_bool`
     * are concrete implementation of `_encode_varint`
     */
    function _encode_uint32(uint32 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_varint(x, p, bs);
    }

    function _encode_uint64(uint64 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_varint(x, p, bs);
    }

    function _encode_int32(int32 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        uint64 twosComplement;
        assembly {
            twosComplement := x
        }
        return _encode_varint(twosComplement, p, bs);
    }

    function _encode_int64(int64 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        uint64 twosComplement;
        assembly {
            twosComplement := x
        }
        return _encode_varint(twosComplement, p, bs);
    }

    function _encode_enum(int32 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_int32(x, p, bs);
    }

    function _encode_bool(bool x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        if (x) {
            return _encode_varint(1, p, bs);
        } else {
            return _encode_varint(0, p, bs);
        }
    }

    /**
     * `_encode_sint(32|64)`, `_encode_enum` and `_encode_bool`
     * are the concrete implementation of `_encode_varints`
     */
    function _encode_sint32(int32 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_varints(x, p, bs);
    }

    function _encode_sint64(int64 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_varints(x, p, bs);
    }

    /**
     * `_encode_(s)fixed(32|64)` is the concrete implementation of `_encode_uintf`
     */
    function _encode_fixed32(uint32 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_uintf(x, p, bs, 4);
    }

    function _encode_fixed64(uint64 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_uintf(x, p, bs, 8);
    }

    function _encode_sfixed32(int32 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        uint32 twosComplement;
        assembly {
            twosComplement := x
        }
        return _encode_uintf(twosComplement, p, bs, 4);
    }

    function _encode_sfixed64(int64 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        uint64 twosComplement;
        assembly {
            twosComplement := x
        }
        return _encode_uintf(twosComplement, p, bs, 8);
    }

    /**
     * @dev Encode ProtoBuf fixed-length integer
     * @param x The unsigned integer to be encoded
     * @param p The offset of bytes array `bs`
     * @param bs The bytes array to encode
     * @return The length of encoded bytes
     */
    function _encode_uintf(uint256 x, uint256 p, bytes memory bs, uint256 sz) internal pure returns (uint256) {
        assembly {
            let bsptr := add(sz, add(bs, p))
            let count := sz
            for {} gt(count, 0) {} {
                bsptr := sub(bsptr, 1)
                mstore8(bsptr, byte(sub(32, count), x))
                count := sub(count, 1)
            }
        }
        return sz;
    }

    /**
     * @dev Encode ProtoBuf zig-zag signed integer
     * @param i The unsigned integer to be encoded
     * @return The encoded unsigned integer
     */
    function _encode_zigzag(int256 i) internal pure returns (uint256) {
        if (i >= 0) {
            return uint256(i) * 2;
        } else {
            return uint256(i * -2) - 1;
        }
    }

    // Estimators
    /**
     * @dev Estimate the length of encoded LengthDelim
     * @param i The length of LengthDelim
     * @return The estimated encoded length
     */
    function _sz_lendelim(uint256 i) internal pure returns (uint256) {
        return i + _sz_varint(i);
    }

    /**
     * @dev Estimate the length of encoded ProtoBuf field ID
     * @param i The field ID
     * @return The estimated encoded length
     */
    function _sz_key(uint256 i) internal pure returns (uint256) {
        if (i < 16) {
            return 1;
        } else if (i < 2048) {
            return 2;
        } else if (i < 262_144) {
            return 3;
        } else {
            revert("not supported");
        }
    }

    /**
     * @dev Estimate the length of encoded ProtoBuf varint
     * @param i The unsigned integer
     * @return The estimated encoded length
     */
    function _sz_varint(uint256 i) internal pure returns (uint256) {
        uint256 count = 1;
        assembly {
            i := shr(7, i)
            for {} gt(i, 0) {} {
                i := shr(7, i)
                count := add(count, 1)
            }
        }
        return count;
    }

    /**
     * `_sz_(u)int(32|64)` and `_sz_enum` are the concrete implementation of `_sz_varint`
     */
    function _sz_uint32(uint32 i) internal pure returns (uint256) {
        return _sz_varint(i);
    }

    function _sz_uint64(uint64 i) internal pure returns (uint256) {
        return _sz_varint(i);
    }

    function _sz_int32(int32 i) internal pure returns (uint256) {
        if (i < 0) {
            return 10;
        } else {
            return _sz_varint(uint32(i));
        }
    }

    function _sz_int64(int64 i) internal pure returns (uint256) {
        if (i < 0) {
            return 10;
        } else {
            return _sz_varint(uint64(i));
        }
    }

    function _sz_enum(int64 i) internal pure returns (uint256) {
        if (i < 0) {
            return 10;
        } else {
            return _sz_varint(uint64(i));
        }
    }

    /**
     * `_sz_sint(32|64)` and `_sz_enum` are the concrete implementation of zig-zag encoding
     */
    function _sz_sint32(int32 i) internal pure returns (uint256) {
        return _sz_varint(_encode_zigzag(i));
    }

    function _sz_sint64(int64 i) internal pure returns (uint256) {
        return _sz_varint(_encode_zigzag(i));
    }

    /**
     * `_estimate_packed_repeated_(uint32|uint64|int32|int64|sint32|sint64)`
     */
    function _estimate_packed_repeated_uint32(uint32[] memory a) internal pure returns (uint256) {
        uint256 e = 0;
        for (uint256 i = 0; i < a.length; i++) {
            e += _sz_uint32(a[i]);
        }
        return e;
    }

    function _estimate_packed_repeated_uint64(uint64[] memory a) internal pure returns (uint256) {
        uint256 e = 0;
        for (uint256 i = 0; i < a.length; i++) {
            e += _sz_uint64(a[i]);
        }
        return e;
    }

    function _estimate_packed_repeated_int32(int32[] memory a) internal pure returns (uint256) {
        uint256 e = 0;
        for (uint256 i = 0; i < a.length; i++) {
            e += _sz_int32(a[i]);
        }
        return e;
    }

    function _estimate_packed_repeated_int64(int64[] memory a) internal pure returns (uint256) {
        uint256 e = 0;
        for (uint256 i = 0; i < a.length; i++) {
            e += _sz_int64(a[i]);
        }
        return e;
    }

    function _estimate_packed_repeated_sint32(int32[] memory a) internal pure returns (uint256) {
        uint256 e = 0;
        for (uint256 i = 0; i < a.length; i++) {
            e += _sz_sint32(a[i]);
        }
        return e;
    }

    function _estimate_packed_repeated_sint64(int64[] memory a) internal pure returns (uint256) {
        uint256 e = 0;
        for (uint256 i = 0; i < a.length; i++) {
            e += _sz_sint64(a[i]);
        }
        return e;
    }

    // Element counters for packed repeated fields
    function _count_packed_repeated_varint(uint256 p, uint256 len, bytes memory bs) internal pure returns (uint256) {
        uint256 count = 0;
        uint256 end = p + len;
        while (p < end) {
            uint256 sz;
            (, sz) = _decode_varint(p, bs);
            p += sz;
            count += 1;
        }
        return count;
    }

    // Soltype extensions
    /**
     * @dev Decode Solidity integer and/or fixed-size bytes array, filling from lowest bit.
     * @param n The maximum number of bytes to read
     * @param p The offset of bytes array `bs`
     * @param bs The bytes array to encode
     * @return The bytes32 representation
     * @return The number of bytes used to decode
     */
    function _decode_sol_bytesN_lower(uint8 n, uint256 p, bytes memory bs) internal pure returns (bytes32, uint256) {
        uint256 r;
        (uint256 len, uint256 sz) = _decode_varint(p, bs);
        if (len + sz > n + 3) {
            revert(OVERFLOW_MESSAGE);
        }
        p += 3;
        assert(p < bs.length + WORD_LENGTH);
        assembly {
            r := mload(add(p, bs))
        }
        for (uint256 i = len - 2; i < WORD_LENGTH; i++) {
            r /= 256;
        }
        return (bytes32(r), len + sz);
    }

    /**
     * @dev Decode Solidity integer and/or fixed-size bytes array, filling from highest bit.
     * @param n The maximum number of bytes to read
     * @param p The offset of bytes array `bs`
     * @param bs The bytes array to encode
     * @return The bytes32 representation
     * @return The number of bytes used to decode
     */
    function _decode_sol_bytesN(uint8 n, uint256 p, bytes memory bs) internal pure returns (bytes32, uint256) {
        (uint256 len, uint256 sz) = _decode_varint(p, bs);
        uint256 wordLength = WORD_LENGTH;
        uint256 byteSize = BYTE_SIZE;
        if (len + sz > n + 3) {
            revert(OVERFLOW_MESSAGE);
        }
        p += 3;
        bytes32 acc;
        assert(p < bs.length + WORD_LENGTH);
        assembly {
            acc := mload(add(p, bs))
            let difference := sub(wordLength, sub(len, 2))
            let bits := mul(byteSize, difference)
            acc := shl(bits, shr(bits, acc))
        }
        return (acc, len + sz);
    }

    /*
     * `_decode_sol*` are the concrete implementation of decoding Solidity types
     */
    function _decode_sol_address(uint256 p, bytes memory bs) internal pure returns (address, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytesN(20, p, bs);
        return (address(bytes20(r)), sz);
    }

    function _decode_sol_bool(uint256 p, bytes memory bs) internal pure returns (bool, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(1, p, bs);
        if (r == 0) {
            return (false, sz);
        }
        return (true, sz);
    }

    function _decode_sol_uint(uint256 p, bytes memory bs) internal pure returns (uint256, uint256) {
        return _decode_sol_uint256(p, bs);
    }

    function _decode_sol_uintN(uint8 n, uint256 p, bytes memory bs) internal pure returns (uint256, uint256) {
        (bytes32 u, uint256 sz) = _decode_sol_bytesN_lower(n, p, bs);
        uint256 r;
        assembly {
            r := u
        }
        return (r, sz);
    }

    function _decode_sol_uint8(uint256 p, bytes memory bs) internal pure returns (uint8, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(1, p, bs);
        return (uint8(r), sz);
    }

    function _decode_sol_uint16(uint256 p, bytes memory bs) internal pure returns (uint16, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(2, p, bs);
        return (uint16(r), sz);
    }

    function _decode_sol_uint24(uint256 p, bytes memory bs) internal pure returns (uint24, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(3, p, bs);
        return (uint24(r), sz);
    }

    function _decode_sol_uint32(uint256 p, bytes memory bs) internal pure returns (uint32, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(4, p, bs);
        return (uint32(r), sz);
    }

    function _decode_sol_uint40(uint256 p, bytes memory bs) internal pure returns (uint40, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(5, p, bs);
        return (uint40(r), sz);
    }

    function _decode_sol_uint48(uint256 p, bytes memory bs) internal pure returns (uint48, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(6, p, bs);
        return (uint48(r), sz);
    }

    function _decode_sol_uint56(uint256 p, bytes memory bs) internal pure returns (uint56, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(7, p, bs);
        return (uint56(r), sz);
    }

    function _decode_sol_uint64(uint256 p, bytes memory bs) internal pure returns (uint64, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(8, p, bs);
        return (uint64(r), sz);
    }

    function _decode_sol_uint72(uint256 p, bytes memory bs) internal pure returns (uint72, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(9, p, bs);
        return (uint72(r), sz);
    }

    function _decode_sol_uint80(uint256 p, bytes memory bs) internal pure returns (uint80, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(10, p, bs);
        return (uint80(r), sz);
    }

    function _decode_sol_uint88(uint256 p, bytes memory bs) internal pure returns (uint88, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(11, p, bs);
        return (uint88(r), sz);
    }

    function _decode_sol_uint96(uint256 p, bytes memory bs) internal pure returns (uint96, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(12, p, bs);
        return (uint96(r), sz);
    }

    function _decode_sol_uint104(uint256 p, bytes memory bs) internal pure returns (uint104, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(13, p, bs);
        return (uint104(r), sz);
    }

    function _decode_sol_uint112(uint256 p, bytes memory bs) internal pure returns (uint112, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(14, p, bs);
        return (uint112(r), sz);
    }

    function _decode_sol_uint120(uint256 p, bytes memory bs) internal pure returns (uint120, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(15, p, bs);
        return (uint120(r), sz);
    }

    function _decode_sol_uint128(uint256 p, bytes memory bs) internal pure returns (uint128, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(16, p, bs);
        return (uint128(r), sz);
    }

    function _decode_sol_uint136(uint256 p, bytes memory bs) internal pure returns (uint136, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(17, p, bs);
        return (uint136(r), sz);
    }

    function _decode_sol_uint144(uint256 p, bytes memory bs) internal pure returns (uint144, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(18, p, bs);
        return (uint144(r), sz);
    }

    function _decode_sol_uint152(uint256 p, bytes memory bs) internal pure returns (uint152, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(19, p, bs);
        return (uint152(r), sz);
    }

    function _decode_sol_uint160(uint256 p, bytes memory bs) internal pure returns (uint160, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(20, p, bs);
        return (uint160(r), sz);
    }

    function _decode_sol_uint168(uint256 p, bytes memory bs) internal pure returns (uint168, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(21, p, bs);
        return (uint168(r), sz);
    }

    function _decode_sol_uint176(uint256 p, bytes memory bs) internal pure returns (uint176, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(22, p, bs);
        return (uint176(r), sz);
    }

    function _decode_sol_uint184(uint256 p, bytes memory bs) internal pure returns (uint184, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(23, p, bs);
        return (uint184(r), sz);
    }

    function _decode_sol_uint192(uint256 p, bytes memory bs) internal pure returns (uint192, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(24, p, bs);
        return (uint192(r), sz);
    }

    function _decode_sol_uint200(uint256 p, bytes memory bs) internal pure returns (uint200, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(25, p, bs);
        return (uint200(r), sz);
    }

    function _decode_sol_uint208(uint256 p, bytes memory bs) internal pure returns (uint208, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(26, p, bs);
        return (uint208(r), sz);
    }

    function _decode_sol_uint216(uint256 p, bytes memory bs) internal pure returns (uint216, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(27, p, bs);
        return (uint216(r), sz);
    }

    function _decode_sol_uint224(uint256 p, bytes memory bs) internal pure returns (uint224, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(28, p, bs);
        return (uint224(r), sz);
    }

    function _decode_sol_uint232(uint256 p, bytes memory bs) internal pure returns (uint232, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(29, p, bs);
        return (uint232(r), sz);
    }

    function _decode_sol_uint240(uint256 p, bytes memory bs) internal pure returns (uint240, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(30, p, bs);
        return (uint240(r), sz);
    }

    function _decode_sol_uint248(uint256 p, bytes memory bs) internal pure returns (uint248, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(31, p, bs);
        return (uint248(r), sz);
    }

    function _decode_sol_uint256(uint256 p, bytes memory bs) internal pure returns (uint256, uint256) {
        (uint256 r, uint256 sz) = _decode_sol_uintN(32, p, bs);
        return (uint256(r), sz);
    }

    function _decode_sol_int(uint256 p, bytes memory bs) internal pure returns (int256, uint256) {
        return _decode_sol_int256(p, bs);
    }

    function _decode_sol_intN(uint8 n, uint256 p, bytes memory bs) internal pure returns (int256, uint256) {
        (bytes32 u, uint256 sz) = _decode_sol_bytesN_lower(n, p, bs);
        int256 r;
        assembly {
            r := u
            r := signextend(sub(sz, 4), r)
        }
        return (r, sz);
    }

    function _decode_sol_bytes(uint8 n, uint256 p, bytes memory bs) internal pure returns (bytes32, uint256) {
        (bytes32 u, uint256 sz) = _decode_sol_bytesN(n, p, bs);
        return (u, sz);
    }

    function _decode_sol_int8(uint256 p, bytes memory bs) internal pure returns (int8, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(1, p, bs);
        return (int8(r), sz);
    }

    function _decode_sol_int16(uint256 p, bytes memory bs) internal pure returns (int16, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(2, p, bs);
        return (int16(r), sz);
    }

    function _decode_sol_int24(uint256 p, bytes memory bs) internal pure returns (int24, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(3, p, bs);
        return (int24(r), sz);
    }

    function _decode_sol_int32(uint256 p, bytes memory bs) internal pure returns (int32, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(4, p, bs);
        return (int32(r), sz);
    }

    function _decode_sol_int40(uint256 p, bytes memory bs) internal pure returns (int40, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(5, p, bs);
        return (int40(r), sz);
    }

    function _decode_sol_int48(uint256 p, bytes memory bs) internal pure returns (int48, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(6, p, bs);
        return (int48(r), sz);
    }

    function _decode_sol_int56(uint256 p, bytes memory bs) internal pure returns (int56, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(7, p, bs);
        return (int56(r), sz);
    }

    function _decode_sol_int64(uint256 p, bytes memory bs) internal pure returns (int64, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(8, p, bs);
        return (int64(r), sz);
    }

    function _decode_sol_int72(uint256 p, bytes memory bs) internal pure returns (int72, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(9, p, bs);
        return (int72(r), sz);
    }

    function _decode_sol_int80(uint256 p, bytes memory bs) internal pure returns (int80, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(10, p, bs);
        return (int80(r), sz);
    }

    function _decode_sol_int88(uint256 p, bytes memory bs) internal pure returns (int88, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(11, p, bs);
        return (int88(r), sz);
    }

    function _decode_sol_int96(uint256 p, bytes memory bs) internal pure returns (int96, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(12, p, bs);
        return (int96(r), sz);
    }

    function _decode_sol_int104(uint256 p, bytes memory bs) internal pure returns (int104, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(13, p, bs);
        return (int104(r), sz);
    }

    function _decode_sol_int112(uint256 p, bytes memory bs) internal pure returns (int112, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(14, p, bs);
        return (int112(r), sz);
    }

    function _decode_sol_int120(uint256 p, bytes memory bs) internal pure returns (int120, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(15, p, bs);
        return (int120(r), sz);
    }

    function _decode_sol_int128(uint256 p, bytes memory bs) internal pure returns (int128, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(16, p, bs);
        return (int128(r), sz);
    }

    function _decode_sol_int136(uint256 p, bytes memory bs) internal pure returns (int136, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(17, p, bs);
        return (int136(r), sz);
    }

    function _decode_sol_int144(uint256 p, bytes memory bs) internal pure returns (int144, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(18, p, bs);
        return (int144(r), sz);
    }

    function _decode_sol_int152(uint256 p, bytes memory bs) internal pure returns (int152, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(19, p, bs);
        return (int152(r), sz);
    }

    function _decode_sol_int160(uint256 p, bytes memory bs) internal pure returns (int160, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(20, p, bs);
        return (int160(r), sz);
    }

    function _decode_sol_int168(uint256 p, bytes memory bs) internal pure returns (int168, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(21, p, bs);
        return (int168(r), sz);
    }

    function _decode_sol_int176(uint256 p, bytes memory bs) internal pure returns (int176, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(22, p, bs);
        return (int176(r), sz);
    }

    function _decode_sol_int184(uint256 p, bytes memory bs) internal pure returns (int184, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(23, p, bs);
        return (int184(r), sz);
    }

    function _decode_sol_int192(uint256 p, bytes memory bs) internal pure returns (int192, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(24, p, bs);
        return (int192(r), sz);
    }

    function _decode_sol_int200(uint256 p, bytes memory bs) internal pure returns (int200, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(25, p, bs);
        return (int200(r), sz);
    }

    function _decode_sol_int208(uint256 p, bytes memory bs) internal pure returns (int208, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(26, p, bs);
        return (int208(r), sz);
    }

    function _decode_sol_int216(uint256 p, bytes memory bs) internal pure returns (int216, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(27, p, bs);
        return (int216(r), sz);
    }

    function _decode_sol_int224(uint256 p, bytes memory bs) internal pure returns (int224, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(28, p, bs);
        return (int224(r), sz);
    }

    function _decode_sol_int232(uint256 p, bytes memory bs) internal pure returns (int232, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(29, p, bs);
        return (int232(r), sz);
    }

    function _decode_sol_int240(uint256 p, bytes memory bs) internal pure returns (int240, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(30, p, bs);
        return (int240(r), sz);
    }

    function _decode_sol_int248(uint256 p, bytes memory bs) internal pure returns (int248, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(31, p, bs);
        return (int248(r), sz);
    }

    function _decode_sol_int256(uint256 p, bytes memory bs) internal pure returns (int256, uint256) {
        (int256 r, uint256 sz) = _decode_sol_intN(32, p, bs);
        return (int256(r), sz);
    }

    function _decode_sol_bytes1(uint256 p, bytes memory bs) internal pure returns (bytes1, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(1, p, bs);
        return (bytes1(r), sz);
    }

    function _decode_sol_bytes2(uint256 p, bytes memory bs) internal pure returns (bytes2, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(2, p, bs);
        return (bytes2(r), sz);
    }

    function _decode_sol_bytes3(uint256 p, bytes memory bs) internal pure returns (bytes3, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(3, p, bs);
        return (bytes3(r), sz);
    }

    function _decode_sol_bytes4(uint256 p, bytes memory bs) internal pure returns (bytes4, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(4, p, bs);
        return (bytes4(r), sz);
    }

    function _decode_sol_bytes5(uint256 p, bytes memory bs) internal pure returns (bytes5, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(5, p, bs);
        return (bytes5(r), sz);
    }

    function _decode_sol_bytes6(uint256 p, bytes memory bs) internal pure returns (bytes6, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(6, p, bs);
        return (bytes6(r), sz);
    }

    function _decode_sol_bytes7(uint256 p, bytes memory bs) internal pure returns (bytes7, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(7, p, bs);
        return (bytes7(r), sz);
    }

    function _decode_sol_bytes8(uint256 p, bytes memory bs) internal pure returns (bytes8, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(8, p, bs);
        return (bytes8(r), sz);
    }

    function _decode_sol_bytes9(uint256 p, bytes memory bs) internal pure returns (bytes9, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(9, p, bs);
        return (bytes9(r), sz);
    }

    function _decode_sol_bytes10(uint256 p, bytes memory bs) internal pure returns (bytes10, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(10, p, bs);
        return (bytes10(r), sz);
    }

    function _decode_sol_bytes11(uint256 p, bytes memory bs) internal pure returns (bytes11, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(11, p, bs);
        return (bytes11(r), sz);
    }

    function _decode_sol_bytes12(uint256 p, bytes memory bs) internal pure returns (bytes12, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(12, p, bs);
        return (bytes12(r), sz);
    }

    function _decode_sol_bytes13(uint256 p, bytes memory bs) internal pure returns (bytes13, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(13, p, bs);
        return (bytes13(r), sz);
    }

    function _decode_sol_bytes14(uint256 p, bytes memory bs) internal pure returns (bytes14, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(14, p, bs);
        return (bytes14(r), sz);
    }

    function _decode_sol_bytes15(uint256 p, bytes memory bs) internal pure returns (bytes15, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(15, p, bs);
        return (bytes15(r), sz);
    }

    function _decode_sol_bytes16(uint256 p, bytes memory bs) internal pure returns (bytes16, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(16, p, bs);
        return (bytes16(r), sz);
    }

    function _decode_sol_bytes17(uint256 p, bytes memory bs) internal pure returns (bytes17, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(17, p, bs);
        return (bytes17(r), sz);
    }

    function _decode_sol_bytes18(uint256 p, bytes memory bs) internal pure returns (bytes18, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(18, p, bs);
        return (bytes18(r), sz);
    }

    function _decode_sol_bytes19(uint256 p, bytes memory bs) internal pure returns (bytes19, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(19, p, bs);
        return (bytes19(r), sz);
    }

    function _decode_sol_bytes20(uint256 p, bytes memory bs) internal pure returns (bytes20, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(20, p, bs);
        return (bytes20(r), sz);
    }

    function _decode_sol_bytes21(uint256 p, bytes memory bs) internal pure returns (bytes21, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(21, p, bs);
        return (bytes21(r), sz);
    }

    function _decode_sol_bytes22(uint256 p, bytes memory bs) internal pure returns (bytes22, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(22, p, bs);
        return (bytes22(r), sz);
    }

    function _decode_sol_bytes23(uint256 p, bytes memory bs) internal pure returns (bytes23, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(23, p, bs);
        return (bytes23(r), sz);
    }

    function _decode_sol_bytes24(uint256 p, bytes memory bs) internal pure returns (bytes24, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(24, p, bs);
        return (bytes24(r), sz);
    }

    function _decode_sol_bytes25(uint256 p, bytes memory bs) internal pure returns (bytes25, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(25, p, bs);
        return (bytes25(r), sz);
    }

    function _decode_sol_bytes26(uint256 p, bytes memory bs) internal pure returns (bytes26, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(26, p, bs);
        return (bytes26(r), sz);
    }

    function _decode_sol_bytes27(uint256 p, bytes memory bs) internal pure returns (bytes27, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(27, p, bs);
        return (bytes27(r), sz);
    }

    function _decode_sol_bytes28(uint256 p, bytes memory bs) internal pure returns (bytes28, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(28, p, bs);
        return (bytes28(r), sz);
    }

    function _decode_sol_bytes29(uint256 p, bytes memory bs) internal pure returns (bytes29, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(29, p, bs);
        return (bytes29(r), sz);
    }

    function _decode_sol_bytes30(uint256 p, bytes memory bs) internal pure returns (bytes30, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(30, p, bs);
        return (bytes30(r), sz);
    }

    function _decode_sol_bytes31(uint256 p, bytes memory bs) internal pure returns (bytes31, uint256) {
        (bytes32 r, uint256 sz) = _decode_sol_bytes(31, p, bs);
        return (bytes31(r), sz);
    }

    function _decode_sol_bytes32(uint256 p, bytes memory bs) internal pure returns (bytes32, uint256) {
        return _decode_sol_bytes(32, p, bs);
    }

    /*
     * `_encode_sol*` are the concrete implementation of encoding Solidity types
     */
    function _encode_sol_address(address x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(uint160(x)), 20, p, bs);
    }

    function _encode_sol_uint(uint256 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 32, p, bs);
    }

    function _encode_sol_uint8(uint8 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 1, p, bs);
    }

    function _encode_sol_uint16(uint16 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 2, p, bs);
    }

    function _encode_sol_uint24(uint24 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 3, p, bs);
    }

    function _encode_sol_uint32(uint32 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 4, p, bs);
    }

    function _encode_sol_uint40(uint40 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 5, p, bs);
    }

    function _encode_sol_uint48(uint48 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 6, p, bs);
    }

    function _encode_sol_uint56(uint56 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 7, p, bs);
    }

    function _encode_sol_uint64(uint64 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 8, p, bs);
    }

    function _encode_sol_uint72(uint72 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 9, p, bs);
    }

    function _encode_sol_uint80(uint80 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 10, p, bs);
    }

    function _encode_sol_uint88(uint88 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 11, p, bs);
    }

    function _encode_sol_uint96(uint96 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 12, p, bs);
    }

    function _encode_sol_uint104(uint104 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 13, p, bs);
    }

    function _encode_sol_uint112(uint112 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 14, p, bs);
    }

    function _encode_sol_uint120(uint120 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 15, p, bs);
    }

    function _encode_sol_uint128(uint128 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 16, p, bs);
    }

    function _encode_sol_uint136(uint136 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 17, p, bs);
    }

    function _encode_sol_uint144(uint144 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 18, p, bs);
    }

    function _encode_sol_uint152(uint152 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 19, p, bs);
    }

    function _encode_sol_uint160(uint160 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 20, p, bs);
    }

    function _encode_sol_uint168(uint168 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 21, p, bs);
    }

    function _encode_sol_uint176(uint176 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 22, p, bs);
    }

    function _encode_sol_uint184(uint184 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 23, p, bs);
    }

    function _encode_sol_uint192(uint192 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 24, p, bs);
    }

    function _encode_sol_uint200(uint200 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 25, p, bs);
    }

    function _encode_sol_uint208(uint208 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 26, p, bs);
    }

    function _encode_sol_uint216(uint216 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 27, p, bs);
    }

    function _encode_sol_uint224(uint224 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 28, p, bs);
    }

    function _encode_sol_uint232(uint232 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 29, p, bs);
    }

    function _encode_sol_uint240(uint240 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 30, p, bs);
    }

    function _encode_sol_uint248(uint248 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 31, p, bs);
    }

    function _encode_sol_uint256(uint256 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(uint256(x), 32, p, bs);
    }

    function _encode_sol_int(int256 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(x, 32, p, bs);
    }

    function _encode_sol_int8(int8 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 1, p, bs);
    }

    function _encode_sol_int16(int16 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 2, p, bs);
    }

    function _encode_sol_int24(int24 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 3, p, bs);
    }

    function _encode_sol_int32(int32 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 4, p, bs);
    }

    function _encode_sol_int40(int40 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 5, p, bs);
    }

    function _encode_sol_int48(int48 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 6, p, bs);
    }

    function _encode_sol_int56(int56 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 7, p, bs);
    }

    function _encode_sol_int64(int64 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 8, p, bs);
    }

    function _encode_sol_int72(int72 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 9, p, bs);
    }

    function _encode_sol_int80(int80 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 10, p, bs);
    }

    function _encode_sol_int88(int88 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 11, p, bs);
    }

    function _encode_sol_int96(int96 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 12, p, bs);
    }

    function _encode_sol_int104(int104 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 13, p, bs);
    }

    function _encode_sol_int112(int112 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 14, p, bs);
    }

    function _encode_sol_int120(int120 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 15, p, bs);
    }

    function _encode_sol_int128(int128 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 16, p, bs);
    }

    function _encode_sol_int136(int136 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 17, p, bs);
    }

    function _encode_sol_int144(int144 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 18, p, bs);
    }

    function _encode_sol_int152(int152 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 19, p, bs);
    }

    function _encode_sol_int160(int160 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 20, p, bs);
    }

    function _encode_sol_int168(int168 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 21, p, bs);
    }

    function _encode_sol_int176(int176 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 22, p, bs);
    }

    function _encode_sol_int184(int184 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 23, p, bs);
    }

    function _encode_sol_int192(int192 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 24, p, bs);
    }

    function _encode_sol_int200(int200 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 25, p, bs);
    }

    function _encode_sol_int208(int208 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 26, p, bs);
    }

    function _encode_sol_int216(int216 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 27, p, bs);
    }

    function _encode_sol_int224(int224 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 28, p, bs);
    }

    function _encode_sol_int232(int232 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 29, p, bs);
    }

    function _encode_sol_int240(int240 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 30, p, bs);
    }

    function _encode_sol_int248(int248 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(int256(x), 31, p, bs);
    }

    function _encode_sol_int256(int256 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol(x, 32, p, bs);
    }

    function _encode_sol_bytes1(bytes1 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 1, p, bs);
    }

    function _encode_sol_bytes2(bytes2 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 2, p, bs);
    }

    function _encode_sol_bytes3(bytes3 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 3, p, bs);
    }

    function _encode_sol_bytes4(bytes4 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 4, p, bs);
    }

    function _encode_sol_bytes5(bytes5 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 5, p, bs);
    }

    function _encode_sol_bytes6(bytes6 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 6, p, bs);
    }

    function _encode_sol_bytes7(bytes7 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 7, p, bs);
    }

    function _encode_sol_bytes8(bytes8 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 8, p, bs);
    }

    function _encode_sol_bytes9(bytes9 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 9, p, bs);
    }

    function _encode_sol_bytes10(bytes10 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 10, p, bs);
    }

    function _encode_sol_bytes11(bytes11 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 11, p, bs);
    }

    function _encode_sol_bytes12(bytes12 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 12, p, bs);
    }

    function _encode_sol_bytes13(bytes13 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 13, p, bs);
    }

    function _encode_sol_bytes14(bytes14 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 14, p, bs);
    }

    function _encode_sol_bytes15(bytes15 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 15, p, bs);
    }

    function _encode_sol_bytes16(bytes16 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 16, p, bs);
    }

    function _encode_sol_bytes17(bytes17 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 17, p, bs);
    }

    function _encode_sol_bytes18(bytes18 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 18, p, bs);
    }

    function _encode_sol_bytes19(bytes19 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 19, p, bs);
    }

    function _encode_sol_bytes20(bytes20 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 20, p, bs);
    }

    function _encode_sol_bytes21(bytes21 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 21, p, bs);
    }

    function _encode_sol_bytes22(bytes22 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 22, p, bs);
    }

    function _encode_sol_bytes23(bytes23 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 23, p, bs);
    }

    function _encode_sol_bytes24(bytes24 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 24, p, bs);
    }

    function _encode_sol_bytes25(bytes25 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 25, p, bs);
    }

    function _encode_sol_bytes26(bytes26 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 26, p, bs);
    }

    function _encode_sol_bytes27(bytes27 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 27, p, bs);
    }

    function _encode_sol_bytes28(bytes28 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 28, p, bs);
    }

    function _encode_sol_bytes29(bytes29 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 29, p, bs);
    }

    function _encode_sol_bytes30(bytes30 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 30, p, bs);
    }

    function _encode_sol_bytes31(bytes31 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(bytes32(x), 31, p, bs);
    }

    function _encode_sol_bytes32(bytes32 x, uint256 p, bytes memory bs) internal pure returns (uint256) {
        return _encode_sol_bytes(x, 32, p, bs);
    }

    /**
     * @dev Encode the key of Solidity integer and/or fixed-size bytes array.
     * @param sz The number of bytes used to encode Solidity types
     * @param p The offset of bytes array `bs`
     * @param bs The bytes array to encode
     * @return The number of bytes used to encode
     */
    function _encode_sol_header(uint256 sz, uint256 p, bytes memory bs) internal pure returns (uint256) {
        uint256 offset = p;
        p += _encode_varint(sz + 2, p, bs);
        p += _encode_key(1, WireType.LengthDelim, p, bs);
        p += _encode_varint(sz, p, bs);
        return p - offset;
    }

    /**
     * @dev Encode Solidity type
     * @param x The unsinged integer to be encoded
     * @param sz The number of bytes used to encode Solidity types
     * @param p The offset of bytes array `bs`
     * @param bs The bytes array to encode
     * @return The number of bytes used to encode
     */
    function _encode_sol(uint256 x, uint256 sz, uint256 p, bytes memory bs) internal pure returns (uint256) {
        uint256 offset = p;
        uint256 size;
        p += 3;
        size = _encode_sol_raw_other(x, p, bs, sz);
        p += size;
        _encode_sol_header(size, offset, bs);
        return p - offset;
    }

    /**
     * @dev Encode Solidity type
     * @param x The signed integer to be encoded
     * @param sz The number of bytes used to encode Solidity types
     * @param p The offset of bytes array `bs`
     * @param bs The bytes array to encode
     * @return The number of bytes used to encode
     */
    function _encode_sol(int256 x, uint256 sz, uint256 p, bytes memory bs) internal pure returns (uint256) {
        uint256 offset = p;
        uint256 size;
        p += 3;
        size = _encode_sol_raw_other(x, p, bs, sz);
        p += size;
        _encode_sol_header(size, offset, bs);
        return p - offset;
    }

    /**
     * @dev Encode Solidity type
     * @param x The fixed-size byte array to be encoded
     * @param sz The number of bytes used to encode Solidity types
     * @param p The offset of bytes array `bs`
     * @param bs The bytes array to encode
     * @return The number of bytes used to encode
     */
    function _encode_sol_bytes(bytes32 x, uint256 sz, uint256 p, bytes memory bs) internal pure returns (uint256) {
        uint256 offset = p;
        uint256 size;
        p += 3;
        size = _encode_sol_raw_bytes_array(x, p, bs, sz);
        p += size;
        _encode_sol_header(size, offset, bs);
        return p - offset;
    }

    /**
     * @dev Get the actual size needed to encoding an unsigned integer
     * @param x The unsigned integer to be encoded
     * @param sz The maximum number of bytes used to encode Solidity types
     * @return The number of bytes needed for encoding `x`
     */
    function _get_real_size(uint256 x, uint256 sz) internal pure returns (uint256) {
        uint256 base = 0xff;
        uint256 realSize = sz;
        while (x & (base << (realSize * BYTE_SIZE - BYTE_SIZE)) == 0 && realSize > 0) {
            realSize -= 1;
        }
        if (realSize == 0) {
            realSize = 1;
        }
        return realSize;
    }

    /**
     * @dev Get the actual size needed to encoding an signed integer
     * @param x The signed integer to be encoded
     * @param sz The maximum number of bytes used to encode Solidity types
     * @return The number of bytes needed for encoding `x`
     */
    function _get_real_size(int256 x, uint256 sz) internal pure returns (uint256) {
        int256 base = 0xff;
        if (x >= 0) {
            uint256 tmp = _get_real_size(uint256(x), sz);
            int256 remainder = (x & (base << (tmp * BYTE_SIZE - BYTE_SIZE))) >> (tmp * BYTE_SIZE - BYTE_SIZE);
            if (remainder >= 128) {
                tmp += 1;
            }
            return tmp;
        }

        uint256 realSize = sz;
        while (
            x & (base << (realSize * BYTE_SIZE - BYTE_SIZE)) == (base << (realSize * BYTE_SIZE - BYTE_SIZE))
                && realSize > 0
        ) {
            realSize -= 1;
        }
        {
            int256 remainder = (x & (base << (realSize * BYTE_SIZE - BYTE_SIZE))) >> (realSize * BYTE_SIZE - BYTE_SIZE);
            if (remainder < 128) {
                realSize += 1;
            }
        }
        return realSize;
    }

    /**
     * @dev Encode the fixed-bytes array
     * @param x The fixed-size byte array to be encoded
     * @param sz The maximum number of bytes used to encode Solidity types
     * @param p The offset of bytes array `bs`
     * @param bs The bytes array to encode
     * @return The number of bytes needed for encoding `x`
     */
    function _encode_sol_raw_bytes_array(bytes32 x, uint256 p, bytes memory bs, uint256 sz)
        internal
        pure
        returns (uint256)
    {
        /**
         * The idea is to not encode the leading bytes of zero.
         */
        uint256 actualSize = sz;
        for (uint256 i = 0; i < sz; i++) {
            uint8 current = uint8(x[sz - 1 - i]);
            if (current == 0 && actualSize > 1) {
                actualSize--;
            } else {
                break;
            }
        }
        assembly {
            let bsptr := add(bs, p)
            let count := actualSize
            for {} gt(count, 0) {} {
                mstore8(bsptr, byte(sub(actualSize, count), x))
                bsptr := add(bsptr, 1)
                count := sub(count, 1)
            }
        }
        return actualSize;
    }

    /**
     * @dev Encode the signed integer
     * @param x The signed integer to be encoded
     * @param sz The maximum number of bytes used to encode Solidity types
     * @param p The offset of bytes array `bs`
     * @param bs The bytes array to encode
     * @return The number of bytes needed for encoding `x`
     */
    function _encode_sol_raw_other(int256 x, uint256 p, bytes memory bs, uint256 sz) internal pure returns (uint256) {
        /**
         * The idea is to not encode the leading bytes of zero.or one,
         * depending on whether it is positive.
         */
        uint256 realSize = _get_real_size(x, sz);
        assembly {
            let bsptr := add(bs, p)
            let count := realSize
            for {} gt(count, 0) {} {
                mstore8(bsptr, byte(sub(32, count), x))
                bsptr := add(bsptr, 1)
                count := sub(count, 1)
            }
        }
        return realSize;
    }

    /**
     * @dev Encode the unsigned integer
     * @param x The unsigned integer to be encoded
     * @param sz The maximum number of bytes used to encode Solidity types
     * @param p The offset of bytes array `bs`
     * @param bs The bytes array to encode
     * @return The number of bytes needed for encoding `x`
     */
    function _encode_sol_raw_other(uint256 x, uint256 p, bytes memory bs, uint256 sz) internal pure returns (uint256) {
        uint256 realSize = _get_real_size(x, sz);
        assembly {
            let bsptr := add(bs, p)
            let count := realSize
            for {} gt(count, 0) {} {
                mstore8(bsptr, byte(sub(32, count), x))
                bsptr := add(bsptr, 1)
                count := sub(count, 1)
            }
        }
        return realSize;
    }
}

// lib/openzeppelin-contracts/contracts/access/Ownable.sol

// OpenZeppelin Contracts (last updated v4.9.0) (access/Ownable.sol)

/**
 * @dev Contract module which provides a basic access control mechanism, where
 * there is an account (an owner) that can be granted exclusive access to
 * specific functions.
 *
 * By default, the owner account will be the one that deploys the contract. This
 * can later be changed with {transferOwnership}.
 *
 * This module is used through inheritance. It will make available the modifier
 * `onlyOwner`, which can be applied to your functions to restrict their use to
 * the owner.
 */
abstract contract Ownable is Context {
    address private _owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev Initializes the contract setting the deployer as the initial owner.
     */
    constructor() {
        _transferOwnership(_msgSender());
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        _checkOwner();
        _;
    }

    /**
     * @dev Returns the address of the current owner.
     */
    function owner() public view virtual returns (address) {
        return _owner;
    }

    /**
     * @dev Throws if the sender is not the owner.
     */
    function _checkOwner() internal view virtual {
        require(owner() == _msgSender(), "Ownable: caller is not the owner");
    }

    /**
     * @dev Leaves the contract without owner. It will not be possible to call
     * `onlyOwner` functions. Can only be called by the current owner.
     *
     * NOTE: Renouncing ownership will leave the contract without an owner,
     * thereby disabling any functionality that is only available to the owner.
     */
    function renounceOwnership() public virtual onlyOwner {
        _transferOwnership(address(0));
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Can only be called by the current owner.
     */
    function transferOwnership(address newOwner) public virtual onlyOwner {
        require(newOwner != address(0), "Ownable: new owner is the zero address");
        _transferOwnership(newOwner);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Internal function without access restriction.
     */
    function _transferOwnership(address newOwner) internal virtual {
        address oldOwner = _owner;
        _owner = newOwner;
        emit OwnershipTransferred(oldOwner, newOwner);
    }
}

// lib/proto/GoogleProtobufAny.sol

library GoogleProtobufAny {
    //struct definition
    struct Data {
        string type_url;
        bytes value;
    }

    // Decoder section

    /**
     * @dev The main decoder for memory
     * @param bs The bytes array to be decoded
     * @return The decoded struct
     */
    function decode(bytes memory bs) internal pure returns (Data memory) {
        (Data memory x,) = _decode(32, bs, bs.length);
        return x;
    }

    /**
     * @dev The main decoder for storage
     * @param self The in-storage struct
     * @param bs The bytes array to be decoded
     */
    function decode(Data storage self, bytes memory bs) internal {
        (Data memory x,) = _decode(32, bs, bs.length);
        store(x, self);
    }
    // inner decoder

    /**
     * @dev The decoder for internal usage
     * @param p The offset of bytes array to start decode
     * @param bs The bytes array to be decoded
     * @param sz The number of bytes expected
     * @return The decoded struct
     * @return The number of bytes decoded
     */
    function _decode(uint256 p, bytes memory bs, uint256 sz) internal pure returns (Data memory, uint256) {
        Data memory r;
        uint256[3] memory counters;
        uint256 fieldId;
        ProtoBufRuntime.WireType wireType;
        uint256 bytesRead;
        uint256 offset = p;
        uint256 pointer = p;
        while (pointer < offset + sz) {
            (fieldId, wireType, bytesRead) = ProtoBufRuntime._decode_key(pointer, bs);
            pointer += bytesRead;
            if (fieldId == 1) {
                pointer += _read_type_url(pointer, bs, r, counters);
            } else if (fieldId == 2) {
                pointer += _read_value(pointer, bs, r, counters);
            } else {
                if (wireType == ProtoBufRuntime.WireType.Fixed64) {
                    uint256 size;
                    (, size) = ProtoBufRuntime._decode_fixed64(pointer, bs);
                    pointer += size;
                }
                if (wireType == ProtoBufRuntime.WireType.Fixed32) {
                    uint256 size;
                    (, size) = ProtoBufRuntime._decode_fixed32(pointer, bs);
                    pointer += size;
                }
                if (wireType == ProtoBufRuntime.WireType.Varint) {
                    uint256 size;
                    (, size) = ProtoBufRuntime._decode_varint(pointer, bs);
                    pointer += size;
                }
                if (wireType == ProtoBufRuntime.WireType.LengthDelim) {
                    uint256 size;
                    (, size) = ProtoBufRuntime._decode_lendelim(pointer, bs);
                    pointer += size;
                }
            }
        }
        return (r, sz);
    }

    // field readers

    /**
     * @dev The decoder for reading a field
     * @param p The offset of bytes array to start decode
     * @param bs The bytes array to be decoded
     * @param r The in-memory struct
     * @param counters The counters for repeated fields
     * @return The number of bytes decoded
     */
    function _read_type_url(uint256 p, bytes memory bs, Data memory r, uint256[3] memory counters)
        internal
        pure
        returns (uint256)
    {
        /**
         * if `r` is NULL, then only counting the number of fields.
         */
        (string memory x, uint256 sz) = ProtoBufRuntime._decode_string(p, bs);
        if (isNil(r)) {
            counters[1] += 1;
        } else {
            r.type_url = x;
            if (counters[1] > 0) counters[1] -= 1;
        }
        return sz;
    }

    /**
     * @dev The decoder for reading a field
     * @param p The offset of bytes array to start decode
     * @param bs The bytes array to be decoded
     * @param r The in-memory struct
     * @param counters The counters for repeated fields
     * @return The number of bytes decoded
     */
    function _read_value(uint256 p, bytes memory bs, Data memory r, uint256[3] memory counters)
        internal
        pure
        returns (uint256)
    {
        /**
         * if `r` is NULL, then only counting the number of fields.
         */
        (bytes memory x, uint256 sz) = ProtoBufRuntime._decode_bytes(p, bs);
        if (isNil(r)) {
            counters[2] += 1;
        } else {
            r.value = x;
            if (counters[2] > 0) counters[2] -= 1;
        }
        return sz;
    }

    // Encoder section

    /**
     * @dev The main encoder for memory
     * @param r The struct to be encoded
     * @return The encoded byte array
     */
    function encode(Data memory r) internal pure returns (bytes memory) {
        bytes memory bs = new bytes(_estimate(r));
        uint256 sz = _encode(r, 32, bs);
        assembly {
            mstore(bs, sz)
        }
        return bs;
    }
    // inner encoder

    /**
     * @dev The encoder for internal usage
     * @param r The struct to be encoded
     * @param p The offset of bytes array to start decode
     * @param bs The bytes array to be decoded
     * @return The number of bytes encoded
     */
    function _encode(Data memory r, uint256 p, bytes memory bs) internal pure returns (uint256) {
        uint256 offset = p;
        uint256 pointer = p;

        pointer += ProtoBufRuntime._encode_key(1, ProtoBufRuntime.WireType.LengthDelim, pointer, bs);
        pointer += ProtoBufRuntime._encode_string(r.type_url, pointer, bs);
        pointer += ProtoBufRuntime._encode_key(2, ProtoBufRuntime.WireType.LengthDelim, pointer, bs);
        pointer += ProtoBufRuntime._encode_bytes(r.value, pointer, bs);
        return pointer - offset;
    }
    // nested encoder

    /**
     * @dev The encoder for inner struct
     * @param r The struct to be encoded
     * @param p The offset of bytes array to start decode
     * @param bs The bytes array to be decoded
     * @return The number of bytes encoded
     */
    function _encode_nested(Data memory r, uint256 p, bytes memory bs) internal pure returns (uint256) {
        /**
         * First encoded `r` into a temporary array, and encode the actual size used.
         * Then copy the temporary array into `bs`.
         */
        uint256 offset = p;
        uint256 pointer = p;
        bytes memory tmp = new bytes(_estimate(r));
        uint256 tmpAddr = ProtoBufRuntime.getMemoryAddress(tmp);
        uint256 bsAddr = ProtoBufRuntime.getMemoryAddress(bs);
        uint256 size = _encode(r, 32, tmp);
        pointer += ProtoBufRuntime._encode_varint(size, pointer, bs);
        ProtoBufRuntime.copyBytes(tmpAddr + 32, bsAddr + pointer, size);
        pointer += size;
        delete tmp;
        return pointer - offset;
    }
    // estimator

    /**
     * @dev The estimator for a struct
     * @param r The struct to be encoded
     * @return The number of bytes encoded in estimation
     */
    function _estimate(Data memory r) internal pure returns (uint256) {
        uint256 e;
        e += 1 + ProtoBufRuntime._sz_lendelim(bytes(r.type_url).length);
        e += 1 + ProtoBufRuntime._sz_lendelim(r.value.length);
        return e;
    }

    //store function
    /**
     * @dev Store in-memory struct to storage
     * @param input The in-memory struct
     * @param output The in-storage struct
     */
    function store(Data memory input, Data storage output) internal {
        output.type_url = input.type_url;
        output.value = input.value;
    }

    //utility functions
    /**
     * @dev Return an empty struct
     * @return r The empty struct
     */
    function nil() internal pure returns (Data memory r) {
        assembly {
            r := 0
        }
    }

    /**
     * @dev Test whether a struct is empty
     * @param x The struct to be tested
     * @return r True if it is empty
     */
    function isNil(Data memory x) internal pure returns (bool r) {
        assembly {
            r := iszero(x)
        }
    }
}
//library Any

// lib/openzeppelin-contracts/contracts/utils/Strings.sol

// OpenZeppelin Contracts (last updated v4.9.0) (utils/Strings.sol)

/**
 * @dev String operations.
 */
library Strings {
    bytes16 private constant _SYMBOLS = "0123456789abcdef";
    uint8 private constant _ADDRESS_LENGTH = 20;

    /**
     * @dev Converts a `uint256` to its ASCII `string` decimal representation.
     */
    function toString(uint256 value) internal pure returns (string memory) {
        unchecked {
            uint256 length = Math.log10(value) + 1;
            string memory buffer = new string(length);
            uint256 ptr;
            /// @solidity memory-safe-assembly
            assembly {
                ptr := add(buffer, add(32, length))
            }
            while (true) {
                ptr--;
                /// @solidity memory-safe-assembly
                assembly {
                    mstore8(ptr, byte(mod(value, 10), _SYMBOLS))
                }
                value /= 10;
                if (value == 0) break;
            }
            return buffer;
        }
    }

    /**
     * @dev Converts a `int256` to its ASCII `string` decimal representation.
     */
    function toString(int256 value) internal pure returns (string memory) {
        return string(abi.encodePacked(value < 0 ? "-" : "", toString(SignedMath.abs(value))));
    }

    /**
     * @dev Converts a `uint256` to its ASCII `string` hexadecimal representation.
     */
    function toHexString(uint256 value) internal pure returns (string memory) {
        unchecked {
            return toHexString(value, Math.log256(value) + 1);
        }
    }

    /**
     * @dev Converts a `uint256` to its ASCII `string` hexadecimal representation with fixed length.
     */
    function toHexString(uint256 value, uint256 length) internal pure returns (string memory) {
        bytes memory buffer = new bytes(2 * length + 2);
        buffer[0] = "0";
        buffer[1] = "x";
        for (uint256 i = 2 * length + 1; i > 1; --i) {
            buffer[i] = _SYMBOLS[value & 0xf];
            value >>= 4;
        }
        require(value == 0, "Strings: hex length insufficient");
        return string(buffer);
    }

    /**
     * @dev Converts an `address` with fixed length of 20 bytes to its not checksummed ASCII `string` hexadecimal
     * representation.
     */
    function toHexString(address addr) internal pure returns (string memory) {
        return toHexString(uint256(uint160(addr)), _ADDRESS_LENGTH);
    }

    /**
     * @dev Returns true if the two strings are equal.
     */
    function equal(string memory a, string memory b) internal pure returns (bool) {
        return keccak256(bytes(a)) == keccak256(bytes(b));
    }
}

// lib/proto/channel.sol

library ProtoChannel {
    //struct definition
    struct Data {
        int32 state;
        int32 ordering;
        ProtoCounterparty.Data counterparty;
        string[] connection_hops;
        string version;
    }

    // Decoder section

    /**
     * @dev The main decoder for memory
     * @param bs The bytes array to be decoded
     * @return The decoded struct
     */
    function decode(bytes memory bs) internal pure returns (Data memory) {
        (Data memory x,) = _decode(32, bs, bs.length);
        return x;
    }

    /**
     * @dev The main decoder for storage
     * @param self The in-storage struct
     * @param bs The bytes array to be decoded
     */
    function decode(Data storage self, bytes memory bs) internal {
        (Data memory x,) = _decode(32, bs, bs.length);
        store(x, self);
    }
    // inner decoder

    /**
     * @dev The decoder for internal usage
     * @param p The offset of bytes array to start decode
     * @param bs The bytes array to be decoded
     * @param sz The number of bytes expected
     * @return The decoded struct
     * @return The number of bytes decoded
     */
    function _decode(uint256 p, bytes memory bs, uint256 sz) internal pure returns (Data memory, uint256) {
        Data memory r;
        uint256[6] memory counters;
        uint256 fieldId;
        ProtoBufRuntime.WireType wireType;
        uint256 bytesRead;
        uint256 offset = p;
        uint256 pointer = p;
        while (pointer < offset + sz) {
            (fieldId, wireType, bytesRead) = ProtoBufRuntime._decode_key(pointer, bs);
            pointer += bytesRead;
            if (fieldId == 1) {
                pointer += _read_state(pointer, bs, r);
            } else if (fieldId == 2) {
                pointer += _read_ordering(pointer, bs, r);
            } else if (fieldId == 3) {
                pointer += _read_counterparty(pointer, bs, r);
            } else if (fieldId == 4) {
                pointer += _read_unpacked_repeated_connection_hops(pointer, bs, nil(), counters);
            } else if (fieldId == 5) {
                pointer += _read_version(pointer, bs, r);
            } else {
                pointer += ProtoBufRuntime._skip_field_decode(wireType, pointer, bs);
            }
        }
        pointer = offset;
        if (counters[4] > 0) {
            require(r.connection_hops.length == 0);
            r.connection_hops = new string[](counters[4]);
        }

        while (pointer < offset + sz) {
            (fieldId, wireType, bytesRead) = ProtoBufRuntime._decode_key(pointer, bs);
            pointer += bytesRead;
            if (fieldId == 4) {
                pointer += _read_unpacked_repeated_connection_hops(pointer, bs, r, counters);
            } else {
                pointer += ProtoBufRuntime._skip_field_decode(wireType, pointer, bs);
            }
        }
        return (r, sz);
    }

    // field readers

    /**
     * @dev The decoder for reading a field
     * @param p The offset of bytes array to start decode
     * @param bs The bytes array to be decoded
     * @param r The in-memory struct
     * @return The number of bytes decoded
     */
    function _read_state(uint256 p, bytes memory bs, Data memory r) internal pure returns (uint256) {
        (int32 x, uint256 sz) = ProtoBufRuntime._decode_int32(p, bs);
        r.state = x;
        return sz;
    }

    /**
     * @dev The decoder for reading a field
     * @param p The offset of bytes array to start decode
     * @param bs The bytes array to be decoded
     * @param r The in-memory struct
     * @return The number of bytes decoded
     */
    function _read_ordering(uint256 p, bytes memory bs, Data memory r) internal pure returns (uint256) {
        (int32 x, uint256 sz) = ProtoBufRuntime._decode_int32(p, bs);
        r.ordering = x;
        return sz;
    }

    /**
     * @dev The decoder for reading a field
     * @param p The offset of bytes array to start decode
     * @param bs The bytes array to be decoded
     * @param r The in-memory struct
     * @return The number of bytes decoded
     */
    function _read_counterparty(uint256 p, bytes memory bs, Data memory r) internal pure returns (uint256) {
        (ProtoCounterparty.Data memory x, uint256 sz) = _decode_ProtoCounterparty(p, bs);
        r.counterparty = x;
        return sz;
    }

    /**
     * @dev The decoder for reading a field
     * @param p The offset of bytes array to start decode
     * @param bs The bytes array to be decoded
     * @param r The in-memory struct
     * @param counters The counters for repeated fields
     * @return The number of bytes decoded
     */
    function _read_unpacked_repeated_connection_hops(
        uint256 p,
        bytes memory bs,
        Data memory r,
        uint256[6] memory counters
    ) internal pure returns (uint256) {
        /**
         * if `r` is NULL, then only counting the number of fields.
         */
        (string memory x, uint256 sz) = ProtoBufRuntime._decode_string(p, bs);
        if (isNil(r)) {
            counters[4] += 1;
        } else {
            r.connection_hops[r.connection_hops.length - counters[4]] = x;
            counters[4] -= 1;
        }
        return sz;
    }

    /**
     * @dev The decoder for reading a field
     * @param p The offset of bytes array to start decode
     * @param bs The bytes array to be decoded
     * @param r The in-memory struct
     * @return The number of bytes decoded
     */
    function _read_version(uint256 p, bytes memory bs, Data memory r) internal pure returns (uint256) {
        (string memory x, uint256 sz) = ProtoBufRuntime._decode_string(p, bs);
        r.version = x;
        return sz;
    }

    // struct decoder
    /**
     * @dev The decoder for reading a inner struct field
     * @param p The offset of bytes array to start decode
     * @param bs The bytes array to be decoded
     * @return The decoded inner-struct
     * @return The number of bytes used to decode
     */
    function _decode_ProtoCounterparty(uint256 p, bytes memory bs)
        internal
        pure
        returns (ProtoCounterparty.Data memory, uint256)
    {
        uint256 pointer = p;
        (uint256 sz, uint256 bytesRead) = ProtoBufRuntime._decode_varint(pointer, bs);
        pointer += bytesRead;
        (ProtoCounterparty.Data memory r,) = ProtoCounterparty._decode(pointer, bs, sz);
        return (r, sz + bytesRead);
    }

    // Encoder section

    /**
     * @dev The main encoder for memory
     * @param r The struct to be encoded
     * @return The encoded byte array
     */
    function encode(Data memory r) internal pure returns (bytes memory) {
        bytes memory bs = new bytes(_estimate(r));
        uint256 sz = _encode(r, 32, bs);
        assembly {
            mstore(bs, sz)
        }
        return bs;
    }
    // inner encoder

    /**
     * @dev The encoder for internal usage
     * @param r The struct to be encoded
     * @param p The offset of bytes array to start decode
     * @param bs The bytes array to be decoded
     * @return The number of bytes encoded
     */
    function _encode(Data memory r, uint256 p, bytes memory bs) internal pure returns (uint256) {
        uint256 offset = p;
        uint256 pointer = p;
        uint256 i;
        if (r.state != 0) {
            pointer += ProtoBufRuntime._encode_key(1, ProtoBufRuntime.WireType.Varint, pointer, bs);
            pointer += ProtoBufRuntime._encode_int32(r.state, pointer, bs);
        }
        if (r.ordering != 0) {
            pointer += ProtoBufRuntime._encode_key(2, ProtoBufRuntime.WireType.Varint, pointer, bs);
            pointer += ProtoBufRuntime._encode_int32(r.ordering, pointer, bs);
        }

        pointer += ProtoBufRuntime._encode_key(3, ProtoBufRuntime.WireType.LengthDelim, pointer, bs);
        pointer += ProtoCounterparty._encode_nested(r.counterparty, pointer, bs);

        if (r.connection_hops.length != 0) {
            for (i = 0; i < r.connection_hops.length; i++) {
                pointer += ProtoBufRuntime._encode_key(4, ProtoBufRuntime.WireType.LengthDelim, pointer, bs);
                pointer += ProtoBufRuntime._encode_string(r.connection_hops[i], pointer, bs);
            }
        }
        if (bytes(r.version).length != 0) {
            pointer += ProtoBufRuntime._encode_key(5, ProtoBufRuntime.WireType.LengthDelim, pointer, bs);
            pointer += ProtoBufRuntime._encode_string(r.version, pointer, bs);
        }
        return pointer - offset;
    }
    // nested encoder

    /**
     * @dev The encoder for inner struct
     * @param r The struct to be encoded
     * @param p The offset of bytes array to start decode
     * @param bs The bytes array to be decoded
     * @return The number of bytes encoded
     */
    function _encode_nested(Data memory r, uint256 p, bytes memory bs) internal pure returns (uint256) {
        /**
         * First encoded `r` into a temporary array, and encode the actual size used.
         * Then copy the temporary array into `bs`.
         */
        uint256 offset = p;
        uint256 pointer = p;
        bytes memory tmp = new bytes(_estimate(r));
        uint256 tmpAddr = ProtoBufRuntime.getMemoryAddress(tmp);
        uint256 bsAddr = ProtoBufRuntime.getMemoryAddress(bs);
        uint256 size = _encode(r, 32, tmp);
        pointer += ProtoBufRuntime._encode_varint(size, pointer, bs);
        ProtoBufRuntime.copyBytes(tmpAddr + 32, bsAddr + pointer, size);
        pointer += size;
        delete tmp;
        return pointer - offset;
    }
    // estimator

    /**
     * @dev The estimator for a struct
     * @param r The struct to be encoded
     * @return The number of bytes encoded in estimation
     */
    function _estimate(Data memory r) internal pure returns (uint256) {
        uint256 e;
        uint256 i;
        e += 1 + ProtoBufRuntime._sz_int32(r.state);
        e += 1 + ProtoBufRuntime._sz_int32(r.ordering);
        e += 1 + ProtoBufRuntime._sz_lendelim(ProtoCounterparty._estimate(r.counterparty));
        for (i = 0; i < r.connection_hops.length; i++) {
            e += 1 + ProtoBufRuntime._sz_lendelim(bytes(r.connection_hops[i]).length);
        }
        e += 1 + ProtoBufRuntime._sz_lendelim(bytes(r.version).length);
        return e;
    }
    // empty checker

    function _empty(Data memory r) internal pure returns (bool) {
        if (r.state != 0) {
            return false;
        }

        if (r.ordering != 0) {
            return false;
        }

        if (r.connection_hops.length != 0) {
            return false;
        }

        if (bytes(r.version).length != 0) {
            return false;
        }

        return true;
    }

    //store function
    /**
     * @dev Store in-memory struct to storage
     * @param input The in-memory struct
     * @param output The in-storage struct
     */
    function store(Data memory input, Data storage output) internal {
        output.state = input.state;
        output.ordering = input.ordering;
        ProtoCounterparty.store(input.counterparty, output.counterparty);
        output.connection_hops = input.connection_hops;
        output.version = input.version;
    }

    //array helpers for ConnectionHops
    /**
     * @dev Add value to an array
     * @param self The in-memory struct
     * @param value The value to add
     */
    function addConnectionHops(Data memory self, string memory value) internal pure {
        /**
         * First resize the array. Then add the new element to the end.
         */
        string[] memory tmp = new string[](self.connection_hops.length + 1);
        for (uint256 i = 0; i < self.connection_hops.length; i++) {
            tmp[i] = self.connection_hops[i];
        }
        tmp[self.connection_hops.length] = value;
        self.connection_hops = tmp;
    }

    //utility functions
    /**
     * @dev Return an empty struct
     * @return r The empty struct
     */
    function nil() internal pure returns (Data memory r) {
        assembly {
            r := 0
        }
    }

    /**
     * @dev Test whether a struct is empty
     * @param x The struct to be tested
     * @return r True if it is empty
     */
    function isNil(Data memory x) internal pure returns (bool r) {
        assembly {
            r := iszero(x)
        }
    }
}
//library ProtoChannel

library ProtoCounterparty {
    //struct definition
    struct Data {
        string port_id;
        string channel_id;
    }

    // Decoder section

    /**
     * @dev The main decoder for memory
     * @param bs The bytes array to be decoded
     * @return The decoded struct
     */
    function decode(bytes memory bs) internal pure returns (Data memory) {
        (Data memory x,) = _decode(32, bs, bs.length);
        return x;
    }

    /**
     * @dev The main decoder for storage
     * @param self The in-storage struct
     * @param bs The bytes array to be decoded
     */
    function decode(Data storage self, bytes memory bs) internal {
        (Data memory x,) = _decode(32, bs, bs.length);
        store(x, self);
    }
    // inner decoder

    /**
     * @dev The decoder for internal usage
     * @param p The offset of bytes array to start decode
     * @param bs The bytes array to be decoded
     * @param sz The number of bytes expected
     * @return The decoded struct
     * @return The number of bytes decoded
     */
    function _decode(uint256 p, bytes memory bs, uint256 sz) internal pure returns (Data memory, uint256) {
        Data memory r;
        uint256 fieldId;
        ProtoBufRuntime.WireType wireType;
        uint256 bytesRead;
        uint256 offset = p;
        uint256 pointer = p;
        while (pointer < offset + sz) {
            (fieldId, wireType, bytesRead) = ProtoBufRuntime._decode_key(pointer, bs);
            pointer += bytesRead;
            if (fieldId == 1) {
                pointer += _read_port_id(pointer, bs, r);
            } else if (fieldId == 2) {
                pointer += _read_channel_id(pointer, bs, r);
            } else {
                pointer += ProtoBufRuntime._skip_field_decode(wireType, pointer, bs);
            }
        }
        return (r, sz);
    }

    // field readers

    /**
     * @dev The decoder for reading a field
     * @param p The offset of bytes array to start decode
     * @param bs The bytes array to be decoded
     * @param r The in-memory struct
     * @return The number of bytes decoded
     */
    function _read_port_id(uint256 p, bytes memory bs, Data memory r) internal pure returns (uint256) {
        (string memory x, uint256 sz) = ProtoBufRuntime._decode_string(p, bs);
        r.port_id = x;
        return sz;
    }

    /**
     * @dev The decoder for reading a field
     * @param p The offset of bytes array to start decode
     * @param bs The bytes array to be decoded
     * @param r The in-memory struct
     * @return The number of bytes decoded
     */
    function _read_channel_id(uint256 p, bytes memory bs, Data memory r) internal pure returns (uint256) {
        (string memory x, uint256 sz) = ProtoBufRuntime._decode_string(p, bs);
        r.channel_id = x;
        return sz;
    }

    // Encoder section

    /**
     * @dev The main encoder for memory
     * @param r The struct to be encoded
     * @return The encoded byte array
     */
    function encode(Data memory r) internal pure returns (bytes memory) {
        bytes memory bs = new bytes(_estimate(r));
        uint256 sz = _encode(r, 32, bs);
        assembly {
            mstore(bs, sz)
        }
        return bs;
    }
    // inner encoder

    /**
     * @dev The encoder for internal usage
     * @param r The struct to be encoded
     * @param p The offset of bytes array to start decode
     * @param bs The bytes array to be decoded
     * @return The number of bytes encoded
     */
    function _encode(Data memory r, uint256 p, bytes memory bs) internal pure returns (uint256) {
        uint256 offset = p;
        uint256 pointer = p;

        if (bytes(r.port_id).length != 0) {
            pointer += ProtoBufRuntime._encode_key(1, ProtoBufRuntime.WireType.LengthDelim, pointer, bs);
            pointer += ProtoBufRuntime._encode_string(r.port_id, pointer, bs);
        }
        if (bytes(r.channel_id).length != 0) {
            pointer += ProtoBufRuntime._encode_key(2, ProtoBufRuntime.WireType.LengthDelim, pointer, bs);
            pointer += ProtoBufRuntime._encode_string(r.channel_id, pointer, bs);
        }
        return pointer - offset;
    }
    // nested encoder

    /**
     * @dev The encoder for inner struct
     * @param r The struct to be encoded
     * @param p The offset of bytes array to start decode
     * @param bs The bytes array to be decoded
     * @return The number of bytes encoded
     */
    function _encode_nested(Data memory r, uint256 p, bytes memory bs) internal pure returns (uint256) {
        /**
         * First encoded `r` into a temporary array, and encode the actual size used.
         * Then copy the temporary array into `bs`.
         */
        uint256 offset = p;
        uint256 pointer = p;
        bytes memory tmp = new bytes(_estimate(r));
        uint256 tmpAddr = ProtoBufRuntime.getMemoryAddress(tmp);
        uint256 bsAddr = ProtoBufRuntime.getMemoryAddress(bs);
        uint256 size = _encode(r, 32, tmp);
        pointer += ProtoBufRuntime._encode_varint(size, pointer, bs);
        ProtoBufRuntime.copyBytes(tmpAddr + 32, bsAddr + pointer, size);
        pointer += size;
        delete tmp;
        return pointer - offset;
    }
    // estimator

    /**
     * @dev The estimator for a struct
     * @param r The struct to be encoded
     * @return The number of bytes encoded in estimation
     */
    function _estimate(Data memory r) internal pure returns (uint256) {
        uint256 e;
        e += 1 + ProtoBufRuntime._sz_lendelim(bytes(r.port_id).length);
        e += 1 + ProtoBufRuntime._sz_lendelim(bytes(r.channel_id).length);
        return e;
    }
    // empty checker

    function _empty(Data memory r) internal pure returns (bool) {
        if (bytes(r.port_id).length != 0) {
            return false;
        }

        if (bytes(r.channel_id).length != 0) {
            return false;
        }

        return true;
    }

    //store function
    /**
     * @dev Store in-memory struct to storage
     * @param input The in-memory struct
     * @param output The in-storage struct
     */
    function store(Data memory input, Data storage output) internal {
        output.port_id = input.port_id;
        output.channel_id = input.channel_id;
    }

    //utility functions
    /**
     * @dev Return an empty struct
     * @return r The empty struct
     */
    function nil() internal pure returns (Data memory r) {
        assembly {
            r := 0
        }
    }

    /**
     * @dev Test whether a struct is empty
     * @param x The struct to be tested
     * @return r True if it is empty
     */
    function isNil(Data memory x) internal pure returns (bool r) {
        assembly {
            r := iszero(x)
        }
    }
}
//library ProtoCounterparty

// contracts/libs/Ibc.sol

/**
 * Ibc.sol
 * Basic IBC data structures and utilities.
 */

/// IbcPacket represents the packet data structure received from a remote chain
/// over an IBC channel.
struct IbcPacket {
    /// identifies the channel and port on the sending chain.
    IbcEndpoint src;
    /// identifies the channel and port on the receiving chain.
    IbcEndpoint dest;
    /// The sequence number of the packet on the given channel
    uint64 sequence;
    bytes data;
    /// block height after which the packet times out
    Height timeoutHeight;
    /// block timestamp (in nanoseconds) after which the packet times out
    uint64 timeoutTimestamp;
}

// UniversalPacke represents the data field of an IbcPacket
struct UniversalPacket {
    bytes32 srcPortAddr;
    // source middleware ids bitmap, ie. logic OR of all MW IDs in the MW stack.
    uint256 mwBitmap;
    bytes32 destPortAddr;
    bytes appData;
}

/// Height is a monotonically increasing data type
/// that can be compared against another Height for the purposes of updating and
/// freezing clients
///
/// Normally the RevisionHeight is incremented at each height while keeping
/// RevisionNumber the same. However some consensus algorithms may choose to
/// reset the height in certain conditions e.g. hard forks, state-machine
/// breaking changes In these cases, the RevisionNumber is incremented so that
/// height continues to be monitonically increasing even as the RevisionHeight
/// gets reset
struct Height {
    uint64 revision_number;
    uint64 revision_height;
}

struct AckPacket {
    // success indicates the dApp-level logic. Even when a dApp fails to process a packet per its dApp logic, the
    // delivery of packet and ack packet are still considered successful.
    bool success;
    bytes data;
}

struct IncentivizedAckPacket {
    bool success;
    // Forward relayer's payee address, an EMV address registered on Polymer chain with `RegisterCounterpartyPayee`
    // endpoint.
    // In case of missing payee, zero address is used on Polymer.
    // The relayer payee address is set when incentivized ack is created on Polymer.
    bytes relayer;
    bytes data;
}

enum ChannelOrder {
    NONE,
    UNORDERED,
    ORDERED
}

enum ChannelState {
    // Default State
    UNINITIALIZED,
    // A channel has just started the opening handshake.
    INIT,
    // A channel has acknowledged the handshake step on the counterparty chain.
    TRYOPEN,
    // A channel has completed the handshake. Open channels are
    // ready to send and receive packets.
    OPEN,
    // A channel has been closed and can no longer be used to send or receive
    // packets.
    CLOSED,
    // A channel has been forced closed due to a frozen client in the connection
    // path.
    FROZEN,
    // A channel has acknowledged the handshake step on the counterparty chain, but not yet confirmed with a virtual
    // chain. Virtual channel end ONLY.
    TRY_PENDING,
    // A channel has finished the ChanOpenAck handshake step on chain A, but not yet confirmed with the corresponding
    // virtual chain. Virtual channel end ONLY.
    ACK_PENDING,
    // A channel has finished the ChanOpenConfirm handshake step on chain B, but not yet confirmed with the
    // corresponding
    // virtual chain. Virtual channel end ONLY.
    CONFIRM_PENDING,
    // A channel has finished the ChanCloseConfirm step on chainB, but not yet confirmed with the corresponding
    // virtual chain. Virtual channel end ONLY.
    CLOSE_CONFIRM_PENDING
}

struct Channel {
    string version;
    ChannelOrder ordering;
    bool feeEnabled;
    string[] connectionHops;
    string counterpartyPortId;
    bytes32 counterpartyChannelId;
}

struct ChannelEnd {
    string portId;
    bytes32 channelId;
    string version;
}

struct IbcEndpoint {
    string portId;
    bytes32 channelId;
}

struct Proof {
    // block height at which the proof is valid for a membership or non-membership at the given keyPath
    Height proofHeight;
    // ics23 merkle proof
    bytes proof;
}

// misc errors.
library IBCErrors {
    error invalidCounterParty();
    error invalidCounterPartyPortId();
    error invalidHexStringLength();
    error invalidRelayerAddress();
    error consensusStateVerificationFailed();
    error packetNotTimedOut();
    error invalidAddress();

    // packet sequence related errors.
    error invalidPacketSequence();
    error unexpectedPacketSequence();

    // channel related errors.
    error channelNotOwnedBySender();
    error channelNotOwnedByPortAddress();

    // client related errors.
    error clientAlreadyCreated();
    error clientNotCreated();

    // packet commitment related errors.
    error packetCommitmentNotFound();
    error ackPacketCommitmentAlreadyExists();
    error packetReceiptAlreadyExists();

    // receiver related errors.
    error receiverNotIntendedPacketDestination();
    error receiverNotOriginPacketSender();

    error invalidChannelType(string channelType);
}

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
    function isChannelOpenTry(ChannelEnd calldata counterparty) public pure returns (bool open) {
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

library Ibc {
    /**
     * Convert a non-0x-prefixed hex string to an address
     * @param hexStr hex string to convert to address. Note that the hex string must not include a 0x prefix.
     * hexStr is case-insensitive.
     */
    function _hexStrToAddress(string memory hexStr) external pure returns (address addr) {
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
    }

    // For XXXX => vIBC direction, SC needs to verify the proof of membership of TRY_PENDING
    // For vIBC initiated channel, SC doesn't need to verify any proof, and these should be all empty
    function _isChannelOpenTry(ChannelEnd calldata counterparty) external pure returns (bool open) {
        if (counterparty.channelId == bytes32(0) && bytes(counterparty.version).length == 0) {
            open = false;
            // ChanOpenInit with unknow conterparty
        } else if (counterparty.channelId != bytes32(0) && bytes(counterparty.version).length != 0) {
            // this is the ChanOpenTry; counterparty must not be zero-value
            open = true;
        } else {
            revert IBCErrors.invalidCounterParty();
        }
    }

    function toStr(bytes32 b) public pure returns (string memory outStr) {
        uint8 i = 0;
        while (i < 32 && b[i] != 0) {
            i++;
        }
        bytes memory bytesArray = new bytes(i);
        for (uint8 j = 0; j < i; j++) {
            bytesArray[j] = b[j];
        }
        outStr = string(bytesArray);
    }

    function toStr(uint256 _number) public pure returns (string memory outStr) {
        if (_number == 0) {
            return "0";
        }

        uint256 length;
        uint256 number = _number;

        // Determine the length of the string
        while (number != 0) {
            length++;
            number /= 10;
        }

        bytes memory buffer = new bytes(length);

        // Convert each digit to its ASCII representation
        for (uint256 i = length; i > 0; i--) {
            buffer[i - 1] = bytes1(uint8(48 + (_number % 10)));
            _number /= 10;
        }

        outStr = string(buffer);
    }

    // https://github.com/open-ibc/ibcx-go/blob/ef80dd6784fd/modules/core/24-host/keys.go#L135
    function channelProofKey(string calldata portId, bytes32 channelId) public pure returns (bytes memory proofKey) {
        proofKey = abi.encodePacked("channelEnds/ports/", portId, "/channels/", toStr(channelId));
    }

    // protobuf encoding of a channel object
    // https://github.com/open-ibc/ibcx-go/blob/ef80dd6784fd/modules/core/04-channel/keeper/keeper.go#L92
    function channelProofValue(
        ChannelState state,
        ChannelOrder ordering,
        string calldata version,
        string[] calldata connectionHops,
        ChannelEnd calldata counterparty
    ) public pure returns (bytes memory proofValue) {
        proofValue = ProtoChannel.encode(
            ProtoChannel.Data(
                int32(uint32(state)),
                int32(uint32(ordering)),
                ProtoCounterparty.Data(counterparty.portId, toStr(counterparty.channelId)),
                connectionHops,
                version
            )
        );
    }

    // https://github.com/open-ibc/ibcx-go/blob/ef80dd6784fd/modules/core/24-host/keys.go#L185
    function packetCommitmentProofKey(IbcPacket calldata packet) public pure returns (bytes memory proofKey) {
        proofKey = abi.encodePacked(
            "commitments/ports/",
            packet.src.portId,
            "/channels/",
            toStr(packet.src.channelId),
            "/sequences/",
            toStr(packet.sequence)
        );
    }

    // https://github.com/open-ibc/ibcx-go/blob/ef80dd6784fd/modules/core/04-channel/types/packet.go#L19
    function packetCommitmentProofValue(IbcPacket calldata packet) public pure returns (bytes32 proofValue) {
        proofValue = sha256(
            abi.encodePacked(
                packet.timeoutTimestamp,
                packet.timeoutHeight.revision_number,
                packet.timeoutHeight.revision_height,
                sha256(packet.data)
            )
        );
    }

    // https://github.com/open-ibc/ibcx-go/blob/ef80dd6784fd/modules/core/24-host/keys.go#L201
    function ackProofKey(IbcPacket calldata packet) public pure returns (bytes memory proofKey) {
        proofKey = abi.encodePacked(
            "acks/ports/",
            packet.dest.portId,
            "/channels/",
            toStr(packet.dest.channelId),
            "/sequences/",
            toStr(packet.sequence)
        );
    }

    // https://github.com/open-ibc/ibcx-go/blob/ef80dd6784fd/modules/core/04-channel/types/packet.go#L38
    function ackProofValue(bytes calldata ack) public pure returns (bytes32 proofValue) {
        proofValue = sha256(ack);
    }

    function parseAckData(bytes calldata ack) public pure returns (AckPacket memory ackData) {
        // this hex value is '"result"'
        ackData = (keccak256(ack[1:9]) == keccak256(hex"22726573756c7422"))
            ? AckPacket(true, Base64.decode(string(ack[11:ack.length - 2]))) // result success
            : AckPacket(false, ack[10:ack.length - 2]); // this is an error
    }
}

// contracts/interfaces/IbcDispatcher.sol

/**
 * @title IbcPacketSender
 * @author Polymer Labs
 * @dev IBC packet sender interface.
 */
interface IbcPacketSender {
    function sendPacket(bytes32 channelId, bytes calldata payload, uint64 timeoutTimestamp) external;
}

/**
 * @title IbcDispatcher
 * @author Polymer Labs
 * @notice IBC dispatcher interface is the Polymer Core Smart Contract that implements the core IBC protocol.
 * @dev IBC-compatible contracts depend on this interface to actively participate in the IBC protocol.
 *         Other features are implemented as callback methods in the IbcReceiver interface.
 */
interface IbcDispatcher is IbcPacketSender {
    function channelOpenInit(
        string calldata version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] calldata connectionHops,
        string calldata counterpartyPortId
    ) external;

    function closeIbcChannel(bytes32 channelId) external;

    function portPrefix() external view returns (string memory portPrefix);
}

/**
 * @title IbcEventsEmitter
 * @notice IBC CoreSC events interface.
 */
interface IbcEventsEmitter {
    //
    // channel events
    //
    event ChannelOpenInit(
        address indexed recevier,
        string version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] connectionHops,
        string counterpartyPortId
    );
    event ChannelOpenInitError(address indexed receiver, bytes error);

    event ChannelOpenTry(
        address indexed receiver,
        string version,
        ChannelOrder ordering,
        bool feeEnabled,
        string[] connectionHops,
        string counterpartyPortId,
        bytes32 counterpartyChannelId
    );
    event ChannelOpenTryError(address indexed receiver, bytes error);

    event ChannelOpenAck(address indexed receiver, bytes32 channelId);
    event ChannelOpenAckError(address indexed receiver, bytes error);

    event ChannelOpenConfirm(address indexed receiver, bytes32 channelId);
    event ChannelOpenConfirmError(address indexed receiver, bytes error);

    event CloseIbcChannel(address indexed portAddress, bytes32 indexed channelId);

    event CloseIbcChannelError(address indexed receiver, bytes error);
    event AcknowledgementError(address indexed receiver, bytes error);
    event TimeoutError(address indexed receiver, bytes error);

    //
    // packet events
    //
    event SendPacket(
        address indexed sourcePortAddress,
        bytes32 indexed sourceChannelId,
        bytes packet,
        uint64 sequence,
        // timeoutTimestamp is in UNIX nano seconds; packet will be rejected if
        // delivered after this timestamp on the receiving chain.
        // Timeout semantics is compliant to IBC spec and ibc-go implementation
        uint64 timeoutTimestamp
    );

    event Acknowledgement(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 sequence);

    event Timeout(address indexed sourcePortAddress, bytes32 indexed sourceChannelId, uint64 indexed sequence);

    event RecvPacket(address indexed destPortAddress, bytes32 indexed destChannelId, uint64 sequence);

    event WriteAckPacket(
        address indexed writerPortAddress, bytes32 indexed writerChannelId, uint64 sequence, AckPacket ackPacket
    );

    event WriteTimeoutPacket(
        address indexed writerPortAddress,
        bytes32 indexed writerChannelId,
        uint64 sequence,
        Height timeoutHeight,
        uint64 timeoutTimestamp
    );
}

// contracts/interfaces/IbcReceiver.sol

/**
 * @title IbcChannelReceiver
 * @dev This interface must be implemented by IBC-enabled contracts that act as channel owners and process channel
 * handshake callbacks.
 */
interface IbcChannelReceiver {
    function onChanOpenInit(
        ChannelOrder order,
        string[] calldata connectionHops,
        string calldata counterpartyPortIdentifier,
        string calldata version
    ) external returns (string memory selectedVersion);

    function onChanOpenTry(
        ChannelOrder order,
        string[] memory connectionHops,
        bytes32 channelId,
        string memory counterpartyPortIdentifier,
        bytes32 counterpartychannelId,
        string memory counterpartyVersion
    ) external returns (string memory selectedVersion);

    function onChanOpenAck(bytes32 channelId, bytes32 counterpartychannelId, string calldata counterpartyVersion)
        external;

    function onChanOpenConfirm(bytes32 channelId) external;
    function onCloseIbcChannel(
        bytes32 channelId,
        string calldata counterpartyPortIdentifier,
        bytes32 counterpartyChannelId
    ) external;
}

/**
 * @title IbcPacketReceiver
 * @notice Packet handler interface must be implemented by a IBC-enabled contract.
 * @dev Packet handling callback methods are invoked by the IBC dispatcher.
 */
interface IbcPacketReceiver {
    function onRecvPacket(IbcPacket calldata packet) external returns (AckPacket memory ackPacket);

    function onAcknowledgementPacket(IbcPacket calldata packet, AckPacket calldata ack) external;

    function onTimeoutPacket(IbcPacket calldata packet) external;
}

/**
 * @title IbcReceiver
 * @author Polymer Labs
 * @notice IBC receiver interface must be implemented by a IBC-enabled contract.
 * The implementer, aka. dApp devs, should implement channel handshake and packet handling methods.
 */
interface IbcReceiver is IbcChannelReceiver, IbcPacketReceiver {}

contract IbcReceiverBase is Ownable {
    IbcDispatcher public dispatcher;

    error notIbcDispatcher();
    error UnsupportedVersion();
    error ChannelNotFound();

    /**
     * @dev Modifier to restrict access to only the IBC dispatcher.
     * Only the address with the IBC_ROLE can execute the function.
     * Should add this modifier to all IBC-related callback functions.
     */
    modifier onlyIbcDispatcher() {
        if (msg.sender != address(dispatcher)) {
            revert notIbcDispatcher();
        }
        _;
    }

    /**
     * @dev Constructor function that takes an IbcDispatcher address and grants the IBC_ROLE to the Polymer IBC
     * Dispatcher.
     * @param _dispatcher The address of the IbcDispatcher contract.
     */
    constructor(IbcDispatcher _dispatcher) Ownable() {
        dispatcher = _dispatcher;
    }

    /// This function is called for plain Ether transfers, i.e. for every call with empty calldata.
    // An empty function body is sufficient to receive packet fee refunds.
    receive() external payable {}
}

// contracts/interfaces/IbcMiddleware.sol

/**
 * @title IbcUniversalPacketSender
 * @dev IbcUniversalPacketSender is called by end-users of IBC middleware contracts to send a packet over a MW stack.
 */
interface IbcUniversalPacketSender {
    function sendUniversalPacket(
        bytes32 channelId,
        bytes32 destPortAddr,
        bytes calldata appData,
        uint64 timeoutTimestamp
    ) external;
}

interface IbcMwPacketSender {
    // sendMWPacket must be called by another authorized IBC middleware contract.
    // NEVER by a dApp contract.
    // The middleware contract may choose to send the packet by calling IbcCoreSC or another IBC MW.
    function sendMWPacket(
        bytes32 channelId,
        // original source address of the packet
        bytes32 srcPortAddr,
        bytes32 destPortAddr,
        // source middleware ids bit AND
        uint256 srcMwIds,
        bytes calldata appData,
        uint64 timeoutTimestamp
    ) external;
}

// IBC middleware contracts must implement this interface to relay universal channel packets to other IBC middleware
// contracts.
// If the MW contract also receives ucPacket as the final destination, it must also implement
// IbcUniversalPacketReceiver.
interface IbcMwPacketReceiver {
    function onRecvMWPacket(
        bytes32 channelId,
        UniversalPacket calldata ucPacket,
        // 0-based receiver middleware index in the MW stack.
        //  0 for the first MW directly called by UniversalChannel MW.
        // `mwIndex-1` is the last MW that delivers the packet to the non-MW dApp.
        // Each mw in the stack must increment mwIndex by 1 before calling the next MW.
        uint256 mwIndex,
        address[] calldata mwAddrs
    ) external returns (AckPacket memory ackPacket);

    function onRecvMWAck(
        bytes32 channelId,
        UniversalPacket calldata ucPacket,
        // 0-based receiver middleware index in the MW stack.
        //  0 for the first MW directly called by UniversalChannel MW.
        // `mwIndex-1` is the last MW that delivers the packet to the non-MW dApp.
        // Each mw in the stack must increment mwIndex by 1 before calling the next MW.
        uint256 mwIndex,
        address[] calldata mwAddrs,
        AckPacket calldata ack
    ) external;

    function onRecvMWTimeout(
        bytes32 channelId,
        UniversalPacket calldata ucPacket,
        uint256 mwIndex,
        address[] calldata mwAddrs
    ) external;
}

// dApps and IBC middleware contracts need to implement this interface to receive universal channel packets as packets'
// final destination.
interface IbcUniversalPacketReceiver {
    function onRecvUniversalPacket(bytes32 channelId, UniversalPacket calldata ucPacket)
        external
        returns (AckPacket memory ackPacket);

    function onUniversalAcknowledgement(bytes32 channelId, UniversalPacket memory packet, AckPacket calldata ack)
        external;

    function onTimeoutUniversalPacket(bytes32 channelId, UniversalPacket calldata packet) external;
}

interface IbcMiddlwareProvider is IbcUniversalPacketSender, IbcMwPacketSender {
    /**
     * @dev MW_ID is the ID of MW contract on all supported virtual chains.
     * MW_ID must:
     * - be globally unique, ie. no two MWs should have the same MW_ID registered on Polymer chain.
     * - be identical on all supported virtual chains.
     * - be identical on one virtual chain across multiple deployed MW instances. Each MW instance belong exclusively to
     * one MW stack.
     * - be 1 << N, where N is a non-negative integer [0,255], and not in conflict with other MWs.
     */
    function MW_ID() external view returns (uint256 MWID);
}

/**
 * @title IbcMiddleware
 * @notice IBC middleware interface must be implemented by a IBC-middleware contract, similar to ICS-30
 *  https://github.com/cosmos/ibc/tree/main/spec/app/ics-030-middleware.
 * Current limitations:
 *   - IbcMiddleware must sit on top of a UniversalChannel MW or another IbcMiddleware.
 *   - IbcMiddleware cannnot own an IBC channel. Instead, UniversalChannel MW owns channels.
 */
interface IbcMiddleware is IbcMiddlwareProvider, IbcMwPacketReceiver, IbcUniversalPacketReceiver {}

/**
 * @title IbcUniversalChannelMW
 * @notice This interface must be implemented by IBC universal channel middleware contracts.
 * IbcUniversalChannelMW is a special type of IbcMiddleware that owns IBC channels, which are multiplexed for other
 * IbcMiddleware.
 * IbcUniversalChannelMW cannot sit on top of other MW, and must talk to IbcDispatcher directly.
 */
interface IbcUniversalChannelMW is IbcMiddlwareProvider, IbcPacketReceiver, IbcChannelReceiver {
    error MwBitmpaCannotBeZero();
}

/**
 * @title IbcMwEventsEmitter
 * @notice IBC middleware events interface.
 * @dev IBC middleware contracts should emit events specific to their own middleware,
 * and can choose to emit events not defined in this interface if needed.
 */
interface IbcMwEventsEmitter {
    event SendMWPacket(
        bytes32 indexed channelId,
        bytes32 indexed srcPortAddr,
        bytes32 indexed destPortAddr,
        // middleware UID
        uint256 mwId,
        bytes appData,
        uint64 timeoutTimestamp,
        bytes mwData
    );
    event RecvMWPacket(
        bytes32 indexed channelId,
        bytes32 indexed srcPortAddr,
        bytes32 indexed destPortAddr,
        // middleware UID
        uint256 mwId,
        bytes appData,
        bytes mwData
    );
    event RecvMWAck(
        bytes32 indexed channelId,
        bytes32 indexed srcPortAddr,
        bytes32 indexed destPortAddr,
        // middleware UID
        uint256 mwId,
        bytes appData,
        bytes mwData,
        AckPacket ack
    );
    event RecvMWTimeout(
        bytes32 indexed channelId,
        bytes32 indexed srcPortAddr,
        bytes32 indexed destPortAddr,
        // middleware UID
        uint256 mwId,
        bytes appData,
        bytes mwData
    );
}

/**
 * @title IbcMwUser
 * @dev Contracts that send and recv universal packets via IBC MW should inherit IbcMwUser or implement similiar
 * functions.
 * @notice IbcMwUser ensures only authorized IBC middleware can call IBC callback functions.
 */
contract IbcMwUser is Ownable {
    // default middleware
    address public mw;
    mapping(address => bool) public authorizedMws;

    error UnauthorizedIbcMiddleware();
    error ackDataTooShort();
    error ackAddressMismatch();

    /**
     * @dev Modifier to restrict access to only the IBC middleware.
     * Only the address with the IBC_ROLE can execute the function.
     * Should add this modifier to all IBC-related callback functions.
     */
    modifier onlyIbcMw() {
        if (!authorizedMws[msg.sender]) {
            revert UnauthorizedIbcMiddleware();
        }
        _;
    }

    /**
     * @dev Constructor function that takes an IbcMiddleware address and grants the IBC_ROLE to the Polymer IBC
     * Dispatcher.
     * @param _middleware The address of the IbcMiddleware contract.
     */
    constructor(address _middleware) Ownable() {
        mw = _middleware;
        _authorizeMiddleware(_middleware);
    }

    /// This function is called for plain Ether transfers, i.e. for every call with empty calldata.
    // An empty function body is sufficient to receive packet fee refunds.
    receive() external payable {}

    /**
     * @dev Set the default IBC middleware contract in the MW stack.
     * When sending packets, the default middleware is the next middleware in the MW stack.
     * When receiving packets, the default middleware is the previous middleware in the MW stack.
     * @param _middleware The address of the IbcMiddleware contract.
     * @notice The default middleware is authorized automatically.
     */
    function setDefaultMw(address _middleware) external onlyOwner {
        _authorizeMiddleware(_middleware);
        mw = _middleware;
    }

    /**
     * @dev register an authorized middleware so that modifier onlyIbcMw can be used to restrict access to only
     * authorized middleware.
     * Only the address with the IBC_ROLE can execute the function.
     * @notice Should add this modifier to all IBC-related callback functions.
     */
    function authorizeMiddleware(address middleware) external onlyOwner {
        _authorizeMiddleware(middleware);
    }

    function _authorizeMiddleware(address middleware) internal {
        authorizedMws[address(middleware)] = true;
    }
}

// contracts/examples/Earth.sol

contract Earth is IbcMwUser, IbcUniversalPacketReceiver {
    struct UcPacketWithChannel {
        bytes32 channelId;
        UniversalPacket packet;
    }

    struct UcAckWithChannel {
        bytes32 channelId;
        UniversalPacket packet;
        AckPacket ack;
    }

    // received packet as chain B
    UcPacketWithChannel[] public recvedPackets;
    // received ack packet as chain A
    UcAckWithChannel[] public ackPackets;
    // received timeout packet as chain A
    UcPacketWithChannel[] public timeoutPackets;

    constructor(address _middleware) IbcMwUser(_middleware) {}

    function greet(address destPortAddr, bytes32 channelId, bytes calldata message, uint64 timeoutTimestamp) external {
        IbcUniversalPacketSender(mw).sendUniversalPacket(
            channelId, IbcUtils.toBytes32(destPortAddr), message, timeoutTimestamp
        );
    }

    function onRecvUniversalPacket(bytes32 channelId, UniversalPacket calldata packet)
        external
        onlyIbcMw
        returns (AckPacket memory ackPacket)
    {
        recvedPackets.push(UcPacketWithChannel(channelId, packet));
        return this.generateAckPacket(channelId, IbcUtils.toAddress(packet.srcPortAddr), packet.appData);
    }

    function onUniversalAcknowledgement(bytes32 channelId, UniversalPacket memory packet, AckPacket calldata ack)
        external
        onlyIbcMw
    {
        // verify packet's destPortAddr is the ack's first encoded address. See generateAckPacket())
        if (ack.data.length < 20) revert ackDataTooShort();
        address ackSender = address(bytes20(ack.data[0:20]));
        if (IbcUtils.toAddress(packet.destPortAddr) != ackSender) revert ackAddressMismatch();
        ackPackets.push(UcAckWithChannel(channelId, packet, ack));
    }

    function onTimeoutUniversalPacket(bytes32 channelId, UniversalPacket calldata packet) external onlyIbcMw {
        timeoutPackets.push(UcPacketWithChannel(channelId, packet));
    }

    // For testing only; real dApps should implment their own ack logic
    function generateAckPacket(bytes32, address srcPortAddr, bytes calldata appData)
        external
        view
        returns (AckPacket memory ackPacket)
    {
        return AckPacket(true, abi.encodePacked(this, srcPortAddr, "ack-", appData));
    }
}
