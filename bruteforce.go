package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

func findKey(encrypted []byte, decrypted string) ([]byte, error) {
	maxKey := math.Exp2(24)
	for key := range int(maxKey) {
		cleanKey := intToBytes(key)
		tryDecrypt := crypt(encrypted, cleanKey)
		if string(tryDecrypt) == decrypted {
			return cleanKey, nil
		}
	}
	return make([]byte, 0), fmt.Errorf("key not found")
}


// Helper function: crypt performs XOR-based encryption/decryption
func crypt(dat, key []byte) []byte {
	final := []byte{}
	for i, d := range dat {
		final = append(final, d^key[i])
	}
	return final
}

// Helper function: intToBytes converts an integer to a 3-byte slice (little-endian)
func intToBytes(num int) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, int64(num))
	if err != nil {
		return nil
	}
	bs := buf.Bytes()
	if len(bs) > 3 {
		return bs[:3]
	}
	return bs
}

