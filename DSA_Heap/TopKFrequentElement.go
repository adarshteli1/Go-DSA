package dsaheap

import "container/heap"

type Pair struct {
	number    int
	frequency int
}
type MaxHeapPair []Pair

func TopKFreqElements(nums []int, k int) []int {
	m := make(map[int]int)
	for _, val := range nums {
		m[val]++
	}
	h := &MaxHeapPair{}
	heap.Init(h)
	for i, val := range m {
		heap.Push(h, Pair{number: i, frequency: val})
	}
	res := []int{}
	for i := 0; i < k && h.Len() > 0; i++ {
		x := heap.Pop(h).(Pair)
		res = append(res, x.number)
	}
	return res
}

func (h MaxHeapPair) Len() int { return len(h) }
func (h MaxHeapPair) Less(i, j int) bool {
	return h[i].frequency > h[j].frequency
}
func (h MaxHeapPair) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MaxHeapPair) Push(val any) {
	*h = append(*h, val.(Pair))
}
func (h *MaxHeapPair) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
