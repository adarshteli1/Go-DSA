package dsahashmap

func IntersectionOfTwoArrays(nums1, nums2 []int) []int {
	res := []int{}
	m := make(map[int]bool)
	for _, val := range nums1 {
		m[val] = true
	}
	for _, val := range nums2 {
		if m[val] {
			res = append(res, val)
			delete(m, val)
		}
	}
	return res
}
