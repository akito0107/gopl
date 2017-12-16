package main

import (
	"crypto/sha256"
	"fmt"

	"github.com/akito0107/gopl/ch02/popcount"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Println(bitcount(c1, c2))
}

func bitcount(a [32]byte, b [32]byte) int {
	var sum int
	for i := 0; i < 32; i++ {
		c := uint64(a[i] ^ b[i])
		sum += popcount.PopCount(c)
	}
	return sum
}
