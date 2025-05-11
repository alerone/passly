package aes

import (
	"crypto/aes"
	"crypto/cipher"
)

func decrypt(key, ciphertext, nonce []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err 
	}

	gcm, err := cipher.NewGCM(block)

	plaintext, err = gcm.Open(plaintext, nonce, ciphertext, nil)
	if err != nil {
		return nil, err 
	}

	return plaintext, nil
}

