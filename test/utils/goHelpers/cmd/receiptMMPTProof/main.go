package main

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/trie/trienode"
	"github.com/ethereum/go-ethereum/triedb"
)

func main() {

	blockNumber := int64(20221924)
	receiptIdx := uint64(1)

	fmt.Println("blockNumber: ", blockNumber)

	rpcUrl := os.Getenv("OP_SEPOLIA_RPC_URL")
	alchemyRpc, err := ethclient.Dial(rpcUrl)
	if err != nil {
		fmt.Println("error dialing alchemyRpc: ", err)
		return
	}

	client := ethclient.NewClient(alchemyRpc.Client())

	ctxWithTimeout, _ := context.WithTimeout(context.Background(), 10000*time.Millisecond)

	receipts, err := client.BlockReceipts(ctxWithTimeout, rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber)))

	if err != nil {
		fmt.Println("error fetching receipts: ", err)
		return
	}

	ctxWithTimeout, _ = context.WithTimeout(context.Background(), 1000*time.Millisecond)

	block, err := client.HeaderByNumber(ctxWithTimeout, big.NewInt(blockNumber))
	if err != nil {
		fmt.Println("error fetching block: ", err)
		return
	}

	derivableReceipts := types.Receipts(receipts)

	// check that the receiptTrieRoot matches the one in the ethHeader

	receiptToProve := derivableReceipts[receiptIdx]

	targetReceiptProof := EthProof{}

	pset := trienode.NewProofSet()
	rlpEncodedKey := rlp.AppendUint64(nil, receiptIdx)

	tr := trie.NewEmpty(triedb.NewDatabase(rawdb.NewMemoryDatabase(), nil))
	trieRoot := types.DeriveSha(derivableReceipts, tr)

	if err := tr.Prove(rlpEncodedKey, pset); err != nil {
		fmt.Println("error proving receipt: ", err)
		return

	}

	if !bytes.Equal(tr.Hash().Bytes(), block.ReceiptHash.Bytes()) {
		fmt.Println("derived receipt trie root does not match receipt trie root in eth header")
		return
	}

	for key, value := range targetReceiptProof {
		fmt.Println("key:", hex.EncodeToString([]byte(key)))
		fmt.Println("value:", hex.EncodeToString(value))
	}

	var buf bytes.Buffer
	receiptToProve.EncodeRLP(&buf)
	rlpEncodedReceiptStr := hex.EncodeToString(buf.Bytes())
	plist := pset.List()
	leaf := plist[len(plist)-1]
	leafHash := crypto.Keccak256(leaf)

	second := crypto.Keccak256(plist[1])
	if !bytes.Contains(plist[0], second) {
		fmt.Println("output proof list not expected")
		return
	}
	if !bytes.Contains(plist[1], leafHash) {
		fmt.Println("output proof list not expected")
		return

	}

	fmt.Println("The following are the arguments that need to be passed into the MerkleTrie.verifyInclusionProof contract call: ")
	fmt.Printf("\n Key: %s \n", hex.EncodeToString(rlpEncodedKey))
	fmt.Printf("\n Value: %s \n", rlpEncodedReceiptStr)
	fmt.Printf("\n Proof :")
	for i, p := range plist {
		fmt.Printf("Proof array index %d: %s \n", i, hex.EncodeToString(p))
	}

	fmt.Printf(" \n Root: %s", hex.EncodeToString(trieRoot.Bytes()))

}
