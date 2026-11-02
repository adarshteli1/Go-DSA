package dsastringsarrays

func FindNumbersDisappearedInArray(nums []int) []int {
	res := []int{}
	for _, val := range nums {
		if val < 0 {
			val = -val
		}
		index := val - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}
