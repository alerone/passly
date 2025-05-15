package salt

import (
	"crypto/rand"
	"crypto/sha256"
)

func generateSalt(length int) ([]byte, error) {
	r := rand.Reader
	buf := make([]byte, length)
	_, err := r.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func hashPassword(password, salt []byte) []byte {
	h := sha256.New()
	psalt := append(password, salt...)
	h.Write(psalt)
	return h.Sum(nil)
}
