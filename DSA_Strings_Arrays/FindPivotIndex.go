package dsastringsarrays

func PivotIndex(nums []int) int {
	var total int
	for _, val := range nums {
		total += val
	}
	leftsum := 0
	for i := 0; i < len(nums); i++ {
		rightsum := total - leftsum - nums[i]
		if rightsum == leftsum {
			return i
		} else {
			leftsum += nums[i]
		}
	}
	return -1
}
