#!/usr/bin/env node
import { ContractRegistryLoader, deployToChain, parseObjFromFile } from "..";

import { getMainLogger } from "../utils/cli";
import { DEPLOY_SPECS_PATH } from "../utils/constants";
import { parseArgsFromCLI } from "../utils/io";

async function main() {
  const { chain, accounts, args, deploySpecs } = await parseArgsFromCLI();
  const contracts = ContractRegistryLoader.loadSingle(
    parseObjFromFile(deploySpecs)
  );
  deployToChain(
    chain,
    accounts.mustGet(chain.chainName),
    contracts.subset(),
    getMainLogger(),
    false,
    false,
    true
  );
}
main();
