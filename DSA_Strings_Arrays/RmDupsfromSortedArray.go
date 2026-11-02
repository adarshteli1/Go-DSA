package dsastringsarrays

func RemoveDuplicatesFromSortedArray(nums []int) []int {
	swap := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] { //swap
			nums[swap], nums[i] = nums[i], nums[swap]
			swap++
		}
	}
	return nums
}
