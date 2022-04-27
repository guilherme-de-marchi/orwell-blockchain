package main

import (
	"fmt"
	"github.com/dgraph-io/badger/v3"
)

func main() {
	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	bc, err := NewBlockchain(db, NewBlock([32]byte{}, "root"))
	if err != nil {
		panic(err)
	}
	fmt.Println(bc.GetLastBlock())
	err = bc.AddBlock(NewBlock(bc.LastBlock.Hash, "block 2"))
	if err != nil {
		panic(err)
	}
	fmt.Println(bc.GetLastBlock())
	fmt.Println("#################")
	iter := bc.Iterator()
	for ok := true; ok; {
		fmt.Println(iter.Value())
		ok, err = iter.Next()
		if err != nil {
			panic(err)
		}
	}
}
