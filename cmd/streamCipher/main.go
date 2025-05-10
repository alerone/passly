package main

import (
	"fmt"
	"log"

	"github.com/alerone/passly/internal/streamCipher"
)

func main() {
	encrypted, err := streamCipher.Encrypt([]byte("hola"), []byte("caro"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("encrypted:", encrypted)
}
