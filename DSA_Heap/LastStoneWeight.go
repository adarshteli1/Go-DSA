package dsaheap

import "container/heap"

func LastWeightStone(stones []int) int {
	m := &MaxHeap{}
	heap.Init(m)
	for _, val := range stones {
		heap.Push(m, val)
	}
	for m.Len() > 1 {
		x := heap.Pop(m).(int)
		y := heap.Pop(m).(int)
		if x != y {
			heap.Push(m, x-y)
		}
	}
	if m.Len() == 0 {
		return 0
	}
	return heap.Pop(m).(int)

}
