package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

// RSA es muy lento en comparación con ECC (curvas elípticas)
// ECC con 512 bits de key equivale a una key de 15360 de RSA
// La seguridad de RSA está en factorizar grandes números
func encryptWithLib(pubKey *rsa.PublicKey, msg []byte) ([]byte, error) {
	hash := sha256.New()
	_, err := hash.Write(msg)
	if err != nil {
		return nil, err
	}
	randomReader := rand.Reader
	res, err := rsa.EncryptOAEP(hash, randomReader, pubKey, msg, nil)
	if err != nil {
		return nil, err
	}
	return res, nil 
}




