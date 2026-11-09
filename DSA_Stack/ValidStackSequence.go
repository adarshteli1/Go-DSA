package dsastack

func ValidStackSequence(pushed, popped []int) bool {
	j := 0
	stack := []int{}
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for j < len(popped) && len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}
	return j == len(popped)
}
