### Deploying via Command Line 
This npm package exposes two commands - one to deploy new contacts (which automatically creates persisted deployment files), and one to send transactions to contracts from persisted artifact files. The following steps are needed to deploy contracts via the command line: 

1. Ensure that your deployer account and constructor arguments are configured. This can either be done through adding contract spec yaml files located in the specs/ from the root of where this npm module is installed from (requires adding a `specs/evm.accounts.yaml` file and either a `specs/contracts.spec.yaml` or  `specs/upgrade.spec.yaml`), or by setting the KEY_DEPLOYER, RPC_URL, DEPLOYMENT_CHAIN_ID, CHAIN_NAME environment variables. For examples of contract and account spec files, see the `/specs` folder in this repo.
2. Pass in optional command arguments:  
   - RPC_URL - the rpc url to submit deploy txs to, can be a local fork as well 
   - ACCOUNTS_SPECS_PATH - the path to the accounts spec file
   - CHAIN_NAME - the name of the chain to deploy to
   - CHAIN_ID - the chain id of the chain to deploy to
   - DEVELOPMENT_ENVIRONMENT - the environment to deploy to, should be one of staging, production, local, or mainnet
  In addition to these params, any sensitive env variables that can't be saved to yaml files can be passed into the environment as well. For example, if you want to avoid saving a private key to the accounts spec yaml file, you can add that reference the name of the environment in the account spec and pass in the private key as an environment variable.

3. Run either `npx deploy-vibc-core-smart-contracts` to deploy contracts from the contract spec, or `npx upgrade-vibc-core-smart-contracts` to  send an upgrade transaction.

The hierarcy of how configuration is read is: 
- any arguments passed in via clia
- any arguments from the env
- default arugments as defined in src/utils/constants.ts

The deployment files are saved in the format `deployments/<CHAIN_ID>/<DEPLOYMENT_ENVIRONMENT>/<CHAIN_ID>-<DEPLOYMENT_NAME>.json`
### Deploying via imports 
Deployments can also be done through calls through the `deployToChain` and the `sendTxToChain` methods. 

#### Deploying new contracts via imports

```
import {
  AccountRegistry,
  Chain,
  ContractRegistryLoader,
  deployToChain,
  parseObjFromFile,
} from "@open-ibc/vibc-core-smart-contracts";

import { getMainLogger } from "@open-ibc/vibc-core-smart-contracts/utils/cli";
import { DEFAULT_RPC_URL } from "../utils/constants";

// Or can parse it form the env
const accountConfig = {
  name: "local",
  registry: [
    {
      name: "KEY_DEPLOYER",
      privateKey: process.env.KEY_DEPLOYER
    },
  ],
};

const accounts = AccountRegistry.loadMultiple([accountConfig]);
const contracts = ContractRegistryLoader.loadSingle(
  parseObjFromFile("specs/contracts.spec.yaml")
);

const chain: Chain = {
  rpc: process.env.RPC_URL ,
  chainId: process.env.DEPLOYMENT_CHAIN_ID,
  chainName: process.env.CHAIN_NAME,
  vmType: "evm",
  description: "local chain",
};

deployToChain(
  chain,
  accounts.mustGet(chain.chainName),
  contracts.subset(),
  getMainLogger(),
  false
);
```

similar to the command line deploy, this will create a deployment artifact file in the `deployments/` folder.

#### Upgrading existing contracts via imports
Proxy upgrades to existing contracts can be done through the `sendTxToChain` method :

```
#!/usr/bin/env node
import {
  AccountRegistry,
  Chain,
  parseObjFromFile,
} from "@open-ibc/vibc-core-smart-contracts";
import { loadTxRegistry } from "@open-ibc/vibc-core-smart-contracts/evm/schemas/tx";
import { sendTxToChain } from "@open-ibc/vibc-core-smart-contracts";

import { getOutputLogger } from "@open-ibc/vibc-core-smart-contracts/utils/cli";
// Or can parse it form the env
const accountConfig = {
  name: "local",
  registry: [
    {
      name: "KEY_DEPLOYER",
      privateKey: process.env.KEY_DEPLOYER,
    },
  ],
};

const accounts = AccountRegistry.loadMultiple([accountConfig]);
const upgradeTxs = loadTxRegistry(parseObjFromFile("specs/upgrade.spec.yaml"));

const chain: Chain = {
  rpc: process.env.RPC_URL,
  chainId: process.env.CHAIN_ID,
  chainName: "local",
  vmType: "evm",
  description: "local chain",
};

sendTxToChain(
  chain,
  accounts.mustGet(chain.chainName),
  upgradeTxs.subset(),
  getOutputLogger(),
  false
);
```

