package dsahashmap

func RansomNote(ransomeNote string, magazine string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		m[magazine[i]]++
	}
	for i := 0; i < len(ransomeNote); i++ {
		if m[ransomeNote[i]] == 0 {
			return false
		}
		m[ransomeNote[i]]--
	}
	return true
}
