package dsaheap

import "container/heap"

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}
func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func LastWeightStone(stones []int) int {
	m := &maxHeap{}
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
