package rsa

import "math/big"

func encrypt(m, e, n *big.Int) *big.Int {
	res := new(big.Int)
	res = res.Exp(m,e,n)
	return res
}
