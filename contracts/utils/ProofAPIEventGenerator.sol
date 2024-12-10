// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

contract ProofAPIEventGenerator {
    address public owner;

    // Event to log the transaction
    event ProofAPIEventHealthCheck(address indexed to, uint256 amount, uint256 srcChainID, uint256 destChainID);

    // Constructor to set the contract owner
    constructor() {
        owner = msg.sender;
    }

    // Function to receive Ether
    receive(uint256 amount uint256 srcChainID, uint256 destChainID) external payable {
        
       emit ProofAPIEventHealthCheck(msg.sender, msg.value, srcChainID, destChainID) 
    }

    // Function to send Ether from the contract to a specified address
    function sendEther(address payable recipient, uint256 amount) public {
        // Ensure only the owner can send Ether
        require(msg.sender == owner, "Only the owner can call this function");

        // Check that the contract has enough balance
        require(address(this).balance >= amount, "Insufficient contract balance");

        // Transfer the Ether
        recipient.transfer(amount);

        // Emit an event for logging
        emit EtherTransferred(recipient, amount);
    }

    // Function to get the contract's balance
    function getBalance() public view returns (uint256) {
        return address(this).balance;
    }
}