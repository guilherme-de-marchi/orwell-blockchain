package blockchain

import (
	"log"

	"github.com/Guilherme-De-Marchi/orwell-blockchain/block"
)

type Iterator struct {
	CurrentHash []byte
	Chain       *Blockchain
}

func (iter *Iterator) Next() (*block.Block, error) {
	if iter.CurrentHash == nil {
		return nil, nil
	}

	log.Println("hey")
	b, err := iter.Chain.GetBlock(iter.CurrentHash)
	if err != nil {
		return nil, err
	}

	iter.CurrentHash = b.PrevHash
	return b, nil
}
