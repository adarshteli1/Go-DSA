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

/*
LC Q. 735 – Asteroid Collision

Problem Statement:

We are given an array asteroids of integers representing asteroids in a row.

For each asteroid, the absolute value represents its size, and the sign
represents its direction (positive meaning right, negative meaning left).
Each asteroid moves at the same speed.

Find out the state of the asteroids after all collisions.

If two asteroids meet, the smaller one will explode. If both are the
same size, both will explode. Two asteroids moving in the same direction
will never meet.

Example 1:

Input:
asteroids = [5, 10, -5]

Output:
[5, 10]

Explanation:
The 10 and -5 collide, resulting in 10. The 5 and 10 never collide.

Example 2:

Input:
asteroids = [8, -8]

Output:
[]

Explanation:
The 8 and -8 collide, exploding each other.

Example 3:

Input:
asteroids = [10, 2, -5]

Output:
[10]

Explanation:
The 2 and -5 collide, resulting in -5.
The 10 and -5 collide, resulting in 10.

Example 4:

Input:
asteroids = [3, 5, -6, 2, -1, 4]

Output:
[-6, 2, 4]
*/
