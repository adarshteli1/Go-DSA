package dsaqueue

func StudentsUnableToEat(students []int, sanwhiches []int) int {
	queue := students
	failure := 0
	i := 0
	for len(queue) > 0 && failure < len(queue) {
		if queue[0] == sanwhiches[i] {
			queue = queue[1:]
			i++
			failure = 0
		} else {
			student := queue[0]
			queue = queue[1:]
			queue = append(queue, student)
			failure++
		}
	}
	return len(queue)
}
