package dsaheap

func parent(index int) int {
	return (index - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}
func right(i int) int {
	return 2*i + 2
}
