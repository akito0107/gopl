package ex09

import "regexp"

var pattern = regexp.MustCompile(`\$\w+`)

func Expand(s string, f func(string) string) string {
	w := func(src string) string {
		return f(src[1:])
	}
	return pattern.ReplaceAllStringFunc(s, w)
}
