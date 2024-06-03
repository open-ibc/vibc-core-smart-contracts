import { ethers } from "ethers";
import {
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
import * as contracts from "./evm/contracts/index";

/**
 * Return deployment libraries, factory, factory constructor,
 * and rendered arguments for a contract deployment
 */
const getDeployData = (
  factoryName: string,
  deployArgs: any[] | undefined,
  env: StringToStringMap,
  libraries: any[] = [],
  init: { args: any[]; signature: string } | undefined
) => {
  // @ts-ignore
  const contractFactoryConstructor = contracts[`${factoryName}__factory`];
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

  // var encodedInitData = "";
  let initData = "";

  if (init) {
    const initArgs = init.args.map((arg: any) => {
      return typeof arg === "string" ? renderString(arg, env) : arg;
    });
    const iFace = new ethers.Interface([`function ${init.signature}`]);
    initData = iFace.encodeFunctionData(init.signature, initArgs);
  }
  return {
    args: renderArgs(deployArgs, initData, env),
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
  writeContracts: boolean = true // True if you want to save persisted artifact files.
) {
  logger.info(
    `deploying ${deploySpec.size} contract(s) to chain ${
      chain.chainName
    } with contractNames: [${deploySpec.keys()}]`
  );

  logger.info(`deploying with contract spec ${deploySpec}`);

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
    await readDeploymentFilesIntoEnv(env);
  }

  // result is the final contract registry after deployment, modified in place
  const result = ContractRegistryLoader.loadSingle(
    JSON.parse(JSON.stringify(deploySpec.serialize()))
  );

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
        contract.init
      );

      logger.info(
        `[${chain.chainName}]: deploying ${contract.name} with args: [${
          constructorData.args
        }] with libraries: ${JSON.stringify(constructorData.libraries)}`
      );
      let deployedAddr = `new.${contract.name}.address`;
      const deployer = accountRegistry.mustGet(
        contract.deployer ? contract.deployer : DEFAULT_DEPLOYER
      );

      if (!dryRun) {
        const deployed = await constructorData.factory
          .connect(deployer)
          .deploy(...constructorData.args);
        await deployed.deploymentTransaction()?.wait();
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
        `[${chain.chainName}]: deployed ${contract.name} to address: ${deployedAddr}`
      );
      if (writeContracts) {
        const contractObject = {
          factory: factoryName,
          address: deployedAddr,
          abi: constructorData.contractFactoryConstructor.abi,
          bytecode: constructorData.factory.bytecode,
          name: contract.name,
          args: constructorData.args,
        };
        writeDeployedContractToFile(chain, contractObject);
      }
    } catch (err) {
      logger.error(
        `[${chain.chainName}] deploy ${contract.name} failed: ${err}`
      );
      throw err;
    }
  };
  for (const contract of result.values()) {
    await eachContract(contract);
  }

  logger.info(
    `[${chain.chainName}]: finished deploying ${result.size} contracts`
  );

  return {
    chainName: chain.chainName,
    contracts: result,
  };
}
