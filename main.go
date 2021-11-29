package main

import (
	"encoding/hex"
	"fmt"

	"github.com/Guilherme-De-Marchi/orwells-blockchain/blockchain"
	"github.com/Guilherme-De-Marchi/orwells-blockchain/security"
)

func main() {
	key, err := security.GenerateRSAKey(2048)
	if err != nil {
		fmt.Println(err)
	}

	genesis, err := blockchain.NewBlock("testing", []byte(""))
	if err != nil {
		fmt.Println(err)
	}

	genesis.Sign(key)

	fmt.Println(
		genesis.Data,
		hex.EncodeToString(genesis.Hash),
		hex.EncodeToString(genesis.PreviousBlockHash),
		genesis.Timestamp,
		hex.EncodeToString(genesis.Signature),
	)

	verif := security.VerifySignature(genesis.Hash, genesis.Signature, &key.PublicKey)
	fmt.Println(verif)
}
