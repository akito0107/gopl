package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	test := []byte("ほげ   ほげ   ")
	fmt.Println(test)
	test = compact(test)
	fmt.Println(string(test))
}

func compact(b []byte) []byte {
	cnt := 0
	var lastrune rune
	for i := 0; i+cnt < len(b); {
		r, size := utf8.DecodeRune(b[i:])
		if size == 0 {
			break
		}
		if unicode.IsSpace(r) && unicode.IsSpace(lastrune) {
			copy(b[i:], b[i+size:])
			cnt += size
		} else {
			i += size
		}
		lastrune = r
	}
	return b[:len(b)-cnt]
}
