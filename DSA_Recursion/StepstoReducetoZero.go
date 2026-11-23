package dsarecursion

func NumberOfStepsToReducetoZero(n int) int {
	return numberOfStepsToReducetoZerohelper(n, 0)
}

func numberOfStepsToReducetoZerohelper(n, steps int) int {
	if n == 0 {
		return steps
	}
	if n%2 == 0 {
		return numberOfStepsToReducetoZerohelper(n/2, steps+1)
	}
	return numberOfStepsToReducetoZerohelper(n-1, steps+1)
}
