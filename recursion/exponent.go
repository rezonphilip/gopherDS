package recursion

func Exponent(base, exp int) int {
	// base condition
	if exp == 0 {
		return 1
	}

	// recursive condition
	return base * Exponent(base, exp-1)
}
