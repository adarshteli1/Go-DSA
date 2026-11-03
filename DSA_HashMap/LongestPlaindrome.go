package dsahashmap

func LongestPalindrome(s string) int {
	m := make(map[byte]int)
	hasOddFreq := false
	var count int
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for _, val := range m {
		if val%2 == 0 {
			count += val
		} else {
			count += val - 1
			hasOddFreq = true
		}
	}
	if hasOddFreq {
		count++
	}
	return count
}
