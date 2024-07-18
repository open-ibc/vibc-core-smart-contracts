import { ethers } from "ethers";
import {
  DeployedContractObject,
  StringToStringMap,
  readDeploymentFilesIntoEnv,
  renderArgs,
  renderString,
  writeDeployedContractToFile,
} from "./utils/io";
import assert from "assert";
import { AccountRegistry } from "./evm/account";
import {
  ContractRegistry,
  ContractRegistryLoader,
} from "./evm/schemas/contract";
import { Logger } from "./utils/cli";
import { DEFAULT_DEPLOYER } from "./utils/constants";
import { Chain } from "./evm/chain";
import * as vibcContractFactories from "./evm/contracts/index";

/**
 * Return deployment libraries, factory, factory constructor,
 * and rendered arguments for a contract deployment
 */
const getDeployData = (
  factoryName: string,
  deployArgs: any[] | undefined,
  env: StringToStringMap,
  libraries: any[] = [],
  init: { args: any[]; signature: string } | undefined,
  contractFactories: Record<string, any>,
) => {
  // @ts-ignore
  const contractFactoryConstructor = contractFactories[`${factoryName}__factory`];
  assert(
    contractFactoryConstructor,
    `cannot find contract factory constructor for contract: ${factoryName}`
  );

  let libs: any = {};
  libraries.forEach((arg: any) => {
    libs[arg.name] = renderString(arg.address, env);
  });
  libs = [libs];

  const factory = new contractFactoryConstructor(...libs);
  if (!factory) {
    throw new Error(
      `cannot load contract factory for contract: ${factoryName} with factory name: ${factoryName}__factory`
    );
  }

  return {
    args: renderArgs(deployArgs, init, env),
    libraries: libs,
    factory,
    contractFactoryConstructor,
  };
};

export async function deployToChain(
  chain: Chain,
  accountRegistry: AccountRegistry,
  deploySpec: ContractRegistry,
  logger: Logger,
  dryRun = false,
  forceDeployNewContracts = false, // True if you want to use existing deployments when possible
  writeContracts: boolean = true, // True if you want to save persisted artifact files.
  extraContractFactories: Record<string, any> = {},
) {
  logger.info(
    `deploying ${deploySpec.size} contract(s) to chain ${chain.chainName}-${
      chain.deploymentEnvironment
    } with contractNames: [${deploySpec.keys()}]`
  );

  let nonces: Record<string, number> = {}; // maps addresses to nonces
  if (!dryRun) {
    const provider = ethers.getDefaultProvider(chain.rpc);
    const newAccounts = accountRegistry.subset([]);
    for (const [name, wallet] of accountRegistry.entries()) {
      newAccounts.set(name, wallet.connect(provider));
    }
    accountRegistry = newAccounts;
  }

  //  @ts-ignore
  const env: StringToStringMap = { chain };
  if (!forceDeployNewContracts) {
    // Only read from existing contract files if we want to deploy new ones
    await readDeploymentFilesIntoEnv(env, chain);
  }

  // result is the final contract registry after deployment, modified in place
  const result = ContractRegistryLoader.loadSingle(
    JSON.parse(JSON.stringify(deploySpec.serialize()))
  );

  // merge extra contracts into the registry
  let contractFactories = {
    ...vibcContractFactories,
    ...extraContractFactories,
  } 

  const eachContract = async (
    contract: ReturnType<ContractRegistry["mustGet"]>
  ) => {
    try {
      const factoryName = contract.factoryName
        ? contract.factoryName
        : contract.name;

      const constructorData = getDeployData(
        factoryName,
        contract.deployArgs,
        env,
        contract.libraries,
        contract.init,
        contractFactories,
      );

      logger.info(
        `[${chain.chainName}-${chain.deploymentEnvironment}]: deploying ${
          contract.name
        } with args: [${constructorData.args}] with libraries: ${JSON.stringify(
          constructorData.libraries
        )}`
      );
      let deployedAddr = `new.${contract.name}.address`;
      const deployer = accountRegistry.mustGet(
        contract.deployer ? contract.deployer : DEFAULT_DEPLOYER
      );

      // To avoid nonce too low bug, we manually increment nonces for each account
      if (!(deployer.address in nonces)) {
        nonces[deployer.address] = await deployer.getNonce();
      } else {
        nonces[deployer.address] += 1;
      }

      const nonce = nonces[deployer.address];
      if (!dryRun) {
        const overrides = {
          nonce,
        };
        const deployed = await constructorData.factory
          .connect(deployer)
          .deploy(...constructorData.args, overrides);
        await deployed.deploymentTransaction()?.wait(1);
        deployedAddr = await deployed.getAddress();
      }
      // save deployed contract address for its dependencies
      logger.info(
        `deployed contract ${chain.chainName} ${contract.name} at ${deployedAddr}`
      );
      env[contract.name] = deployedAddr;
      // update contract in registry as output result
      contract.address = deployedAddr;
      contract.deployer = deployer.address;
      contract.abi = constructorData.contractFactoryConstructor.abi;
      logger.info(
        `[${chain.chainName}-${chain.deploymentEnvironment}]: deployed ${contract.name} to address: ${deployedAddr}`
      );
      if (writeContracts) {
        const contractObject: DeployedContractObject = {
          factory: factoryName,
          address: deployedAddr,
          abi: constructorData.contractFactoryConstructor.abi,
          bytecode: constructorData.factory.bytecode,
          name: contract.name,
          args: constructorData.args,
          libraries: constructorData.libraries,
        };
        writeDeployedContractToFile(chain, contractObject);
      }
    } catch (err) {
      logger.error(
        `[${chain.chainName}-${chain.deploymentEnvironment}] deploy ${contract.name} failed: ${err}`
      );
      throw err;
    }
  };

  for (const contract of result.values()) {
    await eachContract(contract);
  }

  logger.info(
    `[${chain.chainName}-${chain.deploymentEnvironment}]: finished deploying ${result.size} contracts`
  );

  return {
    chainName: chain.chainName,
    contracts: result,
  };
}
