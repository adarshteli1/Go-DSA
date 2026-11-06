package dsastringsarrays

func NumbersWithEvenDigits(nums []int) int {
	count := 0
	for _, num := range nums {
		digit := 0
		for num > 0 {
			num /= 10
			digit++
		}
		if digit%2 == 0 {
			count++
		}
	}
	return count
}
