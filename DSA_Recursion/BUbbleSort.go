package dsarecursion

func BubbleSort(nums []int) []int {
	return bubblesort(nums, 0, len(nums)-1)
}

func bubblesort(nums []int, index int, size int) []int {
	if size <= 0 {
		return nums
	}
	if index < size {
		if nums[index] > nums[index+1] {
			nums[index], nums[index+1] = nums[index+1], nums[index]
		}
		return bubblesort(nums, index+1, size)
	}
	return bubblesort(nums, 0, size-1)
}
