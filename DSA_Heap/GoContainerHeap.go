package dsaheap

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j] // for MinHeap use return h[i]<h[j]
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MaxHeap) Push(val any) {
	*h = append(*h, val.(int))
}
func (h *MaxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func main() {
	h := &MaxHeap{}

	heap.Init(h)

	heap.Push(h, 10)
	heap.Push(h, 30)
	heap.Push(h, 20)

	fmt.Println((*h)) 

	max := heap.Pop(h).(int)
	fmt.Println(max) // 30
}
