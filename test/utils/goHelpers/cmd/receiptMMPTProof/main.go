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
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/ethereum/go-ethereum/triedb"
)

func main() {

	blockNumber := int64(20221924)
	logIdx := uint64(1)
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

	for i, receipt := range receipts {
		fmt.Println("receipt: ", i, receipt.TxHash)
	}

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

	trieFun := trie.NewEmpty(triedb.NewDatabase(rawdb.NewMemoryDatabase(), nil))
	derivableReceipts := types.Receipts(receipts)
	receiptTrieRoot := types.DeriveSha(derivableReceipts, trieFun)

	fmt.Println("receiptTrieRoot: from deriveSha ", receiptTrieRoot.Hex())
	fmt.Println("receiptTrieRoot from block ", block.ReceiptHash)

	if !bytes.Equal(receiptTrieRoot.Bytes(), block.ReceiptHash.Bytes()) {
		fmt.Println("derived receipt trie root does not match receipt trie root in eth header")
		return
	}

	// check that the receiptTrieRoot matches the one in the ethHeader

	targetReceiptProof := EthProof{}

	if err := trieFun.Prove(rlp.AppendUint64(nil, logIdx), &targetReceiptProof); err != nil {
		fmt.Println("error proving receipt: ", err)
		return
	}

	for key, value := range targetReceiptProof {
		fmt.Println("key:", hex.EncodeToString([]byte(key)))
		fmt.Println("value:", hex.EncodeToString(value))
	}

	var buf bytes.Buffer
	receipts[0].EncodeRLP(&buf)
	hexStr := hex.EncodeToString(buf.Bytes())
	fmt.Println("RLP-encoded receipt (hex):", hexStr)
}
