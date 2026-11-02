package dsastringsarrays

func ValidAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var m [26]int
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		if m[t[i]-'a'] == 0 {
			return false
		}
		m[t[i]-'a']--
	}
	return true
}
