package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

func encrypt(pubKey *rsa.PublicKey, msg []byte) ([]byte, error) {
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
