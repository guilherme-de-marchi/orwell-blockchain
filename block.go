package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
)

type Block struct {
	PrevHash, Hash [32]byte
	Data           string
	Nonce          uint64
}

func NewBlock(prevHash [32]byte, data string) *Block {
	b := &Block{
		PrevHash: prevHash,
		Data:     data,
	}
	b.Hash = b.getHash()
	return b
}

func (b *Block) getHash() [32]byte {
	nonce := make([]byte, 8)
	binary.LittleEndian.PutUint64(nonce, b.Nonce)
	info := bytes.Join(
		[][]byte{
			b.PrevHash[:],
			[]byte(b.Data),
			nonce,
		},
		[]byte{},
	)
	return sha256.Sum256(info)
}

func (b *Block) Serialize() ([]byte, error) {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(*b)
	if err != nil {
		return nil, err
	}
	return res.Bytes(), nil
}

func DeserializeBlock(data []byte) (*Block, error) {
	var b Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&b)
	if err != nil {
		return nil, err
	}
	return &b, nil
}
