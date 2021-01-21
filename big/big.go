package big

import "math/big"

// MulInt ...
func MulInt(bigs ...*big.Int) *big.Int {
	if len(bigs) == 0 {
		return big.NewInt(0)
	}
	mul := big.NewInt(1)
	for _, b := range bigs {
		mul = big.NewInt(0).Mul(mul, b)
	}
	return mul
}

// MulFloat ...
func MulFloat(bigs ...*big.Float) *big.Float {
	if len(bigs) == 0 {
		return big.NewFloat(0)
	}
	mul := big.NewFloat(1)
	for _, b := range bigs {
		mul = big.NewFloat(0).Mul(mul, b)
	}
	return mul
}

// QuoInt ...
func QuoInt(big1, big2 *big.Int) *big.Int {
	if big2.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}
	return big.NewInt(0).Quo(big1, big2)
}

// QuoFloat ...
func QuoFloat(big1, big2 *big.Float) *big.Float {
	if big2.Cmp(big.NewFloat(0)) == 0 {
		return big.NewFloat(0)
	}
	return big.NewFloat(0).Quo(big1, big2)
}

// SumInt ...
func SumInt(bigs ...*big.Int) *big.Int {
	sum := big.NewInt(0)
	for _, b := range bigs {
		sum = big.NewInt(0).Add(sum, b)
	}
	return sum
}

// SumFloat ...
func SumFloat(bigs ...*big.Float) *big.Float {
	sum := big.NewFloat(0)
	for _, b := range bigs {
		sum = big.NewFloat(0).Add(sum, b)
	}
	return sum
}
