package dsarecursion

func PalindromeCheck(s string) bool {
	return PalindromeCheckhelper(s, 0, len(s)-1)
}

func PalindromeCheckhelper(s string, left, right int) bool {
	if left >= right {
		return true
	}
	if s[left] != s[right] {
		return false
	}
	return PalindromeCheckhelper(s, left+1, right-1)
}
