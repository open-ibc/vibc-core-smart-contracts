#!/usr/bin/env node
import { ContractRegistryLoader, parseObjFromFile } from "..";
import { loadContractUpdateRegistry } from "../evm/schemas/contractUpdate";
import { updateContractsForChain } from "../updateContract";

import { getOutputLogger } from "../utils/cli";
import { UPDATE_SPECS_PATH } from "../utils/constants";
import { parseArgsFromCLI } from "../utils/io";

async function main() {
  const { chain, accounts, args, extraBindingsPath, extraArtifactsPath } =
    await parseArgsFromCLI();
  const updateSpecs = (args.UPDATE_SPECS_PATH as string) || UPDATE_SPECS_PATH;

  const contractUpdates = loadContractUpdateRegistry(
    parseObjFromFile(updateSpecs)
  );

  const extraContractFactories = extraBindingsPath
    ? await import(extraBindingsPath)
    : null;

  console.log("extra contracts", extraContractFactories);
  console.log(
    "extra contracts",
    extraContractFactories[`GnosisSafe0815__factory`]
  );
  updateContractsForChain(
    chain,
    accounts.mustGet(chain.chainName),
    ContractRegistryLoader.emptySingle(),
    contractUpdates,
    getOutputLogger(),
    false,
    false,
    true,
    extraContractFactories,
    extraArtifactsPath,
  );
}
main();
