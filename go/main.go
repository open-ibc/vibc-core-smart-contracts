package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

// Header represents a block header in the Ethereum blockchain.
type HeaderWithoutParentHash struct {
	UncleHash   common.Hash    `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    common.Address `json:"miner"`
	Root        common.Hash    `json:"stateRoot"        gencodec:"required"`
	TxHash      common.Hash    `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash common.Hash    `json:"receiptsRoot"     gencodec:"required"`
	Bloom       types.Bloom          `json:"logsBloom"        gencodec:"required"`
	Difficulty  *big.Int       `json:"difficulty"       gencodec:"required"`
	Number      *big.Int       `json:"number"           gencodec:"required"`
	GasLimit    uint64         `json:"gasLimit"         gencodec:"required"`
	GasUsed     uint64         `json:"gasUsed"          gencodec:"required"`
	Time        uint64         `json:"timestamp"        gencodec:"required"`
	Extra       []byte         `json:"extraData"        gencodec:"required"`
	MixDigest   common.Hash    `json:"mixHash"`
	Nonce       types.BlockNonce     `json:"nonce"`

	// BaseFee was added by EIP-1559 and is ignored in legacy headers.
	BaseFee *big.Int `json:"baseFeePerGas" rlp:"optional"`

	// WithdrawalsHash was added by EIP-4895 and is ignored in legacy headers.
	WithdrawalsHash *common.Hash `json:"withdrawalsRoot" rlp:"optional"`

	// BlobGasUsed was added by EIP-4844 and is ignored in legacy headers.
	BlobGasUsed *uint64 `json:"blobGasUsed" rlp:"optional"`

	// ExcessBlobGas was added by EIP-4844 and is ignored in legacy headers.
	ExcessBlobGas *uint64 `json:"excessBlobGas" rlp:"optional"`

	// ParentBeaconRoot was added by EIP-4788 and is ignored in legacy headers.
	ParentBeaconRoot *common.Hash `json:"parentBeaconBlockRoot" rlp:"optional"`
}


func getBeforeAndAfterBytes(header *types.Header)([]byte, []byte){
	// Rlp encode bytes as normal
	var encodedHeader bytes.Buffer
	rlp.Encode(&encodedHeader, header)
	headerBytes:= encodedHeader.Bytes()

	// now we find and remoe the hash from the bytes
	hash := header.ParentHash.Bytes()

	index:= bytes.Index(headerBytes, hash)
	fmt.Printf("index of hash in bytes: %x\n", index)

	return headerBytes[:index], headerBytes[index+32:]
}

func calculateHeaderHashLikeInContract(prevHeader *common.Hash, beforeBytes []byte, afterBytes []byte) []byte {
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(beforeBytes)
	hasher.Write(prevHeader.Bytes())
	hasher.Write(afterBytes)
	return hasher.Sum(nil)
}

func removeHeaderHash(header *types.Header) (HeaderWithoutParentHash){

		return HeaderWithoutParentHash{
			UncleHash:         header.UncleHash,
			Coinbase:          header.Coinbase,
			Root:              header.Root,
			TxHash:            header.TxHash,
			ReceiptHash:       header.ReceiptHash,
			Bloom:             header.Bloom,
			Difficulty:        header.Difficulty,
			Number:            header.Number,
			GasLimit:          header.GasLimit,
			GasUsed:           header.GasUsed,
			Time:              header.Time,
			Extra:             header.Extra,
			MixDigest:         header.MixDigest,
			Nonce:             header.Nonce,
			BaseFee:           header.BaseFee,
			WithdrawalsHash:   header.WithdrawalsHash,
			BlobGasUsed:       header.BlobGasUsed,
			ExcessBlobGas:     header.ExcessBlobGas,
			ParentBeaconRoot:  header.ParentBeaconRoot,
		}
}


func calculateHeaderHash(header *types.Header) []byte {
	var encodedHeader bytes.Buffer
	rlp.Encode(&encodedHeader, header)
	fmt.Printf("normal encoded header: %x \n", encodedHeader.Bytes())
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(encodedHeader.Bytes())
	return hasher.Sum(nil)
}


func main() {
	client, err := ethclient.Dial("https://opt-sepolia.g.alchemy.com/v2/lTfokMxkVuwtrXEz07t7GRmWA1oFGftf") // Replace with your Infura project ID
	if err != nil {
		log.Fatal(err)
	}

	blocks := []*big.Int{big.NewInt(1), big.NewInt(2)}

	for _, blockNumber := range blocks {
		block, err := client.BlockByNumber(context.Background(), blockNumber)
		if err != nil {
			log.Fatal(err)
		}

		header := block.Header()
		

		fmt.Printf("Block Number: %x\n", header.Number)
		fmt.Printf("Prevhash: %x \n", header.ParentHash)
		fmt.Printf("stored hash of block %x \n", block.Hash() )
		
		before, after := getBeforeAndAfterBytes(header)
		// fmt.Printf("before and after bytes %x \n %x \n", before, after )
		// fmt.Printf("calculated hash of this block: %x \n", calculateHeaderHash(header))
		nonParentHeader:= removeHeaderHash(header)
		var otherBytes bytes.Buffer
		rlp.Encode(&otherBytes, nonParentHeader)
		// otherBytesSlice := otherBytes.Bytes()
		// fmt.Printf("other bytes rlp encoded %x \n", otherBytesSlice)
		fmt.Printf("calculated hash of this block via contract: %x \n", calculateHeaderHashLikeInContract(&header.ParentHash, before, after))

		
	}

}
