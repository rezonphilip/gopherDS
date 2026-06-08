package recursion

func Sum(n int) int {
	// base condition
	if n == 1 {
		return 1
	}

	// recursive condition
	return n + Sum(n-1)
}
