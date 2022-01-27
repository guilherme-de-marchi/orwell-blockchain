package proof

import (
	"bytes"

	"github.com/Guilherme-De-Marchi/orwell-blockchain/block"
)

type Proofer interface {
	Apply(*block.Block)
	Validate(*block.Block) bool
}

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

func (p PoW) Apply(b *block.Block) {
	for {
		b.DeriveHash()
		if bytes.Equal(b.Hash[:p.Difficulty], p.Target) {
			break
		}
		b.Nonce++
	}
}

func (p PoW) Validate(b *block.Block) bool {
	return bytes.Equal(b.Hash[:p.Difficulty], p.Target)
}
