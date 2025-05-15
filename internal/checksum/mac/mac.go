package mac

import (
	"crypto/sha256"
	"encoding/hex"
)

func macMatches(message, key, checksum string) bool {
	h := sha256.New()
	h.Write([]byte(message + key))
	return checksum == hex.EncodeToString(h.Sum(nil))
}
