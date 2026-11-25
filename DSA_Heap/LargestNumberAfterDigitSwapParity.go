package dsaheap

import "container/heap"

func LargestNumberAfterDigitsSwapByParity(n int) int {
	even := &MaxHeap{}
	odd := &MaxHeap{}
	heap.Init(even)
	heap.Init(odd)
	ans := 0
	digit := []int{}
	for i := n; i > 0; i /= 10 {
		d := i % 10
		digit = append(digit, d)
		if d%2 == 0 {
			heap.Push(even, d)
		} else {
			heap.Push(odd, d)
		}
	}
	for i := len(digit) - 1; i >= 0; i-- {
		if digit[i]%2 == 0 {
			x := heap.Pop(even).(int)
			ans = ans*10 + x
		} else {
			y := heap.Pop(odd).(int)
			ans = ans*10 + y
		}
	}
	return ans
}
