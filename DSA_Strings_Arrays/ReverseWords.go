package dsastringsarrays

import "strings"

func ReverseWordsInString(a string) string {
	s := strings.Fields(a)
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return strings.Join(s, " ")

}
