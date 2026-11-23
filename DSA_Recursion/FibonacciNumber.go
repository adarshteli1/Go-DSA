package dsarecursion

func FibonacciNumber(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return FibonacciNumber(n-1) + FibonacciNumber(n-2)
}

/*

func FibonacciNumber(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciNumber(n-1) + FibonacciNumber(n-2)
}

*/
