package utils

import "math/big"

func ToWei18(num int64) *big.Int {
	n, _ := new(big.Int).SetString("1000000000000000000", 10)
	m := new(big.Int)
	m.SetInt64(num)
	m.Mul(n, m)
	return m

}
