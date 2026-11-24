package dsaheap

import "container/heap"

func KthLargetElementinArray(nums []int, k int) int {
	h := &MaxHeap{}
	heap.Init(h)
	for _, val := range nums {
		heap.Push(h, val)
	}
	max := -1
	for k > 0 {
		x := heap.Pop(h).(int)
		max = x
		k--
	}
	return max
}
