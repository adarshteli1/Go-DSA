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

/*
LC Q. 622 – Design Circular Queue

Problem Statement:

Design your implementation of the circular queue. The circular queue is a linear data structure in which the operations are performed based on FIFO (First In First Out) principle, and the last position is connected back to the first position to make a circle. It is also called "Ring Buffer".

One of the benefits of the circular queue is that we can make use of the spaces in front of the queue. In a normal queue, once the queue becomes full, we cannot insert the next element even if there is a space in front of the queue. But using the circular queue, we can use the space to store new values.

Examples:

Example 1:

Input:
["MyCircularQueue", "enQueue", "enQueue", "enQueue", "enQueue",
 "Rear", "isFull", "deQueue", "enQueue", "Rear"]

[[3], [1], [2], [3], [4], [], [], [], [4], []]

Output:
[null, true, true, true, false, 3, true, true, true, 4]

Example 2:

Input:
["MyCircularQueue", "enQueue", "Rear", "deQueue", "Front",
 "deQueue", "Front", "deQueue", "isEmpty", "deQueue",
 "enQueue", "Front", "Rear"]

[[2], [1], [], [], [], [], [], [], [], [], [3], [], []]

Output:
[null, true, 1, true, 2, true, -1, false, true, false, true, 3, 3]
*/
