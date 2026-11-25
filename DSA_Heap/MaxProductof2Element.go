
package dsaheap
import "container/heap"

type largestHeap []int

func MaxProductof2





func (h largestHeap) Len() int {
	return len(h)
}
func (h largestHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h largestHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *largestHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *largestHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
