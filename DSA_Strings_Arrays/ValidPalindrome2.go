package dsastringsarrays

func ValidPalindromeII(s string) bool {
	start := 0
	end := len(s) - 1
	for start < end {
		if s[start] != s[end] {
			return IsPalindrome(s, start+1, end) || IsPalindrome(s, start, end-1)
		}
		start++
		end--
	}
	return true
}

func IsPalindrome(s string, start int, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
