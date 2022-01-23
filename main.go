package main

import (
	"fmt"

	"github.com/Guilherme-De-Marchi/orwell-blockchain/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()
	bc.Init()

	bc.AddNewBlock([]byte("block 1"))
	bc.AddNewBlock([]byte("block 2"))
	bc.AddNewBlock([]byte("block 3"))

	for _, b := range bc.Blocks {
		fmt.Printf("PrevHash: %x\n", b.PrevHash)
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n\n", b.Hash)
	}
}
