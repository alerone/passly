package iv

import (
	"math"
)


func nonceStrength(nonce []byte) int {
	return int(math.Exp2(float64(len(nonce)*8)))
}


