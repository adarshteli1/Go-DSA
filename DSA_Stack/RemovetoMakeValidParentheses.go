package dsastack

func MinRemoveToMakeValidParentheses(s string) string {
	stack := []int{}
	remove := make(map[int]bool)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			stack = append(stack, i)
		case ')':
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				remove[i] = true
			}
		}
	}
	for _, index := range stack {
		remove[index] = true
	}
	result := []byte{}
	for i := 0; i < len(s); i++ {
		if !remove[i] {
			result = append(result, s[i])
		}
	}
	return string(result)
}
