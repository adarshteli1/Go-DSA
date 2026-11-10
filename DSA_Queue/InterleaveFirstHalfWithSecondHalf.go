package dsaqueue

func InterLeaveFirstHalfWithSecondHalf(queue []int) []int {
	queue1 := []int{}
	result := []int{}
	n := len(queue) / 2
	for i := 0; i < n; i++ {
		front := queue[0]
		queue = queue[1:]
		queue1 = append(queue1, front)
	}
	for len(queue) > 0 && len(queue1) > 0 {
		front := queue1[0]
		queue1 = queue1[1:]
		result = append(result, front)
		front2 := queue[0]
		queue = queue[1:]
		result = append(result, front2)

	}
	return result
}

/*
GFG – Interleave the First Half of the Queue with Second Half

Note:
This is a GeeksforGeeks problem, not an exact LeetCode problem.
There is no direct equivalent LeetCode problem.

Problem Statement:

Given a queue of even size, rearrange the elements by interleaving the
first half of the queue with the second half.

The first element of the first half should be followed by the first
element of the second half, then the second element of the first half,
followed by the second element of the second half, and so on.

Examples:

Example 1:

Input:
N = 4
Q = [1, 2, 3, 4]

Output:
[1, 3, 2, 4]

Explanation:
First half:  [1, 2]
Second half: [3, 4]

Interleaving both halves gives [1, 3, 2, 4].

Example 2:

Input:
N = 6
Q = [1, 2, 3, 4, 5, 6]

Output:
[1, 4, 2, 5, 3, 6]

Explanation:
First half:  [1, 2, 3]
Second half: [4, 5, 6]

Interleaving both halves gives [1, 4, 2, 5, 3, 6].
*/
