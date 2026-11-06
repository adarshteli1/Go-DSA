package dsastack

func AsteroidCollision(nums []int) []int {
	stack := []int{}
	for i := 0; i < len(nums); i++ {
		current := nums[i]
		destroyed := false
		for len(stack) > 0 && stack[len(stack)-1] > 0 && current < 0 {
			top := stack[len(stack)-1]
			if top == -current {
				stack = stack[:len(stack)-1]
				destroyed = true
				break
			} else if top > -current {
				destroyed = true
				break
			} else {
				stack = stack[:len(stack)-1]
			}
		}
		if !destroyed {
			stack = append(stack, current)
		}
	}
	return stack
}
