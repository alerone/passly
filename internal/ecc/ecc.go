package ecc

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

func genKeys() (pubKey *ecdsa.PublicKey, privKey *ecdsa.PrivateKey, err error) {
	randReader := rand.Reader
	privKey, err = ecdsa.GenerateKey(elliptic.P256(), randReader)
	if err != nil {
		return nil, nil, err
	}
	
	return &privKey.PublicKey, privKey, nil
}
