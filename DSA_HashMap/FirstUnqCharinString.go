package dsahashmap

func FirstUniqueCharacterInAString(s string) int {
	var Letters [26]int
	for i := 0; i < len(s); i++ {
		Letters[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if Letters[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
