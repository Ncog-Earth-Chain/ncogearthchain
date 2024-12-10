package utils

import "math/big"

// ToNec number of NEC to Wei
func ToNec(nec uint64) *big.Int {
	return new(big.Int).Mul(new(big.Int).SetUint64(nec), big.NewInt(1e18))
}
