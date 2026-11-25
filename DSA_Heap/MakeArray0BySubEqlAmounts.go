package dsaheap

import "container/heap"

func MakeArray0BySubtractingEqualAmounts(nums []int) int {
	h := &MinHeap{}
	heap.Init(h)
	for _, val := range nums {
		if val > 0 {
			heap.Push(h, val)
		}
	}
	if h.Len() == 0 {
		return 0
	}
	x := heap.Pop(h).(int)
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i] = nums[i] - x
		}
	}
	return 1 + MakeArray0BySubtractingEqualAmounts(nums)
}
