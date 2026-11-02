package dsastack

func NextGreater(nums1, nums2 []int) []int {
	next := make(map[int]int)
	stack := []int{}
	for i := 0; i < len(nums2); i++ {
		cur := nums2[i]
		for len(stack) > 0 && cur > stack[len(stack)-1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			next[top] = cur
		}
		stack = append(stack, cur)
	}
	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		val, exists := next[nums1[i]]
		if exists {
			res[i] = val
		} else {
			res[i] = -1
		}
	}
	return res

}
