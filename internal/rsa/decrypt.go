package rsa

import "math/big"

func decryptWithoutLib(c, d, n *big.Int) *big.Int {
	p := new(big.Int)
	p.Exp(c, d, n)
	return p
}
