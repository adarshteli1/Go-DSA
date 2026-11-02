package dsahashmap

func IsomorphicStrings(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	map1 := make(map[byte]byte)
	map2 := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if value1, exists1 := map1[s[i]]; exists1 {
			if value1 != t[i] {
				return false
			}
		} else {
			map1[s[i]] = t[i]
		}
		if value2, exists2 := map2[t[i]]; exists2 {
			if value2 != s[i] {
				return false
			}
		} else {
			map2[t[i]] = s[i]
		}
	}
	return true
}
