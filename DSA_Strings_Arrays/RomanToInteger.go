package dsastringsarrays

func RomanToInteger(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var total int
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && m[s[i]] < m[s[i+1]] {
			total -= m[s[i]]
		} else {
			total += m[s[i]]
		}
	}
	return total
}
