package main


func base8Char(bits byte) string {
	const base8Alphabet = "ABCDEFGH"
	num := int(bits)
	if num >= len(base8Alphabet) {
		return ""
	}

	return string(base8Alphabet[num])
}
