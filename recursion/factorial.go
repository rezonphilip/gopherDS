package recursion

func Factorial(n int) int {

	// Base case
	if n <= 1 {
		return 1
	}

	// Recursive case
	return n * Factorial(n-1)
}
