package dsarecursion

//Count Digits

func CountDigits(a int) int {
	if a < 10 {
		return a
	}
	return 1 + CountDigits(a/10)
}

//Sum  of Digits

func SumOfDigits(a int, sum int) int {
	if a < 10 {
		return sum + a
	}
	sum += a % 10
	return SumOfDigits(a/10, sum)
}
