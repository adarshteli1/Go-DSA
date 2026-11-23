package dsarecursion

func ReverseString(s []rune, l, r int) string {
	if l >= r {
		return string(s)
	}

	s[l], s[r] = s[r], s[l]

	return ReverseString(s, l+1, r-1)
}
