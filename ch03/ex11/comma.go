package main

import (
	"bytes"
	"strings"
)

func comma(s string) string {
	strs := strings.Split(s, ".")
	if len(strs[0]) <= 3 {
		return s
	}
	var b bytes.Buffer
	init := len(strs[0]) % 3
	if init == 0 {
		init = 3
	}
	b.WriteString(strs[0][0:init])
	for i := init; (i + 3) <= len(strs[0]); i += 3 {
		b.WriteString(",")
		b.WriteString(strs[0][i : i+3])
	}
	if len(strs[1]) != 0 {
		b.WriteString(".")
		b.WriteString(strs[1])
	}
	return b.String()
}
