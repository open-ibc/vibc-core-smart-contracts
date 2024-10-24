import { sendTx } from "./tx";
import { deployContract } from "./deploy";
import { Chain } from "./evm/chain";
import { SingleSigAccountRegistry } from "./evm/schemas/account";
import { ContractItemSchema, ContractRegistry } from "./evm/schemas/contract";
import {
  UpdateContractRegistry,
  loadContractUpdateRegistry,
} from "./evm/schemas/contractUpdate";
import { Logger } from "winston";
import { readAccountsIntoEnv, readDeploymentFilesIntoEnv } from "./utils/io";
import { TxItemSchema } from "./evm/schemas/tx";
import { SendingAccountRegistry } from "./evm/schemas/sendingAccount";

type UpdateContractsOpts = {
  dryRun?: boolean;
  forceDeployNewContracts?: boolean;
  writeContracts?: boolean;
  extraContractFactories?: Record<string, object>;
}

// Combination of sendTxToChain and deployContracts. Can do both from a single deploy file, and uses zod to parse the schema.
export async function updateContractsForChain(
  chain: Chain,
  accountRegistry: SingleSigAccountRegistry | SendingAccountRegistry,
  existingContracts: ContractRegistry,
  updateSpec: UpdateContractRegistry,
  logger: Logger,
  options: UpdateContractsOpts = {
    dryRun: false,
    forceDeployNewContracts: false, // True if you want to use existing deployments when possible
    writeContracts: true, // True if you want to save persisted artifact files.
    extraContractFactories: {},
  }
) {
  const {
    dryRun,
    forceDeployNewContracts,
    writeContracts,
    extraContractFactories,
  } = options;

  logger.info(
    `updating ${updateSpec.size} contract(s) for chain ${chain.chainName}-${
      chain.deploymentEnvironment
    } with contractNames: [${updateSpec.keys()}]`
  );

  const nonces: Record<string, number> = {}; // Maps addresses to nonces. Used to avoid nonce too low errors.

  if (!dryRun) {
    accountRegistry.connectProviderAccounts(chain.rpc);
  }

  const existingContractAddresses: Record<string, string> = {};
  existingContracts.toList().forEach((item) => {
    existingContractAddresses[item.name] = item.address ? item.address : "0x";
  });

  let env = await readDeploymentFilesIntoEnv({}, chain); // Read from existing deployment files first, then overwrite with explicitly given contract addresses
  env = await readAccountsIntoEnv(env, accountRegistry); // Read from rendered accounts, useful for accessing things like multisig address from a signer, etc.
  env = { ...process.env, chain, ...existingContractAddresses, ...env };
  if (!forceDeployNewContracts) {
    // Only read from existing contract files if we want to deploy new ones
    await readDeploymentFilesIntoEnv(env, chain);
  }

  // result is the final contract registry after deployment, modified in place
  const cleanedRegistry = loadContractUpdateRegistry(
    JSON.parse(JSON.stringify(updateSpec.serialize()))
  );

  const result = new UpdateContractRegistry([]);

  for (const updateContract of cleanedRegistry.values()) {
    const parsedContractSchema = ContractItemSchema.safeParse(updateContract);
    if (parsedContractSchema.success) {
      const deployed = await deployContract(
        chain,
        accountRegistry,
        parsedContractSchema.data,
        logger,
        dryRun,
        writeContracts,
        extraContractFactories,
        nonces,
        env
      );
      result.set(deployed.name, deployed);
      continue;
    }

    // If not a valid contract schema, then it should be a valid tx schema
    const parsedTxItem = TxItemSchema.safeParse(updateContract);
    if (parsedTxItem.success) {
      await sendTx(
        chain,
        accountRegistry,
        existingContracts,
        parsedTxItem.data,
        logger,
        dryRun,
        nonces,
        env,
        extraContractFactories
      );
      continue;
    }
    logger.error("Invalid schema! Aborting contract setup.");
    throw new Error("Invalid schema! Aborting contract setup.");
  }

  logger.info(
    `[${chain.chainName}-${chain.deploymentEnvironment}]: finished updating ${updateSpec.size} contracts`
  );

  return {
    chainName: chain.chainName,
    updatedContracts: result,
  };
}
