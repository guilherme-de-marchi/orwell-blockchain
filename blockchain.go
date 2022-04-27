package main

import (
	"github.com/dgraph-io/badger/v3"
)

type Blockchain struct {
	LastBlock *Block
	db        *badger.DB
}

func NewBlockchain(db *badger.DB, root *Block) (*Blockchain, error) {
	bc := &Blockchain{db: db}
	return bc, bc.AddBlock(root)
}

func LoadBlockchain(db *badger.DB) (*Blockchain, error) {
	bc := &Blockchain{db: db}
	lastB, err := bc.GetLastBlock()
	bc.LastBlock = lastB
	return bc, err
}

func (bc *Blockchain) GetLastBlock() (*Block, error) {
	var b *Block
	err := bc.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("last_block"))
		if err != nil {
			return err
		}
		var hash []byte
		err = item.Value(func(val []byte) error {
			hash = append([]byte{}, val...)
			return nil
		})
		if err != nil {
			return err
		}
		b, err = bc.GetBlock(hash)
		return err
	})
	return b, err
}

func (bc *Blockchain) GetBlock(key []byte) (*Block, error) {
	var b *Block
	err := bc.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			b, err = DeserializeBlock(append([]byte{}, val...))
			return err
		})
	})
	return b, err
}

func (bc *Blockchain) AddBlock(b *Block) error {
	err := bc.db.Update(func(txn *badger.Txn) error {
		s, err := b.Serialize()
		if err != nil {
			return err
		}
		err = txn.Set(b.Hash[:], s)
		if err != nil {
			return err
		}
		return txn.Set([]byte("last_block"), b.Hash[:])
	})
	if err != nil {
		return err
	}
	bc.LastBlock = b
	return nil
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	return NewIterator(bc)
}
