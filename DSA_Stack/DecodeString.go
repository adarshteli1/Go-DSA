package dsastack

import "strings"

func DecodeString(s string) string {
	countstack := []int{}
	stringstack := []string{}
	num := 0
	currentstring := ""
	for i := 0; i < len(s); i++ {
		switch {
		case s[i] >= '0' && s[i] <= '9':
			num = num*10 + int(s[i]-'0')
		case s[i] == '[':
			countstack = append(countstack, num)
			num = 0
			stringstack = append(stringstack, currentstring)
			currentstring = ""
		case s[i] == ']':
			count := countstack[len(countstack)-1]
			countstack = countstack[:len(countstack)-1]
			previousstring := stringstack[len(stringstack)-1]
			stringstack = stringstack[:len(stringstack)-1]

			currentstring = previousstring + strings.Repeat(currentstring, count)
		default:
			currentstring += string(s[i])
		}

	}
	return currentstring
}
