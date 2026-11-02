package dsahashmap

import "strings"

func WordPattern(s string, pattern string) bool {
	m1 := make(map[byte]string)
	m2 := make(map[string]byte)
	p := strings.Fields(s)
	if len(pattern) != len(p) {
		return false
	}
	for i := 0; i < len(s); i++ {
		if value, exists := m1[s[i]]; exists {
			if value != p[i] {
				return false
			}
		} else {
			m1[s[i]] = p[i]
		}
		if value2, exists2 := m2[p[i]]; exists2 {
			if value2 != s[i] {
				return false
			}
		} else {
			m2[p[i]] = s[i]
		}
	}
	return true
}
