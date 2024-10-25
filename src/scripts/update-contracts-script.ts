#!/usr/bin/env node
import { ContractRegistryLoader, parseObjFromFile } from "..";
import { loadContractUpdateRegistry } from "../evm/schemas/contractUpdate";
import { updateContractsForChain } from "../updateContract";

import { getOutputLogger } from "../utils/cli";
import { UPDATE_SPECS_PATH } from "../utils/constants";
import { parseArgsFromCLI } from "../utils/io";

async function main() {
  const { chain, accounts, args, extraBindingsPath, externalContractsPath } =
    await parseArgsFromCLI();
  const updateSpecs = (args.UPDATE_SPECS_PATH as string) || UPDATE_SPECS_PATH;

  const contractUpdates = loadContractUpdateRegistry(
    parseObjFromFile(updateSpecs)
  );

  const extraContractFactories = extraBindingsPath
    ? await import(extraBindingsPath)
    : null;

  const existingContracts = externalContractsPath
    ? ContractRegistryLoader.loadSingle(externalContractsPath)
    : ContractRegistryLoader.emptySingle();

  updateContractsForChain(
    chain,
    accounts.mustGet(chain.chainName),
    existingContracts,
    contractUpdates,
    getOutputLogger(),
    false,
    false,
    true,
    extraContractFactories
  );
}
main();
