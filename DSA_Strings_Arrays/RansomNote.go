package dsastringsarrays

func RansomNote(ransomNote string, magazine string) bool {
	var letters [26]int

	for i := 0; i < len(magazine); i++ {
		letters[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		if letters[ransomNote[i]-'a'] == 0 {
			return false
		}
		letters[ransomNote[i]-'a']--
	}
	return true
}
