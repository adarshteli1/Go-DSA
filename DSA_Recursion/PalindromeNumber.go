package dsarecursion

func PalindromeNumber(number int) bool {
	return PalindromeNumberHelper(number, number, 0)
}

func PalindromeNumberHelper(number, original, reversed int) bool {
	if number == 0 {
		return original == reversed
	}
	reversed = reversed*10 + number%10
	return PalindromeNumberHelper(number/10, original, reversed)
}
