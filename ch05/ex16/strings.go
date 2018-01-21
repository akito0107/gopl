package ex16

import "strings"

func Join(sep string, args ...string) string {
	return strings.Join(args, sep)
}
