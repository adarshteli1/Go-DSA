package dsahashmap

func UniqueNumberOfOccurence(nums []int) bool {
	m := make(map[int]int)
	for _, val := range nums {
		m[val]++
	}
	m2 := make(map[int]bool)
	for _, val := range m {
		if m2[val] {
			return false
		}
		m2[val] = true
	}
	return true
}
