package dsastack

import "strconv"

func EvaluateReversePolishNotation(nums []string) int {
	stack := []int{}
	for _, val := range nums {
		switch val {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return -1
			}
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			var result int
			switch val {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				// if b==0{return -1}
				result = a / b
			}
			stack = append(stack, result)
		default:
			num, _ := strconv.Atoi(val)
			stack = append(stack, num)
		}
	}
	return stack[len(stack)-1]

}
