package dsaheap

import "container/heap"

func MaxProductof2ElementsinArray(nums []int) int {
	h := &MaxHeap{}
	heap.Init(h)
	for _, val := range nums {
		heap.Push(h, val)
	}
	x := heap.Pop(h).(int)
	y := heap.Pop(h).(int)
	return (x - 1) * (y - 1)
}
