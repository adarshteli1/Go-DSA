package dsaqueue

type Queue struct {
	Data  []int
	Front int
	Rear  int
	Size  int
	Cap   int
}

func Constructor2(k int) Queue {
	return Queue{
		Data:  make([]int, k),
		Front: 0,
		Rear:  -1,
		Size:  0,
		Cap:   k,
	}
}

func (q *Queue) Enqueue(value int) bool {
	if q.Rear == q.Cap-1 {
		return false
	}
	q.Rear++
	q.Data[q.Rear] = value
	q.Size++
	return true
}

func (q *Queue) Dequeue() bool {
	if q.Size == 0 {
		return false
	}
	q.Size--
	q.Front++
	return true
}

func (q *Queue) FrontValue() int {
	if q.Size == 0 {
		return -1
	}
	return q.Data[q.Front]
}

func (q *Queue) RearValue() int {
	if q.Size == 0 {
		return -1
	}
	return q.Data[q.Rear]
}

func (q *Queue) IsEmpty() bool {
	return q.Size == 0
}
func (q *Queue) IsFull() bool {
	return q.Size == q.Cap
}
