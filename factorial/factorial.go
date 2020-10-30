package factorial

import "math/big"

// The simplest method for calculating the factorial.
// Uses the int type, so it is only applicable for numbers up to 20.
func GetFactorial(n int) int {
	if n < 1 {
		return 1
	}

	var nFactorial int = 1

	for i := 1; i <= n; i++ {
		nFactorial *= i
	}

	return nFactorial
}

// The simplest method for calculating the factorial.
// Uses the big int type, so it can be applied for numbers greater than 20.
func GetBigFactorial(n int) *big.Int {
	var nFactorial = big.NewInt(int64(1))

	if n < 1 {
		return nFactorial
	}

	for i := 1; i <= n; i++ {
		bigI := big.NewInt(int64(i))
		nFactorial.Mul(nFactorial, bigI)
	}

	return nFactorial
}
