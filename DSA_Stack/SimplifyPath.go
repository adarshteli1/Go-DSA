package dsastack

import (
	"strings"
)

func SimplifyPath(path string) string {
	parts := strings.Split(path, "/")
	stack := []string{}
	for i := 0; i < len(parts); i++ {
		switch parts[i] {
		case ".":
			continue
		case "":
			continue
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, parts[i])
		}
	}
	return "/" + strings.Join(stack, "/")
}
