package dsarecursion

func LinearSearch(nums []int, target int) bool {
	return linearSearch(nums, target, 0)
}

func linearSearch(nums []int, target int, index int) bool {
	if index >= len(nums) {
		return false
	}
	return nums[index] == target || linearSearch(nums, target, index+1)
}
