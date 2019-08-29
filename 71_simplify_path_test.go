package leetcode

import "strings"

func simplifyPath(path string) string {
	segs := strings.Split(path, "/")
	if len(segs) < 2 {
		return "/"
	}

	var stack []string
	for i := 1; i < len(segs); i++ {
		if len(segs[i]) == 0 {
			continue
		}
		if segs[i] == ".." {
			n := len(stack) - 1
			if n < 0 {
				n = 0
			}
			stack = stack[:n]
		} else if segs[i] != "." {
			stack = append(stack, segs[i])
		}
	}

	return "/" + strings.Join(stack, "/")
}
