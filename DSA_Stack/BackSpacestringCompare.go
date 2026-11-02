package dsastack

func BackSpaceStringCompare(s, t string) bool {
	stack1 := []rune{}
	stack2 := []rune{}
	for _, val := range s {
		if val == '#' {
			if len(stack1) > 0 {
				stack1 = stack1[:len(stack1)-1]
			}
		} else {
			stack1 = append(stack1, (val))
		}
	}
	for _, val := range t {
		if val == '#' {
			if len(stack2) > 0 {
				stack2 = stack2[:len(stack2)-1]
			}
		} else {
			stack2 = append(stack2, (val))
		}
	}
	if len(stack1) != len(stack2) {
		return false
	}
	for i := range stack1 {
		if stack1[i] != stack2[i] {
			return false
		}
	}

	return true
}
