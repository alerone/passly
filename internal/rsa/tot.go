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
	p1 := new(big.Int)
	q1 := new(big.Int)
	big1 := big.NewInt(1)
	p1 = p1.Sub(p, big1)
	q1 = q1.Sub(q, big1)

	p1 = p1.Mul(p1, q1)
	return p1
}

func getE(tot *big.Int) *big.Int {
	res := new(big.Int)
	big1 := big.NewInt(1)
	for {
		res, err := crand.Int(randReader, tot)
		if err != nil {
			continue
		}

		if res.Cmp(big1) == 0 {
			continue
		}

		if gcd(res, tot).Cmp(big1) == 0 {
			break
		}
	}

	return res
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
