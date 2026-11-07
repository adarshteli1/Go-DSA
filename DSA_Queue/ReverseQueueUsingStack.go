package dsaqueue

func ReverseQueueUsingStack(queue []int) []int {
	stack := []int{}
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		stack = append(stack, front)
	}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		queue = append(queue, top)
	}

	return queue
}
