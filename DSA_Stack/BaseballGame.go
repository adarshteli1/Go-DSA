package dsastack

import "strconv"

func BaseBallGame(nums []string) int {
	stack := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] == "C" {
			stack = stack[:len(stack)-1]
		} else if nums[i] == "D" {
			top := stack[len(stack)-1]
			stack = append(stack, 2*top)
		} else if nums[i] == "+" {
			n := len(stack)
			add := stack[n-1] + stack[n-2]
			stack = append(stack, add)

		} else {
			num, _ := strconv.Atoi(nums[i])
			stack = append(stack, num)
		}
	}
	total := 0
	for _, val := range stack {
		total += val
	}
	return total
}
