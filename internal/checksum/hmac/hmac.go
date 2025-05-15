package hmac

import (
	"crypto/sha256"
	"encoding/hex"
)

func hmac(message, key string) string {
	half := len(key) / 2
	firstHalf := key[:half]
	secondHalf := key[half:]

	h := sha256.New()
	h.Write([]byte(secondHalf + message))
	secCheck := h.Sum(nil)
	h.Reset()
	h.Write([]byte(firstHalf + string(secCheck)))


	return hex.EncodeToString(h.Sum(nil))
}
