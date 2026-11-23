package dsarecursion

func SortedArray(nums []int) bool {
	return sortedarrayhelper(nums, 0)
}

func sortedarrayhelper(nums []int, index int) bool {
	if index == len(nums)-1 {
		return true
	}
	if nums[index] > nums[index+1] {
		return false
	}
	return sortedarrayhelper(nums, index+1)
}
