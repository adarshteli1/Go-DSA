package dsarecursion

func SkipCharacter(s string, char string) string {
	return skipcharacter(s, char, "")
}
func skipcharacter(s, char, ans string) string {
	if len(s) == 0 {
		return ans
	}
	if s[0] != char[0] {
		ans += string(s[0])
	}
	return skipcharacter(s[1:], char, ans)
}
