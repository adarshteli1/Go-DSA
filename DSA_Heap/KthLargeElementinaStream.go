package dsaheap

import "container/heap"

type Kth struct {
	k    int
	heap *MinHeap
}

func KthLargetElementinStream(nums []int, k int) Kth {
	h := &MinHeap{}
	heap.Init(h)
	obj := Kth{k: k, heap: h}
	for _, val := range nums {
		obj.Add(val)
	}
	return obj
}

func (obj *Kth) Add(val int) int {
	heap.Push(obj.heap, val)
	for obj.heap.Len() > obj.k {
		heap.Pop(obj.heap)
	}
	return (*obj.heap)[0]
}
