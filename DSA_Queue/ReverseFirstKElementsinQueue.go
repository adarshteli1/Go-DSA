package dsaqueue

func ReverseFirstKElementsInQueue(queue []int, k int) []int {
	queue2 := []int{}
	stack := []int{}
	for i := 0; i < k; i++ {
		front := queue[0]
		queue = queue[1:]
		stack = append(stack, front)
	}
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		queue2 = append(queue2, front)
	}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		queue = append(queue, top)
	}
	for len(queue2) > 0 {
		front := queue2[0]
		queue2 = queue2[1:]
		queue = append(queue, front)
	}
	return queue

}
