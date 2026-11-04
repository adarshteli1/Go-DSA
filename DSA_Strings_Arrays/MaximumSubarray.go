package dsastringsarrays

func MaximumSubaaray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	count := 0
	maxcount := nums[0]
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > maxcount {
			maxcount = count
		}
		if count < 0 {
			count = 0
		}
	}
	return maxcount
}
