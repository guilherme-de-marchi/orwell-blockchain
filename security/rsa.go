package security

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func GenerateRSAKey(bits int) (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, bits)
}

func Sign(dataHash []byte, priv *rsa.PrivateKey) ([]byte, error) {
	return rsa.SignPSS(rand.Reader, priv, crypto.SHA256, dataHash, nil)
}

func VerifySignature(dataHash []byte, signature []byte, pub *rsa.PublicKey) error {
	return rsa.VerifyPSS(pub, crypto.SHA256, dataHash, signature, nil)
}

func HashfySHA256(data []byte, complements [][]byte) ([]byte, error) {
	dataHash := sha256.New()
	_, err := dataHash.Write(data)
	if err != nil {
		return nil, err
	}

	for _, v := range complements {
		dataHash.Sum(v)
	}
	return dataHash.Sum(nil), nil
}
