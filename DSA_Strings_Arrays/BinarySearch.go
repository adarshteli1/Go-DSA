package dsastringsarrays

func BinarySearch(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start <= end {
		middle := start + (end-start)/2
		if nums[middle] == target {
			return middle
		}
		if nums[middle] < target {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return -1

}
