package factorial

func Factorial(n int) int {

	if n == 1 {
		return n
	} else {
		return n * Factorial(n-1)
	}

}
