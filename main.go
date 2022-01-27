package main

import (
	"log"
	"os"

	"github.com/Guilherme-De-Marchi/orwell-blockchain/blockchain"
	"github.com/Guilherme-De-Marchi/orwell-blockchain/cli"
	"github.com/Guilherme-De-Marchi/orwell-blockchain/proof"
	"github.com/dgraph-io/badger/v3"
)

func main() {
	defer os.Exit(0)
	db, err := badger.Open(badger.DefaultOptions("./tmp/blockchain"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	p := proof.NewPoW(2)
	chain := blockchain.NewBlockchain(db, p)
	chain.Load()

	c := cli.NewCommandLine(chain)
	c.Run()
}
