package dsastringsarrays

func MoveZeros(nums []int) []int {
	Write := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[Write], nums[i] = nums[i], nums[Write]
			Write++
		}
	}
	return nums
}
