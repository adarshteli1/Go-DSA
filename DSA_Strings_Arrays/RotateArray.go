package dsastringsarrays

func RotateArray(nums []int, k int) []int {
	n := len(nums) - 1
	if n == 0 {
		return nums
	}
	k = k % n

	Rotate(nums, 0, n)
	Rotate(nums, 0, k-1)
	Rotate(nums, k, n)

	return nums
}

func Rotate(nums []int, start, end int) []int {
	for start <= end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
	return nums
}
