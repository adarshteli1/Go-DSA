package dsastringsarrays

func SquaresOfASortedArray(nums []int) []int {
	start := 0
	end := len(nums) - 1
	res := make([]int, len(nums))
	k := len(res) - 1
	for start <= end {
		if nums[start]*nums[start] > nums[end]*nums[end] {
			res[k] = nums[start] * nums[start]
			start++
			k--
		} else {
			res[k] = nums[end] * nums[end]
			end--
			k--
		}
	}
	return res
}
