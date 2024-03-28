# vIBC Core Smart Contracts

This project includes the core smart contracts for the vIBC protocol, and a few demo contracts that simulate testing and serve as a template for integrating dapp devs.

![](./diagrams/vibcContractsOverview.jpg)


## Repo Structure

All contracts internal to this project are in the `contracts`. This directory contains the following subdirectories:
- `core/`: The contracts that are core to the vIBC protocol. These are the contracts that will be deployed and used in production. 
- `interfaces/`: Interfaces for core and testing contracts 
- `libs/`: Libraries used by the core vIbc protocol.
- The `utils/`, `base/`, and `example/` directories all contain contracts that are not core to the vIBC protocol. These contracts are used only for testing or as templates for dapp contracts that will integrate with the vIBC protocol. 


# Core Contracts
## Dispatcher
The Dispatcher contract routes packets to dapps by mapping channels to dapp addresses. The dispatcher also calls packet and channel callbacks on sending/receiving dapps. This contract uses a UUPS proxy compliant with ERC1967 for upgrades. The proxy is intended to be managed by a single owner address, which will be replaced by a multisig in production. 

Since the optimistic light client contract is used to prove that events happened on the Polymer chain, all methods in the dispatcher are permissionless, aside from the methods to set the connection to client mapping, upgrading the implementation, and setting the dispatcher's port prefix. 

Due to the nature of ibc callbacks, the dispatcher should safely be able to integrate its handler methods with any arbitrary (i.e. potentially malicious) contracts. 

The dispatcher can integrate with multiple light clients of the peptide chain to fit differing needs of security and ux. Currently, the dispatcher integrates with the DummyLightClient (used for testing to reduce the need for generating proofs), and the OptimisticLightClient - an EVM implementation of the Optimistic proof verification currently used by Optimism. Support for future clients is also possible. 

Any dapps which send/receive Ibc packets can do so directly through the dispatcher, or through a middleware contract like the `UniversalChannelHandler`. Dapps that directly integrate with the dispatcher are expected to implement the `IbcReceiver` interface, and those which use the `UniversalChannelHandler` are assumed to implement the `IbcUniversalPacketReceiver` interface.

## OptimisticLightClient
The OptimisticLightClient contract abstracts away proof verification from the dispatcher contract. This light client represents a view of the Polymer chain, and is used to prove that channel handshake and packet sending events happened. 

## OptimisticProofVerifier
The optimisticProofVerifier verifies proofs for the optimistic light client. 

## UniversalChannelHandler
The UniversalChannelHandler is a middleware contract that can be used to save dapps from having to go through the 4-step channel handshake to send or receive Ibc packets. 

## Quick Start with Forge/Foundry

### Install Forge

```sh
curl -L https://foundry.paradigm.xyz | bash
```

This will install Foundryup, then simply follow the instructions on-screen, which will make the `foundryup` command available in your CLI.

Running `foundryup` by itself will install the latest (nightly) precompiled binaries: `forge`, `cast`, `anvil`, and `chisel`. See `foundryup --help` for more options, like installing from a specific version or commit.

Or go to https://book.getfoundry.sh/getting-started/installation for more installation options.

### Build contracts

```sh
forge build
```

### Run Tests

```sh
forge test
```

### Clean environment

```sh
forge clean
```

