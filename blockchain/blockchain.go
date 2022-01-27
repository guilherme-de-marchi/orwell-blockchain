package blockchain

import (
	"log"

	"github.com/dgraph-io/badger/v3"

	"github.com/Guilherme-De-Marchi/orwell-blockchain/block"
	"github.com/Guilherme-De-Marchi/orwell-blockchain/proof"
)

type Blockchain struct {
	LastHash []byte
	Database *badger.DB
	Proof    proof.Proofer
}

func NewBlockchain(db *badger.DB, p proof.Proofer) *Blockchain {
	return &Blockchain{
		Database: db,
		Proof:    p,
	}
}

func (chain *Blockchain) Load() error {
	var bs []byte
	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("last_hash"))
		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			bs = append([]byte{}, val...)
			return nil
		})
	})
	if err != nil {
		return err
	}

	chain.LastHash = bs
	return nil
}

func (chain *Blockchain) GetBlock(hash []byte) (*block.Block, error) {
	if hash == nil {
		return new(block.Block), nil
	}

	var bs []byte
	err := chain.Database.View(func(txn *badger.Txn) error {
		log.Println("haha")
		item, err := txn.Get(hash)
		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			bs = append([]byte{}, val...)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}

	b, err := block.DeserializeBlock(bs)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (chain *Blockchain) AddNewBlock(data []byte) (*block.Block, error) {
	b := block.NewBlock(chain.LastHash, data)
	chain.Proof.Apply(b)
	return chain.AddBlock(b)
}

func (chain *Blockchain) AddBlock(b *block.Block) (*block.Block, error) {
	bs, err := b.Serialize()
	if err != nil {
		return nil, err
	}

	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(b.Hash, bs)
		if err != nil {
			return err
		}

		err = txn.Set([]byte("last_hash"), b.Hash)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	chain.LastHash = b.Hash
	return b, nil
}

func (chain *Blockchain) GetIterator() *Iterator {
	return &Iterator{
		CurrentHash: chain.LastHash,
	}
}
