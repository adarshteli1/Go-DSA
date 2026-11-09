package dsastringsarrays

func ContainerWithMostWater(heights []int) int {
	left := 0
	right := len(heights) - 1
	maxArea := 0
	for left < right {
		width := right - left
		length := min(heights[right], heights[left])
		area := width * length
		if area > maxArea {
			maxArea = area
		}
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
