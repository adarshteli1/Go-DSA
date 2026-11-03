package dsahashmap

func MajorityElement(nums []int) int {
	lead := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			lead = nums[i]
		}
		if nums[i] == lead {
			count++
		} else {
			count--
		}
	}
	return lead
}
