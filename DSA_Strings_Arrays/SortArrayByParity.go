package dsastringsarrays

func SortArrayByParity(nums []int) []int {
	evenpointer := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[evenpointer], nums[i] = nums[i], nums[evenpointer]
			evenpointer++
		}
	}
	return nums
}
