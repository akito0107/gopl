package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	counts := Wordfreq(in)
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q \t %d\n", c, n)
	}
}

func Wordfreq(src io.Reader) map[string]int {
	scanner := bufio.NewScanner(src)
	scanner.Split(bufio.ScanWords)
	counts := make(map[string]int)
	for scanner.Scan() {
		t := scanner.Text()
		counts[t]++
	}
	return counts
}
