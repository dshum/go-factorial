package main

import (
	"fmt"
	"math/big"

	"github.com/dshum/go-factorial/factorial"
)

func main() {
	// Calculate 10! using int type
	number := 10
	result0 := factorial.GetFactorial(number)

	fmt.Printf("Factorial of %d is %d\n", number, result0)

	// Trying to calculate 100! using int type.
	// The received result is out of range.
	number = 100
	result1 := factorial.GetFactorial(number)

	// Calculate the factorial of numbers from 1 to 100, using big int type.
	fmt.Printf("Factorial of %d is %d\n", number, result1)

	var result2 *big.Int

	for i := 1; i <= 100; i++ {
		result2 = factorial.GetBigFactorial(i)
		fmt.Printf("Factorial of %v is %v\n", i, result2)
	}
}
