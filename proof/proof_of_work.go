package proof

import (
	"bytes"

	"github.com/Guilherme-De-Marchi/orwell-blockchain/blockchain"
)

type ProofOfWork struct {
	Block      *blockchain.Block
	Difficulty int
	Target     []byte
}

func NewProofOfWork(b *blockchain.Block, diff int) *ProofOfWork {
	return &ProofOfWork{
		Block:      b,
		Difficulty: diff,
		Target:     make([]byte, diff),
	}
}

func (p *ProofOfWork) Run() {
	for {
		p.Block.Hash = p.Block.DeriveHash()
		if bytes.Equal(p.Block.Hash[:p.Difficulty], p.Target) {
			break
		}
		p.Block.Nonce++
	}
}

func (p *ProofOfWork) Validate() bool {
	return bytes.Equal(p.Block.Hash[:p.Difficulty], p.Target)
}
