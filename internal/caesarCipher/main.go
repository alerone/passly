package caesarcipher

import "strings"

func encrypt(plaintext string, key int) string {
	return crypt(plaintext, key)
}

func decrypt(ciphertext string, key int) string {
	return crypt(ciphertext, -key)
}

func crypt(text string, key int) string {
	out := ""
	for _, c := range text {
		out = out + getOffsetChar(c, key)
	}
	return out
}

func getOffsetChar(c rune, offset int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	index := strings.Index(alphabet, string(c))
	index = (index + offset) % len(alphabet)
	if index < 0 {
		index = index + len(alphabet)
	}

	return string(alphabet[index])
}
