import { ethers } from "ethers";
import { AccountRegistry } from "./evm/account";
import { Chain } from "./evm/chain";
import { TxRegistry, loadTxRegistry } from "./evm/schemas/tx";
import { Logger } from "./utils/cli";
import {
  readDeploymentFilesIntoEnv,
  readFromDeploymentFile,
  renderArgs,
} from "./utils/io";
import { DEFAULT_DEPLOYER } from "./utils/constants";
import { ContractRegistry } from "./evm/schemas/contract";

/**
 * Send a tx to an existing contract. Reads contract from the _existingContracts args. Can be used for upgrading proxy to new implementation contracts as well
 */
export async function sendTxToChain(
  chain: Chain, // existing contract registry for this chain
  accountRegistry: AccountRegistry,
  existingContracts: ContractRegistry,
  transactionSpec: TxRegistry,
  logger: Logger,
  dryRun = false
) {
  logger.debug(
    `sending ${transactionSpec.size} transaction(s) to chain ${
      chain.chainName
    } with contractNames: [${transactionSpec.keys()}]`
  );

  if (!dryRun) {
    const provider = ethers.getDefaultProvider(chain.rpc);
    const newAccounts = accountRegistry.subset([]);
    for (const [name, wallet] of accountRegistry.entries()) {
      newAccounts.set(name, wallet.connect(provider));
    }
    accountRegistry = newAccounts;
  }

  // result is the final contract registry after deployment, modified in place
  const result = loadTxRegistry(
    JSON.parse(JSON.stringify(transactionSpec.serialize()))
  );
  const existingContractAddresses: Record<string, string> = {};
  existingContracts.toList().forEach((item) => {
    existingContractAddresses[item.name] = item.address ? item.address : "0x";
  });

  // @ts-ignore
  let env = { ...existingContractAddresses, chain };
  env = await readDeploymentFilesIntoEnv(env);

  const eachTx = async (tx: ReturnType<TxRegistry["mustGet"]>) => {
    try {
      const factoryName = tx.factoryName ? tx.factoryName : tx.name;
      let deployedContractAbi: any;

      const existingContract = existingContracts.get(factoryName);
      // Fetch the ABI from the existing contract if it exists; otherwise read from deployment files
      if (existingContract && existingContract.abi) {
        deployedContractAbi = existingContract.abi;
      } else {
        const deployedContract: any = await readFromDeploymentFile(
          factoryName,
          chain
        );
        if (!deployedContract) {
          throw new Error(`Contract deployment for ${factoryName} not found!`);
        }
        deployedContractAbi = deployedContract.abi;
      }

      const deployer = accountRegistry.mustGet(
        tx.deployer ? tx.deployer : DEFAULT_DEPLOYER
      );
      const deployedContractAddress = renderArgs([tx.address], "", env)[0];

      const ethersContract = new ethers.Contract(
        deployedContractAddress,
        deployedContractAbi,
        deployer
      );
      const args = renderArgs(tx.args, "", env);
      logger.info(
        `calling ${tx.signature} on ${tx.name} @:${deployedContractAddress} with args: \n [${args}]`
      );
      if (!dryRun) {
        const sentTx = await ethersContract.getFunction(tx.signature!)(...args);
        try {
          await sentTx.wait();
        } catch (err) {
          logger.error(`[${chain.chainName}] sendTx ${tx.name} failed: ${err}`);
          throw err;
        }
      }
    } catch (err) {
      logger.error(`[${chain.chainName}] sendTx ${tx.name} failed: ${err}`);
      throw err;
    }
  };

  for (const tx of result.values()) {
    await eachTx(tx);
  }
}
