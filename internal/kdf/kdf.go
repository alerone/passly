package kdf

import (
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	secure, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	if err != nil {
		return "", err
	}

	return string(secure), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}

