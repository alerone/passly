package main

import (
	"encoding/hex"
	"strings"
)

func getHexBytes(s string) ([]byte, error) {
	byteStrings := strings.Split(s, ":")
	sout := ""
	for _, val := range byteStrings {
		sout = sout + val
	}

	return hex.DecodeString(sout)
}
