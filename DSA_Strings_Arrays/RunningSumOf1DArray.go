package dsastringsarrays

func RunningSumOf1DArray(nums []int) []int {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		nums[i] = sum
	}
	return nums
}
