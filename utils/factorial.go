package utils

import (
	"fmt"
	"math/big"
)

// factorial function
// @return *big.Int
// @param num int

func Factorial(num int) *big.Int {
	if num < 0 {
		fmt.Print("Factorial of a negative number is undefined.")
		return big.NewInt(-1)
	}

	result := big.NewInt(1)

	for i := 2; i <= num; i++ {
		result.Mul(result, big.NewInt(int64(i)))
	}

	return result
}
