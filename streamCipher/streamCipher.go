package streamcipher

func crypt(textCh, keyCh <- chan byte, result chan <- byte) {
	defer close(result)
	for {
		text, okText := <- textCh
		key, okKey := <- keyCh
		if !okKey || !okText{
			return
		}

		result <- text ^ key
	}
}
