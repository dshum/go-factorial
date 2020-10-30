package main

import (
	"fmt"

	"github.com/dshum/go-factorial/factorial"
)

func main() {
	number := 10
	result := factorial.GetFactorial(number)

	fmt.Printf("Factorial of %d is %d\n", number, result)

	number = 100
	result = factorial.GetFactorial(number)

	fmt.Printf("Factorial of %d is %d\n", number, result)
}
