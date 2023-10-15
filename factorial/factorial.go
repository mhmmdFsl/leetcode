package factorial

func factorial(n int) int {
	if n == 2 {
		return 2
	}
	return n * factorial(n-1)
}
