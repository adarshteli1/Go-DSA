package dsaheap

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(val any) {
	*h = append(*h, val.(int))
}

func (h *MaxHeap) Pop() any {
	old := *h
	x := old[0]
	*h = old[1:]
	return x
}
