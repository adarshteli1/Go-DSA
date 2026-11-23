package dsarecursion

func ReverseString(s string) string {
	runes := []rune(s)
	return revrseStringhelper(runes, 0, len(runes)-1)
}

func revrseStringhelper(s []rune, left, right int) string {
	if left >= right {
		return string(s)
	}

	s[left], s[right] = s[right], s[left]

	return revrseStringhelper(s, left+1, right-1)
}
