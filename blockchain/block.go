package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	PrevHash []byte
	Data     []byte
	Hash     []byte
}

func NewBlock(prevHash, data []byte) *Block {
	b := &Block{
		PrevHash: prevHash,
		Data:     data,
	}
	b.DeriveHash()

	return b
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.PrevHash, b.Data}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}
