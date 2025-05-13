package main

import (
	"encoding/hex"
	"fmt"

	"github.com/alerone/passly/internal/hash"
)


func main() {

	gpas := make(map[string]float64)

	names := []string{"Lane","Yoojin","Jonny","Christine"}

	gpas[names[0]] = 3.8
	gpas[names[1]] = 3.5
	gpas[names[2]] = 2.0
	gpas[names[3]] = 4.5

	
	for _, name := range names {
		gpa := gpas[name]
		msg := fmt.Sprintf("%v has a GPA of %v\n", name, gpa)
		fmt.Print(msg)
		hashVal := hash.ToyHash([]byte(msg))
		hashed := make([]byte, len(hashVal))
		copy(hashed, hashVal[:])
		fmt.Println(hex.EncodeToString(hashed))
	}
}
