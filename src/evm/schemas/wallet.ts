import { ethers } from 'ethers';
import { z } from 'zod';

export const wallet = z.union([
  z.instanceof(ethers.Wallet),
  z.instanceof(ethers.HDNodeWallet),
]);

export const privateKey = z
  .object({
    name: z.string().min(1),
    // privateKey should be a hex string prefixed with 0x
    privateKey: z.string().min(1),
  })
  .strict();

export const mnemonic = z
  .object({
    name: z.string().min(1),
    // a 12-word mnemonic; or more words per BIP-39 spec
    mnemonic: z.string().min(1),
    path: z.optional(z.string().min(1)),
    index: z.optional(z.number().int().min(0)),
  })
  .strict();

export const singleSigAccount = z.union([privateKey, mnemonic]);

export type PrivateKey = z.infer<typeof privateKey>;
export type Mnemonic = z.infer<typeof mnemonic>;
export type SingleSigAccount = z.infer<typeof singleSigAccount>;

export const isPrivateKey = (account: unknown): account is PrivateKey=> {
  return privateKey.safeParse(account).success;
};

export const isMnemonic = (account: unknown): account is Mnemonic=> {
  return mnemonic.safeParse(account).success;
};

export const isSingleSigAccount = (
  account: unknown
): account is SingleSigAccount=> {
  return singleSigAccount.safeParse(account).success;
};

export const isWallet = (account: unknown): account is Wallet => {
  return wallet.safeParse(account).success;
}

export type Wallet = ethers.Wallet | ethers.HDNodeWallet;
