package main

import (
	"fmt"
	"math/rand"
)

func generateRandomKey(length int) (string, error) {
	randReader := rand.New(rand.NewSource(0))


	buf := make([]byte, length)
	_, err := randReader.Read(buf)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", buf), nil
}
