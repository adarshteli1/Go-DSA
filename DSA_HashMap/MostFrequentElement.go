package dsahashmap

func MostFrequentElement(nums []int) int {
	m := make(map[int]int)
	for _, value := range nums {
		m[value]++
	}
	smallest := -1
	maxFrequency := 0
	for number, frequency := range m {
		if number%2 == 0 {
			if frequency > maxFrequency {
				smallest = number
				maxFrequency = frequency
			} else if frequency == maxFrequency && (smallest == -1 || number < smallest) {
				smallest = number
			}
		}
	}
	return smallest
}
