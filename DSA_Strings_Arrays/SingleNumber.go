package dsastringsarrays

func SingleNumber(nums []int) int {
	var ans int
	for _, val := range nums {
		ans ^= val
	}
	return ans
}
