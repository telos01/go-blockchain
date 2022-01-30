package main

import (
	"fmt"
	"strconv"

	"github.com/telos01/go-chain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("Second block")
	chain.AddBlock("Third block")
	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

		pow := blockchain.NewProof(block)
		fmt.Printf("PoW:%s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
