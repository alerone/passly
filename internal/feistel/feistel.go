package feistel

func feistel(msg []byte, roundKeys [][]byte) []byte {
	nextLHS := msg[:len(msg)/2]
	nextRHS := msg[len(msg)/2:]
	for _, key := range roundKeys {
		aux := nextRHS
		nextRHS = xor(nextLHS, hash(aux, key, len(key)))
		nextLHS = aux
	}

	return append(nextRHS, nextLHS...)
}
