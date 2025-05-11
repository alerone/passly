package iv

import "math/rand"

func generateIV(length int) ([]byte, error) {
	buf := make([]byte, length)
	_, err := rand.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
