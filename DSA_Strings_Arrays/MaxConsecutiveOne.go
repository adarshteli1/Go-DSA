package dsastringsarrays

func MaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			count--
		}
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}
