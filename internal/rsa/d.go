package rsa

import "math/big"


func getD(e, tot *big.Int) *big.Int {
	d := new(big.Int)
	d.ModInverse(e, tot)
	return d
}
