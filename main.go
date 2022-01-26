package main

import (
	"log"

	"github.com/Guilherme-De-Marchi/orwell-blockchain/blockchain"
	"github.com/Guilherme-De-Marchi/orwell-blockchain/proof"
	"github.com/dgraph-io/badger/v3"
)

func main() {
	db, err := badger.Open(badger.DefaultOptions("./tmp/blockchain"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	p := proof.NewPoW(2)
	chain := blockchain.NewBlockchain(db, p)

	_, err = chain.AddNewBlock([]byte("genesis"))
	if err != nil {
		log.Fatal(err)
	}
	chain.GetBlock(chain.LastHash)

	_, err = chain.AddNewBlock([]byte("block 1"))
	if err != nil {
		log.Fatal(err)
	}
	chain.GetBlock(chain.LastHash)

	_, err = chain.AddNewBlock([]byte("block 2"))
	if err != nil {
		log.Fatal(err)
	}
	chain.GetBlock(chain.LastHash)

	_, err = chain.AddNewBlock([]byte("block 3"))
	if err != nil {
		log.Fatal(err)
	}
	chain.GetBlock(chain.LastHash)

	_, err = chain.AddNewBlock([]byte("block 4"))
	if err != nil {
		log.Fatal(err)
	}
	chain.GetBlock(chain.LastHash)
}
