package dsaheap

import (
	"container/heap"
	"math"
)

func TakeGiftFromTheRichestPile(nums []int,k int)int64{
	h:=&MaxHeap{}
	heap.Init(h)
	for _,val:=range nums{
		heap.Push(h,val)
	}
	for i:=0;i<k;i++{
		x:=heap.Pop(h).(int)
		sqr:=int(math.Sqrt(float64(x)))
		heap.Push(h,sqr)
	}
	var sum int64
	for h.Len()>0{
		x:=heap.Pop(h).(int)
		sum+=int64(x)
	}
	return sum
}