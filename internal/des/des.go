package des

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

func encrypt(key, plaintext []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("error while creating cipher: %s", err.Error()) 
	}

	paddedText := padMsg(plaintext, block.BlockSize())
	cipherText := make([]byte, block.BlockSize() + len(paddedText))
	iv := cipherText[:block.BlockSize()]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[block.BlockSize():], paddedText)
	return cipherText, nil
}

func padMsg(plaintext []byte, blockSize int) []byte {
	lastBlockSize := len(plaintext) % blockSize
	if lastBlockSize == 0 {
		return plaintext
	}

	lastBlock := plaintext[len(plaintext)-lastBlockSize:]
	padded := padWithZeros(lastBlock, blockSize)
	return append(plaintext[:len(plaintext)-lastBlockSize], padded...)
}


func decrypt(key, ciphertext []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < des.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:des.BlockSize]
	ciphertext = ciphertext[des.BlockSize:]
	if len(ciphertext)%des.BlockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	return ciphertext, nil
}

func padWithZeros(block []byte, desiredSize int) []byte {
	for len(block) < desiredSize {
		block = append(block, 0)
	}
	return block
}

