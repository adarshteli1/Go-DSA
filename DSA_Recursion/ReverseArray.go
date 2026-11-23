package dsarecursion

func ReverseArray(nums []int) []int {
	return ReverseArrayHelper(nums, 0, len(nums)-1)
}

func ReverseArrayHelper(nums []int, start, end int) []int {
	if start >= end {
		return nums
	}
	nums[start], nums[end] = nums[end], nums[start]
	return ReverseArrayHelper(nums, start+1, end-1)
}
