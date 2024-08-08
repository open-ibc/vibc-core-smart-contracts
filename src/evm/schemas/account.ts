import { z } from "zod";
import { ethers } from "ethers";
import fs from "fs";
import path from "path";
import { Registry } from "../../utils/registry";
import { parseZodSchema, renderString } from "../../utils/io";

const privateKey = z
  .object({
    name: z.string().min(1),
    // privateKey should be a hex string prefixed with 0x
    privateKey: z.string().min(1),
  })
  .strict();

const mnemonic = z
  .object({
    name: z.string().min(1),
    // a 12-word mnemonic; or more words per BIP-39 spec
    mnemonic: z.string().min(1),
    path: z.optional(z.string().min(1)),
    index: z.optional(z.number().int().min(0)),
  })
  .strict();

const singleSigAccount = z.union([privateKey, mnemonic]);

// Uninitialized User input type
const multisigConfig = z.object({
  name: z.string().min(1),
  privateKey: z.string().min(1),
  multisigAddress: z.string().min(1),
  chainId: z.bigint(),
});

// Type that loadEvmAccounts will return for multisig types
const multisigAccount = z
  .object({
    name: z.string().min(1),
    privateKey: privateKey,
    multisigAddress: z.string().min(1),
    chainId: z.bigint(),
    wallet: z.instanceof(ethers.Wallet),
  })
  .strict();

type Privatekey = z.infer<typeof privateKey>;
type Mnemonic = z.infer<typeof mnemonic>;
type SingleSigAccount = z.infer<typeof singleSigAccount>;
type MultiSigConfig = z.infer<typeof multisigConfig>;
type ParsedMultiSigWallet = z.infer<typeof multisigAccount>;
export type SendingAccount = Wallet | ParsedMultiSigWallet;

export const isPrivateKey = (account: any): account is Privatekey => {
  return privateKey.safeParse(account).success;
};
export const isMnemonic = (account: any): account is Mnemonic => {
  return mnemonic.safeParse(account).success;
  ``;
};

export const isSingleSigAccount = (
  account: any
): account is SingleSigAccount => {
  return singleSigAccount.safeParse(account).success;
};

export const isMultiSigConfig = (account: any): account is MultiSigConfig => {
  return multisigConfig.safeParse(account).success;
};

// Note: only use this for already declared accounts.
export const isParsedMultiSigWallet = (
  account: any
): account is ParsedMultiSigWallet => {
  return multisigAccount.safeParse(account).success;
};

// geth compatible keystore
const keyStore = z.object({
  dir: z.string().min(1),
  password: z.optional(z.string()),
});

export const EvmAccountsSchema = z.union([
  z.array(z.union([privateKey, mnemonic, multisigAccount])),
  keyStore,
]);

// ethers wallet with encryption
export type Wallet = ethers.Wallet | ethers.HDNodeWallet;

export class AccountRegistry extends Registry<SendingAccount> {
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

  constructor(
    r: Registry<SendingAccount>,
    private config: any[],
    name: string
  ) {
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
        address: isParsedMultiSigWallet(wallets[index])
          ? wallets[index].wallet
          : wallets[index],
        ...item,
      };
    });
  }
}

// load a Map of { [name: string]: Wallet } from EvmAccountsSchema object
export function loadEvmAccounts(config: any): Registry<SendingAccount> {
  const accountsConfig = config.map((c: any) => {
    const parsed = EvmAccountsSchema.safeParse(c);
    if (!parsed.success) {
      throw new Error(`Invalid account config: ${parsed.error.errors}`);
    }
    return parsed.data;
  });

  const walletMap = new Registry<SendingAccount>([]);
  if (Array.isArray(accountsConfig)) {
    for (const account of accountsConfig) {
      if (isSingleSigAccount(account)) {
        walletMap.set(account.name, createWallet(account));
      } else if (isMultiSigConfig(account)) {
        walletMap.set(
          account.name,
          createWallet({ privateKey: account.privateKey, name: account.name })
        );
      }
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
  for (const [name, account] of accountRegistry.entries()) {
    if (isParsedMultiSigWallet(account)) {
      newAccounts.set(name, account.wallet.connect(provider));
    } else {
      newAccounts.set(name, account.connect(provider));
    }
  }

  return newAccounts;
};

export function createWallet(opt: SingleSigAccount): Wallet {
  if (isPrivateKey(opt)) {
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

  // Otherwise is a mnemonic
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
