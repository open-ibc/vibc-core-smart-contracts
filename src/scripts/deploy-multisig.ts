#!/usr/bin/env node
import { ethers } from "ethers";
import { AccountRegistry, parseObjFromFile } from "..";
import { newSafeFromOwner } from "../multisig/safe";

import {
  parseMultiSigInitArgsFromCLI,
  saveMultisigAddressToAccountsSpec,
} from "../utils/io";

async function main() {
  const { rpcUrl, owners, initiator, accountsSpecPath, threshold } =
    await parseMultiSigInitArgsFromCLI();

  const accountConfigFromYaml = {
    name: "multisig-accounts",
    registry: parseObjFromFile(accountsSpecPath),
  };

  const accounts = AccountRegistry.loadMultiple([
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
    threshold,
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
