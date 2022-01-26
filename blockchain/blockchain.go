package blockchain

import (
	"github.com/dgraph-io/badger/v3"
)

type Proofer interface {
	Apply(*Block)
	Validate(*Block) bool
}

type Blockchain struct {
	LastHash []byte
	Database *badger.DB
	Proof    Proofer
}

func NewBlockchain(db *badger.DB, p Proofer) *Blockchain {
	return &Blockchain{
		Database: db,
		Proof:    p,
	}
}

func (chain *Blockchain) GetBlock(hash []byte) (*Block, error) {
	if hash == nil {
		return new(Block), nil
	}

	var b *Block
	var bs []byte
	err := chain.Database.View(func(txn *badger.Txn) error {
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

	b, err = DeserializeBlock(bs)
	if err != nil {
		return nil, err
	}

	b.Print()

	return b, nil
}

func (chain *Blockchain) AddNewBlock(data []byte) (*Block, error) {
	b := NewBlock(chain.LastHash, data)
	chain.Proof.Apply(b)
	return chain.AddBlock(b)
}

func (chain *Blockchain) AddBlock(b *Block) (*Block, error) {
	bs, err := b.Serialize()
	if err != nil {
		return nil, err
	}

	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(b.Hash, bs)
		return err
	})
	if err != nil {
		return nil, err
	}

	chain.LastHash = b.Hash
	return b, nil
}
