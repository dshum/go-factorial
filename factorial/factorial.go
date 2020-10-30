package factorial

func GetFactorial(n int) int {
	if n < 1 {
		return 1
	}

	var n_factorial int = 1

	for i := 1; i <= n; i++ {
		n_factorial = n_factorial * i
	}

	return n_factorial
}
