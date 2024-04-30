#!/usr/bin/env node
import { ContractRegistryLoader, parseObjFromFile } from "..";
import { loadTxRegistry } from "../evm/schemas/tx";
import { sendTxToChain } from "../tx";

import { getOutputLogger } from "../utils/cli";
import { UPGRADE_SPECS_PATH } from "../utils/constants";
import { parseArgsFromCLI } from "../utils/io";

async function main() {
  const { chain, accounts, args } = await parseArgsFromCLI();
  const upgradeSpecs =
    (args.UPGRADE_SPECS_PATH as string) || UPGRADE_SPECS_PATH;

  const upgradeTxs = loadTxRegistry(parseObjFromFile(upgradeSpecs));

  sendTxToChain(
    chain,
    accounts.mustGet(chain.chainName),
    ContractRegistryLoader.emptySingle(),
    upgradeTxs.subset(),
    getOutputLogger(),
    false
  );
}
main();
