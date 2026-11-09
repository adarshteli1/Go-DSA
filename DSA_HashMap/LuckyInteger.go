package dsahashmap

func LuckyInteger(nums []int) int {
	max := -1
	m := make(map[int]int)
	for _, val := range nums {
		m[val]++
	}
	for number, frequency := range m {
		if number == frequency {
			if number > max {
				max = number
			}
		}
	}
	return max
}
