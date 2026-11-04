package dsahashmap

func SubArraySumEqualsK(nums []int, k int) int {
	m := make(map[int]int)
	m[0] = 1
	sum := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		count += m[sum-k]
		m[sum]++
	}
	return count
}
