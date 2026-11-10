	package dsahashmap

	func LongestSubstringWithoutRepeatingCharacters(s string) int {
		left := 0
		m := make(map[byte]bool)
		maxlen := 0
		for right := 0; right < len(s); right++ {
			for m[s[right]] {
				delete(m, s[left])
				left++
			}
			m[s[right]] = true
			if right-left+1 > maxlen {
				maxlen = right - left + 1
			}
		}
		return maxlen
	}
