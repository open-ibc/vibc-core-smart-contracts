import { ethers } from "ethers";
import { AccountRegistry, connectProviderAccounts } from "./evm/account";
import { Chain } from "./evm/chain";
import { TxItem, TxRegistry, loadTxRegistry } from "./evm/schemas/tx";
import { Logger } from "./utils/cli";
import {
  StringToStringMap,
  readDeploymentFilesIntoEnv,
  readFromDeploymentFile,
  renderArgs,
} from "./utils/io";
import { DEFAULT_DEPLOYER } from "./utils/constants";
import { ContractRegistry } from "./evm/schemas/contract";
import { updateNoncesForSender } from "./deploy";

export const sendTx = async (
  chain: Chain,
  accountRegistry: AccountRegistry,
  existingContracts: ContractRegistry,
  tx: TxItem,
  logger: Logger,
  dryRun: boolean = false,
  nonces: Record<string, number>,
  env: StringToStringMap
) => {
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

    const deployedContractAddress = renderArgs([tx.address], tx.init, env)[0];

    const ethersContract = new ethers.Contract(
      deployedContractAddress,
      deployedContractAbi,
      deployer
    );
    const args = renderArgs(tx.args, tx.init, env);
    logger.info(
      `calling ${tx.signature} on ${tx.name} @:${deployedContractAddress} with args: \n [${args}]`
    );
    if (!dryRun) {
      const updatedNonces = await updateNoncesForSender(nonces, deployer);
      const overrides = {
        nonce: updatedNonces[deployer.address],
      };
      const sentTx = await ethersContract.getFunction(tx.signature!)(
        ...args,
        overrides
      );
      try {
        await sentTx.wait();
      } catch (err) {
        logger.error(
          `[${chain.chainName}-${chain.deploymentEnvironment}] sendTx ${tx.name} failed: ${err}`
        );
        throw err;
      }
    }
  } catch (err) {
    logger.error(
      `[${chain.chainName}-${chain.deploymentEnvironment}] sendTx ${tx.name} failed: ${err}`
    );
    throw err;
  }
};

/**
 * Send a tx to an existing contract. Reads contract from the _existingContracts args. Can be used for upgrading proxy to new implementation contracts as well
 */
export async function sendTxToChain(
  chain: Chain, // existing contract registry for this chain
  accountRegistry: AccountRegistry, // Set of accounts to send txs from
  existingContracts: ContractRegistry, // Used as a supplementary address source for tests where the contract address is not saved as artifacts
  transactionSpec: TxRegistry, // The txs to send
  logger: Logger,
  dryRun = false
) {
  logger.debug(
    `sending ${transactionSpec.size} transaction(s) to chain ${
      chain.chainName
    }-${
      chain.deploymentEnvironment
    } with contractNames: [${transactionSpec.keys()}]`
  );

  if (!dryRun) {
    accountRegistry = connectProviderAccounts(accountRegistry, chain.rpc);
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
  let env = await readDeploymentFilesIntoEnv({}, chain); // Read from existing deployment files first, then overwrite with explicitly given contract addresses
  env = { ...env, ...existingContractAddresses, chain };

  for (const tx of result.values()) {
    await sendTx(
      chain,
      accountRegistry,
      existingContracts,
      tx,
      logger,
      dryRun,
      {},
      env
    );
  }
}
