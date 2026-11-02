package dsastringsarrays

func ReverseVowels(Input string) string {
	s := []byte(Input)
	left := 0
	right := len(s) - 1
	for left < right {
		for left < right && !IsVowel(s[left]) {
			left++
		}
		for left < right && !IsVowel(s[right]) {
			right--
		}
		s[right], s[left] = s[left], s[right]
		left++
		right--
	}
	return string(s)
}

func IsVowel(a byte) bool {
	switch a {
	case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
		return true
	default:
		return false
	}

}
