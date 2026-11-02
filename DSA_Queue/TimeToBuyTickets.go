package dsaqueue

func TimeToBuyTickets(tickets []int, k int) int {
	queue := []int{}
	for i := range tickets {
		queue = append(queue, i)
	}
	c := 0
	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]
		tickets[start]--
		c++
		if tickets[start] == 0 && start == k {
			return c
		}

		if tickets[start] > 0 {
			queue = append(queue, start)
		}

	}
	return c
}
