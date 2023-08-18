//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.9;

import '@openzeppelin/contracts/utils/Strings.sol';
import '@openzeppelin/contracts/access/Ownable.sol';
import './Ibc.sol';
import './IbcDispatcher.sol';
import './IbcReceiver.sol';
import './IbcVerifier.sol';

contract Escrow is Ownable {
    /// Polymer dispatcher contract address.
    /// Only dispatcher can call `distributeFee` to distribute packet fee to relayers.
    address public dispatcher;

    /// if locked, no fee can be transferred out of Escrow
    bool public locked = false;

    /// This function is called for plain Ether transfers, i.e. for every call with empty calldata.
    receive() external payable {}

    /// setDispatcher sets the dispatcher contract address
    function setDispatcher(address _dispatcher) external onlyOwner {
        dispatcher = _dispatcher;
    }

    /// lockEscrow disables `distributedFee` function
    function lockEscrow() external onlyOwner {
        locked = true;
    }

    /// unlockEscrow enables `distributedFee` function
    function unlockEscrow() external onlyOwner {
        locked = false;
    }

    /// Distribute the packet recv and ack fees to forward and reverse relayers' payees accounts.
    /// It can only be called by the dispatcher contract.
    function distributeAckFees(address payable[2] memory relayers, uint256[2] memory fees) external {
        // preconditions check
        require(!locked, 'Escrow is locked');
        require(msg.sender == dispatcher, 'Only dispatcher can call distributeFee');

        for (uint256 i = 0; i < relayers.length; i++) {
            _transfer(relayers[i], fees[i]);
        }
    }

    /// Distribute packet timeout fee to forward relayer's payee account.
    /// It can only be called by the dispatcher contract.
    function distributeTimeoutFee(address payable relayer, uint256 fee) external {
        // preconditions check
        require(!locked, 'Escrow is locked');
        require(msg.sender == dispatcher, 'Only dispatcher can call distributeFee');

        _transfer(relayer, fee);
    }

    /// refund returns the extra packet fee to packet sender
    function refund(address payable sender, uint256 amount) external {
        // preconditions check
        require(!locked, 'Escrow is locked');
        require(msg.sender == dispatcher, 'Only dispatcher can call refund');

        _transfer(sender, amount);
    }

    /// _transfer transfers the given amount of Ether to the given address.
    function _transfer(address payable payee, uint256 amount) internal {
        (bool sent, ) = payee.call{value: amount}('');
        require(sent, 'Failed to transfer Ether');
    }
}
