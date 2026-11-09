package dsahashmap

func SetMisMatch(nums []int) []int {
	var duplicate int
	var missing int
	m := make(map[int]bool)
	for _, value := range nums {
		if m[value] {
			duplicate = value
		}
		m[value] = true
	}
	for i := 1; i <= len(nums); i++ {
		if !m[i] {
			missing = i
			break
		}
	}
	return []int{duplicate, missing}
}
