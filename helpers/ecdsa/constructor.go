package ecdsa

import (
	"crypto/x509"
	"encoding/pem"
)

func NewWithPriKey(k string) *Ecc {
	pemBlock, _ := pem.Decode([]byte(k))

	if pemBlock == nil {
		panic("pemBlock is nil, but it shouldn't be.")
	}

	pri, err := x509.ParseECPrivateKey(pemBlock.Bytes)

	if err != nil {
		panic(err)
	}

	return &Ecc{
		PriKey: pri,
		PubKey: &pri.PublicKey,
	}
}
