package code

import (
	"strings"
)

// https://leetcode.com/problems/simplify-path/

func simplifyPath(path string) string {
	if path[0] != '/' {
		return ""
	}
	steps := strings.Split(path[1:], "/")
	new_steps := make([]string, 0, len(steps))
	for _, step := range steps {
		if len(step) < 1 {
			continue
		}
		if step == "." {
			continue
		}
		if step == ".." {
			if len(new_steps) > 0 {
				new_steps = new_steps[:len(new_steps)-1]
			}
		} else {
			new_steps = append(new_steps, step)
		}
	}
	return "/" + strings.Join(new_steps, "/")
}
