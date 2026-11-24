package dsaheap

import "container/heap"

func KthLargetElementinArray(nums []int, k int) int {
	h := &LargestHeap{}
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

type LargestHeap []int

func (h LargestHeap) Len() int {
	return len(h)
}
func (h LargestHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h LargestHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *LargestHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *LargestHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
