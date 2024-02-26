package ecdsa

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
)

type Ecc struct {
	PriKey *ecdsa.PrivateKey
	PubKey *ecdsa.PublicKey
}

func (e *Ecc) EncodePriKey() (en string, err error) {
	encoded, err := x509.MarshalECPrivateKey(e.PriKey)

	if err != nil {
		return
	}

	keyBytes := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: encoded})

	en = string(keyBytes)
	return
}
