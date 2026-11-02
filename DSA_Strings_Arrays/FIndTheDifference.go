package dsastringsarrays

func FindTheDiffernece(s, t string) byte {
	var Letters [26]int
	for i := 0; i < len(t); i++ {
		Letters[t[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		Letters[s[i]-'a']--
	}
	for i, val := range Letters {
		if val > 0 {
			return byte(i + 'a')
		}
	}
	return 0
}

// or we can solve by XOR

// func FindTheDiffernece(s, t string) byte {
// 	var a byte
// 	for i:=range s{
// 		a^=s[i]
// 	}
// 	for i:=range t{
// 		a^=t[i]
// 	}
// 	return a
// }
