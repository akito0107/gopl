package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	test := []byte("あいうえお")
	fmt.Println(string(test))
	test2 := Reverse(test)
	fmt.Println(string(test))
	fmt.Println(string(test2))
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
