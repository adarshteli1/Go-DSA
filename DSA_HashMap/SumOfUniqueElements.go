package dsahashmap

func SumOfUniqueElements(nums []int) int {
	m := make(map[int]int)
	for _, val := range nums {
		m[val]++
	}
	var sum int
	for number, Frequency := range m {
		if Frequency < 2 {
			sum += number
		}
	}
	return sum
}
