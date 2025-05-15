package ecdsa

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func createECDSAMessage(message string, privateKey *ecdsa.PrivateKey) (string, error) {
	h := sha256.New()
	h.Write([]byte(message))
	r := rand.Reader
	s, err := ecdsa.SignASN1(r, privateKey, h.Sum(nil))
	if err != nil {
		return "", err
	}

	return message + "." + hex.EncodeToString(s), nil
}
