import { z } from "zod";
import { ethers } from "ethers";
import fs from "fs";
import path from "path";
import { Registry } from "../utils/registry";
import { parseZodSchema, renderString } from "../utils/io";

const privateKey = z.object({
  name: z.string().min(1),
  // privateKey should be a hex string prefixed with 0x
  privateKey: z.string().min(1),
});

const mnemonic = z.object({
  name: z.string().min(1),
  // a 12-word mnemonic; or more words per BIP-39 spec
  mnemonic: z.string().min(1),
  path: z.optional(z.string().min(1)),
  index: z.optional(z.number().int().min(0)),
});

// geth compatible keystore
const keyStore = z.object({
  dir: z.string().min(1),
  password: z.optional(z.string()),
});

export const EvmAccountsSchema = z.union([
  z.array(z.union([privateKey, mnemonic])),
  keyStore,
]);

// ethers wallet with encryption
export type Wallet = ethers.Wallet | ethers.HDNodeWallet;

export class AccountRegistry extends Registry<Wallet> {
  static load(config: any[], name: string): AccountRegistry {
    return new AccountRegistry(loadEvmAccounts(config), config, name);
  }

  static loadMultiple(registryItems: { name: string; registry: any }[]) {
    const result = new Registry([] as AccountRegistry[], {
      toObj: (t) => {
        return { name: t.name, registry: t.serialize() };
      },
    });
    for (const item of registryItems) {
      result.set(item.name, AccountRegistry.load(item.registry, item.name));
    }
    return result;
  }

  constructor(r: Registry<Wallet>, private config: any[], name: string) {
    super([], { nameInParent: name });
    for (const [name, wallet] of r.entries()) {
      this.set(name, wallet);
    }
  }

  // return the same config obj that was used to load the accounts, but filtered by current account names
  public serialize() {
    const wallets = this.toList();
    return this.config.map((item, index) => {
      return {
        name: item.name,
        privateKey: wallets[index].privateKey,
        address: wallets[index].address,
        ...item,
      };
    });
  }
}

// load a Map of { [name: string]: Wallet } from EvmAccountsSchema object
export function loadEvmAccounts(config: any): Registry<Wallet> {
  const accountsConfig = parseZodSchema(
    "EvmAccountsSchema",
    config,
    EvmAccountsSchema.parse
  );
  const walletMap = new Registry<Wallet>([]);
  if (Array.isArray(accountsConfig)) {
    for (const account of accountsConfig) {
      walletMap.set(account.name, createWallet(account));
    }
  } else if ("dir" in accountsConfig) {
    const files = fs.readdirSync(accountsConfig.dir);
    for (const file of files) {
      const filePath = path.join(accountsConfig.dir, file);
      const json = fs.readFileSync(filePath, "utf8");
      const wallet = ethers.Wallet.fromEncryptedJsonSync(
        json,
        accountsConfig.password ?? ""
      );
      walletMap.set(wallet.address, wallet);
    }
  } else {
    throw new Error(
      `invalid accounts config: ${JSON.stringify(accountsConfig)}`
    );
  }
  return walletMap;
}

// Connect all accounts to the provider
export const connectProviderAccounts = (
  accountRegistry: AccountRegistry,
  rpc: string
) => {
  const provider = ethers.getDefaultProvider(rpc);
  const newAccounts = accountRegistry.subset([]);
  for (const [name, wallet] of accountRegistry.entries()) {
    newAccounts.set(name, wallet.connect(provider));
  }

  return newAccounts;
};

export function createWallet(opt: {
  privateKey?: string;
  mnemonic?: string;
  path?: string;
  index?: number;
}): Wallet {
  if (opt.privateKey && typeof opt.privateKey === "string") {
    let renderedPrivatekey = opt.privateKey;
    if (!ethers.isHexString(renderedPrivatekey, 32)) {
      // check if is a valid private key - if not, see if it represents an env variable that represents a private key.
      try {
        renderedPrivatekey = renderString(opt.privateKey, process.env); // look up in env if not.
      } catch (e) {
        console.log("no valid private key found for account spec", e);
      }
    }
    return new ethers.Wallet(renderedPrivatekey);
  }
  if (opt.mnemonic && typeof opt.mnemonic === "string") {
    let wallet = ethers.Wallet.fromPhrase(opt.mnemonic);
    if (typeof opt.path === "string" && opt.path.length > 0) {
      wallet = ethers.HDNodeWallet.fromPhrase(
        opt.mnemonic,
        undefined,
        opt.path
      );
    }
    // if account.index is specified, derive the child wallet by index
    if (Number.isInteger(opt.index)) {
      wallet = wallet.deriveChild(opt.index!);
    }
    return wallet;
  }
  throw new Error(
    `invalid wallet config, must provide at least one of {privateKey, mnemonic}, but got: ${JSON.stringify(
      opt
    )}`
  );
}
