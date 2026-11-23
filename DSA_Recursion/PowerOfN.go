package dsarecursion

func MaximumInArray(arr []int, max, index int) int {
	if index == len(arr) {
		return max
	}
	if arr[index] > max {
		max = arr[index]
	}
	return MaximumInArray(arr, max, index+1)
}
