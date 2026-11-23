package dsarecursion

func ReverseNumber(number int) int {
	return reverseNumberHelper(number, 0)
}

func reverseNumberHelper(number, reversed int) int {
	if number == 0 {
		return reversed
	}
	reversed = reversed*10 + number%10
	return reverseNumberHelper(number/10, reversed)
}
