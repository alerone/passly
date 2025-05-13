package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"hash"
)

type hasher struct {
	hash hash.Hash
}

func newHasher() *hasher {
	return &hasher{hash: sha256.New()}
}

func (h *hasher) Write(p string) (int, error){
	return h.hash.Write([]byte(p))
}

func (h *hasher) GetHex() (string){
	sum := h.hash.Sum(nil)
	return hex.EncodeToString(sum)
}




