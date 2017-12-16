package main

import "bytes"

func comma(s string) string {
	if len(s) <= 3 {
		return s
	}
	var b bytes.Buffer
	init := len(s) % 3
	if init == 0 {
		init = 3
	}
	b.WriteString(s[0:init])
	for i := init; (i + 3) <= len(s); i += 3 {
		b.WriteString(",")
		b.WriteString(s[i : i+3])
	}
	return b.String()
}
