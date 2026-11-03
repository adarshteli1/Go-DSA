package dsastringsarrays

func LengthOfTheLastWord(s string) int {
	count := 0
	index := len(s) - 1
	for index >= 0 && s[index] == ' ' {
		index--
	}
	for index >= 0 && s[index] != ' ' {
		index--
		count++
	}
	return count
}
