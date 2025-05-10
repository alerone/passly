package main

import (
	"fmt"
	"strings"
)


func base8Char(bits byte) string {
	const base8Alphabet = "ABCDEFGH"
	num := int(bits)
	if num >= len(base8Alphabet) {
		return ""
	}

	return string(base8Alphabet[num])
}

func getHexString(b []byte) string {
	format := make([]string, len(b))
	for i, byt := range b {
		format[i] = fmt.Sprintf("%02x", byt)
	}

	return strings.Join(format, ":")
}

func getBinaryString(b []byte) string {
	format := make([]string, len(b))
	for i, byt := range b {
		format[i] = fmt.Sprintf("%08b", byt)
	}

	return strings.Join(format, ":")
}
