package dsarecursion

func PalindromeCheck(s string, left, right int) bool {
	if left >= right {
		return true
	}
	if s[left] != s[right] {
		return false
	}
	return PalindromeCheck(s, left+1, right-1)
}
