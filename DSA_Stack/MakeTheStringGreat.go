package dsastack

func MakeTheStringGreat(s string) string {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			differnce := int(top) - int(s[i])
			if differnce == 32 || differnce == -32 {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, s[i])
	}
	return string(stack)
}
