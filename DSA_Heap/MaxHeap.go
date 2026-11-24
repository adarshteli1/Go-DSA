package dsaheap

type Maxheap struct {
	arr []int
}

func (h *Maxheap) Insert(val int) {
	h.arr = append(h.arr, val)
	h.maxheapifyup(len(h.arr) - 1)
}
func (h *Maxheap) maxheapifyup(length int) {
	for h.arr[parent(length)] < h.arr[length] {
		h.swap(parent(length), length)
		length = parent(length)
	}
}

func (h *Maxheap) Extract() int {
	if len(h.arr) == 0 {
		return -1
	}
	extracted := h.arr[0]
	l := len(h.arr) - 1
	h.arr[0] = h.arr[l]
	h.arr = h.arr[:l]
	h.maxheapifydown(0)
	return extracted
}

func (h *Maxheap) maxheapifydown(index int) {
	l, r := left(index), right(index)
	childtocompare := 0
	for l <= len(h.arr)-1 {
		if l == len(h.arr)-1 {
			childtocompare = l
		} else if h.arr[l] > h.arr[r] {
			childtocompare = l
		} else {
			childtocompare = r
		}

		if h.arr[index] < h.arr[childtocompare] {
			h.swap(index, childtocompare)
			index = childtocompare
			l = left(index)
			r = right(index)
		} else {
			return
		}
	}
}

func (h *Maxheap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}
