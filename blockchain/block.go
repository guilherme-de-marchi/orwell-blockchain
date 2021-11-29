package blockchain

import (
	"crypto/rsa"
	"time"

	"github.com/Guilherme-De-Marchi/orwells-blockchain/security"
)

type Block struct {
	Data                               string
	Timestamp                          int64
	Signature, Hash, PreviousBlockHash []byte
}

func NewBlock(data string, previousBlockHash []byte) (*Block, error) {
	block := Block{
		Data:              data,
		Timestamp:         time.Now().UnixMicro(),
		PreviousBlockHash: previousBlockHash,
	}

	// Block's hash = SHA256(data, timestamp, previous_block_hash)
	hash, err := security.HashfySHA256(
		[]byte(block.Data),
		[][]byte{[]byte(string(block.Timestamp)), block.PreviousBlockHash},
	)
	if err != nil {
		return nil, err
	}
	block.Hash = hash

	return &block, nil
}

func (blk *Block) Sign(priv *rsa.PrivateKey) error {
	signature, err := security.Sign(blk.Hash, priv)
	if err != nil {
		return err
	}
	blk.Signature = signature
	return nil
}
