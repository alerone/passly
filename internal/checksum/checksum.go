package checksum

import (
	"crypto/sha256"
	"encoding/hex"
)

func checksumMatches(message, checksum string) bool {
	h := sha256.New()
	h.Write([]byte(message))

	hashedMsg := hex.EncodeToString(h.Sum(nil))
	return hashedMsg == checksum
}
