package rsa

import (
	crand "crypto/rand"
	"math/big"
)

// el totiente es el número de enteros positivos que son coprimos con un numero n
// coprimos son los números que descompuestos solo comparten el 1 GCD (greatest common divisor)
// como primo divisible. este calculo si tenemos ya los dos primos que dan un numero es sencillo:
// n = (p-1) * (q-1)
func getTot(p, q *big.Int) *big.Int {
	tot := new(big.Int)
	tot.Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))
	return tot
}

func getE(tot *big.Int) *big.Int {
	totMinusTwo := new(big.Int)
	totMinusTwo.Sub(tot, big.NewInt(2))

	e, _ := crand.Int(randReader, totMinusTwo)
	e.Add(e, big.NewInt(2))
	for gcd(e, tot).Cmp(big.NewInt(1)) != 0 {
		e, _ = crand.Int(randReader, totMinusTwo)
		e.Add(e, big.NewInt(2))
	}
	return e
}

// Calcula el máximo común divisor de dos big-int (GCD en inglés)
func gcd(x, y *big.Int) *big.Int {
	xCopy := new(big.Int).Set(x)
	yCopy := new(big.Int).Set(y)

	for yCopy.Cmp(big.NewInt(0)) != 0 {
		xCopy, yCopy = yCopy, xCopy.Mod(xCopy, yCopy)
	}
	return xCopy
}
