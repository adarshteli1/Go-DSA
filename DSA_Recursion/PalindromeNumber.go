package dsarecursion

func PalindromeNumber(number, original, reversed int) bool {
	if number == 0 {
		return original == reversed
	}
	reversed = reversed*10 + number%10
	return PalindromeNumber(number/10, original, reversed)
}
