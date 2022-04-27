package main

import "github.com/dgraph-io/badger/v3"

type BlockchainIterator struct {
	chain   *Blockchain
	current *Block
}

func NewIterator(bc *Blockchain) *BlockchainIterator {
	return &BlockchainIterator{
		chain:   bc,
		current: bc.LastBlock,
	}
}

func (iter *BlockchainIterator) Next() (bool, error) {
	b, err := iter.chain.GetBlock(iter.current.PrevHash[:])
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return false, nil
		}
		return false, err
	}
	iter.current = b
	return true, nil
}

func (iter *BlockchainIterator) Value() *Block {
	return iter.current
}
