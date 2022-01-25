package proof

import (
	"bytes"

	"github.com/Guilherme-De-Marchi/orwell-blockchain/blockchain"
)

type PoW struct {
	Difficulty int
	Target     []byte
}

func NewPoW(diff int) PoW {
	return PoW{
		Difficulty: diff,
		Target:     make([]byte, diff),
	}
}

func (p PoW) Apply(b *blockchain.Block) {
	for {
		b.DeriveHash()
		if bytes.Equal(b.Hash[:p.Difficulty], p.Target) {
			break
		}
		b.Nonce++
	}
}

func (p PoW) Validate(b *blockchain.Block) bool {
	return bytes.Equal(b.Hash[:p.Difficulty], p.Target)
}
