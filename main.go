package main

import (
	"log"

	"github.com/Guilherme-De-Marchi/orwell-blockchain/blockchain"
	"github.com/Guilherme-De-Marchi/orwell-blockchain/proof"
)

func main() {
	p := proof.NewPoW(2)
	bc := blockchain.NewBlockchain(p)

	_, err := bc.AddNewBlock([]byte("genesis"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = bc.AddNewBlock([]byte("block 1"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = bc.AddNewBlock([]byte("block 2"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = bc.AddNewBlock([]byte("block 3"))
	if err != nil {
		log.Fatal(err)
	}

	for _, b := range bc.Blocks {
		log.Printf("PrevHash: %x\n", b.PrevHash)
		log.Printf("Data: %s\n", b.Data)
		log.Printf("Hash: %x\n", b.Hash)
		log.Printf("Valid: %v\n\n", bc.Proof.Validate(b))
	}
}
