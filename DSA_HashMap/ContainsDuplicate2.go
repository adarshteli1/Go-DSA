package dsahashmap

func ContiansDuplicate2(nums []int, k int) bool {
	m := make(map[int]int)
	for i, val := range nums {
		if value, exists := m[val]; exists && i-value <= k {
			return true
		}
		m[val] = i
	}
	return false
}
