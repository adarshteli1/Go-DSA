package dsaqueue

func InterLeaveFirstHalfWithSecondHalf(queue []int) []int {
	queue1 := []int{}
	result := []int{}
	n := len(queue) / 2
	for i := 0; i < n; i++ {
		front := queue[0]
		queue = queue[1:]
		queue1 = append(queue1, front)
	}
	for len(queue) > 0 && len(queue1) > 0 {
		front := queue1[0]
		queue1 = queue1[1:]
		result = append(result, front)
		front2 := queue[0]
		queue = queue[1:]
		result = append(result, front2)

	}
	return result
}
