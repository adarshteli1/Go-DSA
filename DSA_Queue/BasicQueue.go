package dsaqueue

type MyCircularQueue struct {
	Data  []int
	Front int
	Rear  int
	Size  int
	Cap   int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		Data:  make([]int, k),
		Front: 0,
		Rear:  -1,
		Size:  0,
		Cap:   k,
	}
}

func (q *MyCircularQueue) Enqueue(value int) bool {
	if q.Size == q.Cap {
		return false
	}
	q.Rear = (q.Rear + 1) % q.Cap
	q.Data[q.Rear] = value
	q.Size++
	return true
}

func (q *MyCircularQueue) Dequeue() bool {
	if q.Size == 0 {
		return false
	}
	q.Front = (q.Front + 1) % q.Cap
	q.Size--
	return true
}

func (q *MyCircularQueue) FrontValue() int {
	if q.Size == 0 {
		return -1
	}
	return q.Data[q.Front]
}

func (q *MyCircularQueue) RearValue() int {
	if q.Size == 0 {
		return -1
	}
	return q.Data[q.Rear]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return q.Size == 0
}

func (q *MyCircularQueue) IsFull() bool {
	return q.Cap == q.Size
}
