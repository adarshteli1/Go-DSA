package dsarecursion

func ReverseArray(nums []int, s, e int) []int {
	if s >= e {
		return nums
	}
	nums[s], nums[e] = nums[e], nums[s]
	return ReverseArray(nums, s+1, e-1)
}
