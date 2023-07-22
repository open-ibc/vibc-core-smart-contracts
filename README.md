# vIBC Core Smart Contracts

This project includes

- vIBC Core Smart Contracts (CoreSC)
- a few demo contracts that simulate dev users protocol contracts, eg. Mars, etc.

## Quick Start with Forge/Foundry

### Install Forge

```sh
curl -L https://foundry.paradigm.xyz | bash
```

This will install Foundryup, then simply follow the instructions on-screen, which will make the `foundryup` command available in your CLI.

Running `foundryup` by itself will install the latest (nightly) precompiled binaries: `forge`, `cast`, `anvil`, and `chisel`. See `foundryup --help` for more options, like installing from a specific version or commit.

Or go to https://book.getfoundry.sh/getting-started/installation for more installation options.

### Install NPM Dependencies

```sh
npm install
```

NOTE: Since we're migrating from Hardhat to Forge, we need to install both Forge and Hardhat dependencies.
If we get rid of Hardhat completely, this step will be removed. No NPM project is needed anymore by then.

### Run Tests

```sh
forge test
```

## Quick Start with Hardhat

```shell
npm install

# Run CoreSC tests
npx hardhat test
```

Try running some of the following tasks:

```shell
npx hardhat help
npx hardhat test
GAS_REPORT=true npx hardhat test
npx hardhat node
npx hardhat run scripts/deploy.ts
```
