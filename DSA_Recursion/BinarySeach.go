package dsarecursion

func BinarySearch(nums []int, target int) int {
	return binarysearch(nums, target, 0, len(nums)-1)
}

func binarysearch(nums []int, target, start, end int) int {
	if start > end {
		return -1
	}
	mid := start + (end-start)/2
	if nums[mid] == target {
		return mid
	}
	if nums[mid] < target {
		return binarysearch(nums, target, mid+1, end)
	}
	return binarysearch(nums, target, start, mid-1)
}
