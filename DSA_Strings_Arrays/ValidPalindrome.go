package dsastringsarrays

import "unicode"

func Validpalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left <= right {
		for left <= right && !IsAlphaNumberic(rune(s[left])) {
			left++
		}
		for left <= right && !IsAlphaNumberic(rune(s[right])) {
			right--
		}
		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func IsAlphaNumberic(a rune) bool {
	return unicode.IsDigit(a) || unicode.IsLetter(a)
}
