package mathutils

import (
	"fmt"
	"math/big"
)

func Factorial(n int64) (*big.Int, error) {
	if n < 0 {
		return nil, fmt.Errorf("negative input: %d", n)
	}

	result := big.NewInt(1)
	if n <= 1 {
		return result, nil
	}

	tmp := new(big.Int)
	for i := int64(2); i <= n; i++ {
		tmp.SetInt64(i)
		result.Mul(result, tmp)
	}

	return result, nil
}
