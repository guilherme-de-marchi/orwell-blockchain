package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
)

type Block struct {
	PrevHash []byte
	Data     []byte
	Hash     []byte
	Nonce    uint64
}

func NewBlock(prevHash, data []byte) *Block {
	return &Block{
		PrevHash: prevHash,
		Data:     data,
	}
}

func (b *Block) DeriveHash() {
	nonce := make([]byte, 8)
	binary.LittleEndian.PutUint64(nonce, b.Nonce)

	info := bytes.Join(
		[][]byte{
			b.PrevHash,
			b.Data,
			nonce,
		},
		[]byte{},
	)

	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}
