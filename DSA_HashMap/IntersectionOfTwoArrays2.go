package dsahashmap

func IntersectionOfTwoArrays2(nums1, nums2 []int) []int {
	m := make(map[int]int)
	res := []int{}
	for _, val := range nums1 {
		m[val]++
	}
	for _, val := range nums2 {
		if count, exists := m[val]; exists && count > 0 {
			res = append(res, val)
			m[val]--
		}
	}
	return res
}
