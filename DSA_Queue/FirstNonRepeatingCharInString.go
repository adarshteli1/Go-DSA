package dsaqueue

func FirstNonRepeatingCharacterInString(stream string) string {
	queue := []byte{}
	frequecny := make(map[byte]int)
	var res string
	for i := 0; i < len(stream); i++ {
		frequecny[stream[i]]++
		queue = append(queue, stream[i])

		for len(queue) > 0 && frequecny[queue[0]] > 1 {
			queue = queue[1:]
		}
		if len(queue) == 0 {
			res += "#"
		} else {
			res += string(queue[0])
		}
	}
	return res
}

/*
GFG – First Non-Repeating Character in a Stream

Note:
This is a GeeksforGeeks problem, not an exact LeetCode problem.
The closest LeetCode problem is LC Q. 387 – First Unique Character in a String,
but LC 387 asks for the index of the first non-repeating character, not the
first non-repeating character after every position in a stream.

Problem Statement:

Given a string A consisting of lowercase alphabets, find the first
non-repeating character in the string for every character added to the stream.

If there is no non-repeating character at any position, append '#' to
the result.

Examples:

Example 1:

Input:
A = "aabc"

Output:
" a#bb"

Correct Output:
"a#bb"

Explanation:
- After reading 'a': First non-repeating character is 'a'.
- After reading "aa": No non-repeating character, so '#'.
- After reading "aab": First non-repeating character is 'b'.
- After reading "aabc": First non-repeating character is 'b'.

Example 2:

Input:
A = "zz"

Output:
"z#"

*/
