import SafeApiKit from "@safe-global/api-kit";
import Safe, { SafeFactory } from "@safe-global/protocol-kit";
import {
  OperationType,
  SafeTransactionDataPartial,
  TransactionResult,
} from "@safe-global/safe-core-sdk-types";
import { ethers } from "ethers";

/*
 * Init a safe owner from a given private key to a give RPC's network
 * Submits an evm tx to create a new safe proxy
 */
export const newSafeFromOwner = async (
  RPC_URL: string,
  ownerKey: string,
  owners: string[],
  threshold: number
) => {
  const safeFactory = await SafeFactory.init({
    provider: RPC_URL,
    signer: ownerKey,
  });

  const safe = await safeFactory.deploySafe({
    safeAccountConfig: { owners, threshold },
  });
  const safeAddress = await safe.getAddress();

  console.log(`Safe has been deployed to: ${safeAddress}`);

  return safeAddress;
};

/**
 *
 * @param data Propose a transaction from a signer account
 */
export async function proposeTransaction(
  safeAddress: string,
  toAddress: string,
  txData: string,
  proposerPrivateKey: string,
  chainId: bigint,
  rpcUrl: string,
  nonce: number,
  value = "0",
  operation = OperationType.Call
) {
  const apiKit = new SafeApiKit({ chainId });

  const txProposer = await Safe.init({
    provider: rpcUrl,
    signer: proposerPrivateKey,
    safeAddress: safeAddress,
  });

  console.log(
    `proposing transaction from ${txProposer} @safe ${safeAddress} to ${toAddress} with value: ${value} and data: ${txData}, and nonce : ${nonce}`
  );

  const txProposerAddress = new ethers.Wallet(proposerPrivateKey).address;

  const safeTransactionData: SafeTransactionDataPartial = {
    to: toAddress,
    value,
    data: txData,
    operation: operation,
    nonce,
  };

  const safeTransaction = await txProposer.createTransaction({
    transactions: [safeTransactionData],
  });

  const safeTxHash = await txProposer.getTransactionHash(safeTransaction);

  const proposerSignature = await txProposer.signHash(safeTxHash);

  // Propose transaction to the service
  await apiKit.proposeTransaction({
    safeAddress,
    safeTransactionData: safeTransaction.data,
    safeTxHash,
    senderAddress: txProposerAddress,
    senderSignature: proposerSignature.data,
  });

  console.log(
    `Transaction has been proposed with hash: ${safeTxHash} from address ${txProposerAddress}`
  );
  return await apiKit.getTransaction(safeTxHash);
}

/**
 * Execute a multisig tx from an account generated from proposeTransaction
 */
export const executeMultisigTx = async (
  safeAddress: string,
  executorPrivateKey: string,
  chainId: bigint,
  rpcUrl: string,
  pendingTransactionIndex: number
):Promise<TransactionResult> => {
  const apiKit = new SafeApiKit({
    chainId,
  });

  const executor = await Safe.init({
    provider: rpcUrl,
    signer: executorPrivateKey,
    safeAddress,
  });

  const transactions = await apiKit.getPendingTransactions(safeAddress);

  if (transactions.results.length <= pendingTransactionIndex) {
    throw new Error(
      `Invalid transaction index - trying to access index ${pendingTransactionIndex} but only ${transactions.results.length} pending txs in safe`
    );
  }

  const tx = transactions.results[pendingTransactionIndex];
  return await executor.executeTransaction(tx);
};

export async function fetchNonceFromSafeAddress(
  rpcUrl: string,
  safeAddress: string
) {
  const safe = await Safe.init({
    provider: rpcUrl,
    safeAddress,
  });

  return safe.getNonce();
}
