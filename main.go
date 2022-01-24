package main

import (
	"fmt"

	"github.com/Guilherme-De-Marchi/orwell-blockchain/blockchain"
	"github.com/Guilherme-De-Marchi/orwell-blockchain/proof"
)

func main() {
	bc := blockchain.NewBlockchain()
	bc.Init()

	bc.AddNewBlock([]byte("block 1"))
	bc.AddNewBlock([]byte("block 2"))
	bc.AddNewBlock([]byte("block 3"))

	for _, b := range bc.Blocks {
		p := proof.NewProofOfWork(b, 2)
		p.Run()

		fmt.Printf("PrevHash: %x\n", b.PrevHash)
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)
		fmt.Printf("Valid: %v\n\n", p.Validate())
	}
}
