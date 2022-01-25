package blockchain

import (
	"fmt"
)

type Proofer interface {
	Apply(*Block)
	Validate(*Block) bool
}

type Blockchain struct {
	Blocks []*Block
	Proof  Proofer
}

func NewBlockchain(p Proofer) *Blockchain {
	return &Blockchain{
		Proof: p,
	}
}

func (bc *Blockchain) GetBlock(i int) (*Block, error) {
	if len(bc.Blocks) == 0 {
		if i == -1 {
			return &Block{}, nil
		}
		return nil, fmt.Errorf("void Blockchain.Blocks\nLen: %v", len(bc.Blocks))
	}
	if i == -1 {
		return bc.Blocks[len(bc.Blocks)-1], nil
	}
	if len(bc.Blocks) == 0 || len(bc.Blocks)-1 < i {
		return nil, fmt.Errorf("block index not in range of Blockchain.Blocks\nLen: %v\nIndex: %v", len(bc.Blocks), i)
	}
	return bc.Blocks[i], nil
}

func (bc *Blockchain) GetLastBlock() (*Block, error) {
	return bc.GetBlock(-1)
}

func (bc *Blockchain) AddBlock(b *Block) {
	bc.Blocks = append(bc.Blocks, b)
}

func (bc *Blockchain) AddNewBlock(data []byte) (*Block, error) {
	lastb, err := bc.GetLastBlock()
	if err != nil {
		return nil, err
	}

	b := NewBlock(lastb.Hash, data)
	bc.Proof.Apply(b)
	bc.AddBlock(b)

	return b, nil
}
