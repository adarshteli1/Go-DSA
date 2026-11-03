package dsastringsarrays

func RemoveElement(nums []int, value int) int {
	Swap := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != value {
			nums[Swap], nums[i] = nums[i], nums[Swap]
			Swap++
		}
	}
	return Swap
}
