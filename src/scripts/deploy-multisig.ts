#!/usr/bin/env node
import { ethers } from "ethers";
import { SingleSigAccountRegistry, parseObjFromFile } from "..";
import { newSafeFromOwner } from "../multisig/safe";

import {
  parseMultisigInitArgsFromCLI,
  saveMultisigAddressToAccountsSpec,
} from "../utils/io";

async function main() {
  const { rpcUrl, owners, initiator, accountsSpecPath, threshold } =
    await parseMultisigInitArgsFromCLI();

  const accountConfigFromYaml = {
    name: "multisig-accounts",
    registry: parseObjFromFile(accountsSpecPath),
  };

  const accounts = SingleSigAccountRegistry.loadMultiple([
    accountConfigFromYaml,
  ]).mustGet("multisig-accounts");

  const senderPrivateKey = accounts.getSinglePrivateKeyFromAccount(initiator);
  if (!senderPrivateKey) {
    throw new Error(`Could not find private key for owner ${initiator}`);
  }

  const newSafeAddress = await newSafeFromOwner(
    rpcUrl,
    senderPrivateKey,
    owners,
    threshold
  );

  const provider = new ethers.JsonRpcProvider(rpcUrl);
  const chainId = (await provider.getNetwork()).chainId;

  await saveMultisigAddressToAccountsSpec(
    newSafeAddress,
    accountsSpecPath,
    initiator,
    chainId
  );
}

main();
