package dsarecursion

//Count Digits

func CountDigits(a int) int {
	if a < 10 {
		return a
	}
	return 1 + CountDigits(a/10)
}

//Sum  of Digits

func SumOfDigits(a int) int {
	if a < 10 {
		return a
	}
	return (a % 10) + SumOfDigits(a/10)
}
