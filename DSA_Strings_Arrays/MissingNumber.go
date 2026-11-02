package dsastringsarrays

func MissingNumber(nums []int) int {
	ans := len(nums)
	for i, val := range nums {
		ans ^= i ^ val
	}
	return ans
}
