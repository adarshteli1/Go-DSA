package dsarecursion

func ReverseNumber(number, reversed int) int {
	if number == 0 {
		return reversed
	}
	reversed = reversed*10 + number%10
	return ReverseNumber(number/10, reversed)
}
