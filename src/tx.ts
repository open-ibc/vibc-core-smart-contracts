import { ethers } from "ethers";
import {
  AccountRegistry,
  connectProviderAccounts,
  isParsedMultiSigWallet,
} from "./evm/schemas/account";
import { Chain } from "./evm/chain";
import { TxItem, TxRegistry, loadTxRegistry } from "./evm/schemas/tx";
import { Logger } from "./utils/cli";
import {
  StringToStringMap,
  readDeploymentFilesIntoEnv,
  readFactoryAbi,
  readFromDeploymentFile,
  renderArgs,
} from "./utils/io";
import { DEFAULT_DEPLOYER } from "./utils/constants";
import { ContractRegistry } from "./evm/schemas/contract";
import { updateNoncesForSender } from "./deploy";
import { fetchNonceFromSafeAddress, proposeTransaction } from "./multisig/safe";
import * as vibcContractFactories from "./evm/contracts/index";

export async function readAbiFromDeployedContract(
  existingContractOverrides: ContractRegistry,
  factoryName: string,
  factories: Record<string, any> = {}
) {
  const existingContractOverride = existingContractOverrides.get(factoryName);

  // Read from overrides first
  if (existingContractOverride && existingContractOverride.abi) {
    return existingContractOverride.abi;
  }

  return readFactoryAbi(factoryName, factories)
}

export async function sendTx(
  chain: Chain,
  accountRegistry: AccountRegistry,
  existingContractOverrides: ContractRegistry,
  tx: TxItem,
  logger: Logger,
  dryRun: boolean = false,
  nonces: Record<string, number>,
  env: StringToStringMap,
  extraContractFactories: Record<string, any> = {},
) {
  try {
    const factoryName = tx.factoryName ? tx.factoryName : tx.name;
    const contractFactories = {
      ...vibcContractFactories,
      ...extraContractFactories,
    };
  

    const deployedContractAbi = await readAbiFromDeployedContract(existingContractOverrides, factoryName, contractFactories)
    if (!deployedContractAbi) {
      throw new Error(`Could not find ABI for contract ${factoryName}`);
    }

    const account = accountRegistry.mustGet(
      tx.deployer ? tx.deployer : DEFAULT_DEPLOYER
    );

    const deployedContractAddress = renderArgs([tx.address], tx.init, env)[0];

    const args = renderArgs(tx.args, tx.init, env);
    if (isParsedMultiSigWallet(account)) {
      const updatedNonces = await updateNoncesForSender(
        nonces,
        account.safeAddress,
        () => fetchNonceFromSafeAddress(chain.rpc, account.safeAddress)
      );
      // If multisig, we don't directly send the tx, but instead propose it to the safe tx service
      const deployer = account.wallet;
      const contractInterface = new ethers.Interface(deployedContractAbi);
      const callData = contractInterface.encodeFunctionData(
        tx.signature!,
        args
      );

      proposeTransaction(
        account.safeAddress,
        deployedContractAddress,
        callData,
        deployer.privateKey,
        BigInt(chain.chainId),
        chain.rpc,
        updatedNonces[account.safeAddress]
      );
    } else {
      // Send from single acount
      const deployer = account;
      const ethersContract = new ethers.Contract(
        deployedContractAddress,
        deployedContractAbi,
        deployer
      );

      logger.info(
        `calling ${tx.signature} on ${tx.name} @:${deployedContractAddress} with args: \n [${args}]`
      );
      if (!dryRun) {
        const updatedNonces = await updateNoncesForSender(
          nonces,
          deployer.address,
          () => deployer.getNonce()
        );
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
    }
  } catch (err) {
    logger.error(
      `[${chain.chainName}-${chain.deploymentEnvironment}] sendTx ${tx.name} failed: ${err}`
    );
    throw err;
  }
}
