package recursion

func Reverse(str string) string {
	// base condition
	if len(str) == 1 {
		return str
	}

	// recursive condition
	return Reverse(str[1:]) + str[:1]
}
