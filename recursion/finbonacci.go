package recursion

func Fibonacci(n int) int {
	// base condition
	if n <= 1 {
		return n
	}

	// recursive condition
	return Fibonacci(n-1) + Fibonacci(n-2)
}
