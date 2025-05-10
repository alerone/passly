package blockcipher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func getBlockSize(keyLen, cipherType int) (int, error) {
	key := make([]byte, keyLen)

	cType := getCipherTypeName(cipherType)

	var c cipher.Block
	var err error
	switch cType {
	case "AES":
		{
			c, err = aes.NewCipher(key)
			if err != nil {
				return 0, err
			}
		}
	case "DES":
		{
			c, err = des.NewCipher(key)
			if err != nil {
				return 0, err
			}
		}
	default:
		return 0, fmt.Errorf("invalid cipher type")
	}

	return c.BlockSize(), nil
}
