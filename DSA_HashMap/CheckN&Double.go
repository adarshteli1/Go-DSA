package dsahashmap

func CheckNandItsDoubleExists(nums []int) bool {
	m := make(map[int]bool)
	for _, val := range nums {
		if m[val] || (val%2 == 0 && m[val/2]) {
			return true
		}
		m[val] = true
	}
	return false
}
