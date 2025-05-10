package roundkey

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func deriveRoundKey(masterKey [4]byte, roundNumber int) [4]byte {
	var roundKey [4]byte
	for i, val := range masterKey {
		roundKey[i] = val ^ byte(roundNumber)
	}

	return roundKey
}

func intToBytes(num int) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, int64(num))
	if err != nil {
		fmt.Println(err)
	}
	return buf.Bytes()
}
