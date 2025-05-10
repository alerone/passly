package streamCipher

import "fmt"

func crypt(textCh, keyCh <- chan byte, result chan <- byte) {
	defer close(result)
	index := 0
	for {
		text, textOk := <- textCh
		key, keyOk := <- keyCh
		if !keyOk || !textOk{
			return
		}

		fmt.Printf("Crypted byte: %d\n", index)
		result <- text ^ key
		index++
	}
}
