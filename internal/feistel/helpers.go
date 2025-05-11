package feistel

import "crypto/sha256"

func xor(lhs, rhs []byte) []byte {
	res := []byte{}

	for i := range lhs {
		res = append(res, lhs[i] ^ rhs[i])
	}

	return res
}

func hash(first, second []byte, outputLength int) []byte {
	h := sha256.New()
	h.Write(append(first, second...))
	return h.Sum(nil)[:outputLength]
}

func reverse[T any](s []T) []T {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
