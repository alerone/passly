package hash

import "math/bits"

func ToyHash(input []byte) [4]byte {
	modified := make([]byte, len(input))
	for i, bVal := range input{
		modified[i] = bits.RotateLeft8(bVal, 3)
		modified[i] = bVal << 2
	}
	final := [4]byte{}
	for i, bVal := range modified {
		final[i%4] ^= bVal
	}

	return final
}
