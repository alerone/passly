package blockcipher

func padWithZeros(block []byte, desiredSize int) []byte {
	res := make([]byte, desiredSize)
	copy(res, block)
	return res
}
