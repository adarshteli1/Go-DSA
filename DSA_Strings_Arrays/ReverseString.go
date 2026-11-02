package dsastringsarrays

func Reverse(s []string) []string {
	left := 0
	right := len(s) - 1
	for left <= right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return s
}
