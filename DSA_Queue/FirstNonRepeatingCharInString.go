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
