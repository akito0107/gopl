package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	test := []byte("あいうえお")
	fmt.Println(test)
	test = Reverse(test)
	fmt.Println(string(test))
}
func _rev(b []byte) {
	size := len(b)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[size-1-i] = b[size-1-i], b[i]
	}
}

func Reverse(b []byte) []byte {
	for i := 0; i < len(b); {
		_, size := utf8.DecodeRune(b[i:])
		_rev(b[i : i+size])
		i += size
	}
	_rev(b)
	return b
}
