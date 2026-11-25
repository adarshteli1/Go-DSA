package dsaheap

type MinHeap struct {
	arr []int
}

// Len implements [heap.Interface].
func (h *MinHeap) Len() int {
	panic("unimplemented")
}

// Less implements [heap.Interface].
func (h *MinHeap) Less(i int, j int) bool {
	panic("unimplemented")
}

// Pop implements [heap.Interface].
func (h *MinHeap) Pop() any {
	panic("unimplemented")
}

// Push implements [heap.Interface].
func (h *MinHeap) Push(x any) {
	panic("unimplemented")
}

// Swap implements [heap.Interface].
func (h *MinHeap) Swap(i int, j int) {
	panic("unimplemented")
}

func (h *MinHeap) Insert(val int) {
	h.arr = append(h.arr, val)
	h.minHeapifyup(len(h.arr) - 1)
}

func (h *MinHeap) minHeapifyup(length int) {
	for length > 0 && h.arr[parent(length)] > h.arr[length] {
		h.swap(parent(length), length)
		length = parent(length)
	}
}

func (h *MinHeap) Extract() (int, bool) {
	if len(h.arr) == 0 {
		return 0, false
	}
	extract := h.arr[0]
	l := len(h.arr) - 1
	h.arr[0] = h.arr[l]
	h.arr = h.arr[:l]
	if len(h.arr) > 0 {
		h.minHeapifydown(0)
	}

	return extract, true
}

func (h *MinHeap) minHeapifydown(index int) {
	childtocompare := 0
	l := left(index)
	r := right(index)
	for l <= len(h.arr)-1 {
		if l == len(h.arr)-1 {
			childtocompare = l
		} else if h.arr[l] < h.arr[r] {
			childtocompare = l
		} else {
			childtocompare = r
		}

		if h.arr[index] > h.arr[childtocompare] {
			h.swap(index, childtocompare)
			index = childtocompare
			l = left(index)
			r = right(index)
		} else {
			return
		}
	}
}

func (h *MinHeap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}
