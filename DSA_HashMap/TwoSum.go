package dsahashmap

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, val := range nums {
		if v, e := m[target-val]; e {
			return []int{v, i}
		}
		m[val] = i
	}
	return nil
}
