package otp

func crypt(plaintext, key []byte) []byte {
	out := make([]byte, len(plaintext))
	if len(plaintext) > len(key) {
		return out
	}
	for i, val := range plaintext {
		out[i] = val ^ key[i]
	}

	return out
}

// encrypt function uses the crypt function for XOR encryption
func encrypt(plaintext, key []byte) []byte {
	return crypt(plaintext, key)
}

// decrypt function uses the crypt function for XOR decryption
func decrypt(ciphertext, key []byte) []byte {
	return crypt(ciphertext, key)
}

