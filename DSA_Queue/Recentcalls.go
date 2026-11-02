package dsaqueue

func RecentCalls(times []int, t int) int {
	times = append(times, t)
	for len(times) > 0 && times[0] < t-3000 {
		times = times[1:]
	}
	return len(times)
}
