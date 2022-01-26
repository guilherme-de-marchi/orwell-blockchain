package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/gob"
	"fmt"
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

func (b *Block) Print() {
	fmt.Printf("PrevHash: %x\n", b.PrevHash)
	fmt.Printf("Data: %s\n", b.Data)
	fmt.Printf("Hash: %x\n", b.Hash)
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

func (b *Block) Serialize() ([]byte, error) {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)
	err := encoder.Encode(b)
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
