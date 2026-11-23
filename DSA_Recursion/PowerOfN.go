package dsarecursion

func PowerOfN(n, pow int) int {
	if pow == 0 {
		return 1
	}
	return n * PowerOfN(n, pow-1)
}
