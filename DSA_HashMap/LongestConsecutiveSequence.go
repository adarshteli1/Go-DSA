package dsahashmap

func LongestConsecutiveSequence(nums []int) int {
	var maxcount int
	m := make(map[int]bool)
	for _, value := range nums {
		m[value] = true
	}
	for i := 0; i < len(nums); i++ {
		if !m[nums[i]-1] {
			start := nums[i]
			count := 1
			for m[start+1] {
				start++
				count++
			}
			if count > maxcount {
				maxcount = count
			}
		}
	}

	return maxcount
}
