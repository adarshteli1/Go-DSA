package dsaheap

import "container/heap"

func NumberGame(nums []int) []int {
	h := &MinHeap{}
	heap.Init(h)
	arr := []int{}
	for _, val := range nums {
		heap.Push(h, val)
	}
	for h.Len() > 1 {
		x := heap.Pop(h).(int)
		y := heap.Pop(h).(int)
		arr = append(arr, y)
		arr = append(arr, x)
	}
	return arr
}
