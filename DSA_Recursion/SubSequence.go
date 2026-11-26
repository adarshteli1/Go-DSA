package dsarecursion

import "fmt"

func Subsequence(p, up string) {
	if len(up) == 0 {
		fmt.Print(p)
		return
	}
	ch := up[0]
	up = up[1:]
	Subsequence(p+string(ch), up)
	Subsequence(p, up)
}
