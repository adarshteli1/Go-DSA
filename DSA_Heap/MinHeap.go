package dsaheap

type MinHeap struct {
	arr []int
}

func (h *MinHeap) Insert(val int) {
	h.arr = append(h.arr, val)
	h.minHeapifyup(len(h.arr) - 1)
}

func (h *MinHeap) minHeapifyup(length int) {
	for h.arr[parent(length)] > h.arr[length] {
		h.swap(parent(length), length)
		length = parent(length)
	}
}

func (h *MinHeap) Extract() int {
	if len(h.arr) == 0 {
		return -1
	}
	extract := h.arr[0]
	l := len(h.arr) - 1
	h.arr[0] = h.arr[l]
	h.arr = h.arr[:l]
	h.minHeapifydown(0)
	return extract
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
		}else{
			return
		}
	}
}

func (h *MinHeap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}
