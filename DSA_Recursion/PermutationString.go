package dsarecursion

import "fmt"

func PermutationString(old, newstring string) {
	if len(old) == 0 {
		fmt.Println(newstring)
		return
	}
	ch := old[0]
	for i := 0; i <= len(newstring); i++ {
		f := newstring[:i]
		s := newstring[i:]
		PermutationString(old[1:], f+string(ch)+s)
	}
}
