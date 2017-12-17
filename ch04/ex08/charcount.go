package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	counts, invalid, err := Count(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q \t %d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for k, v := range counts {
		fmt.Printf("%s\t%d\n", k, v)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

func Count(src io.Reader) (map[string]int, int, error) {
	in := bufio.NewReader(src)
	counts := make(map[string]int)
	invalid := 0
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, invalid, err
		}
		if unicode.IsControl(r) {
			counts["control"]++
			continue
		}
		if unicode.IsNumber(r) {
			counts["number"]++
			continue
		}
		if unicode.IsMark(r) {
			counts["mark"]++
			continue
		}
		if unicode.IsSymbol(r) {
			counts["symbol"]++
			continue
		}
		if unicode.IsPunct(r) {
			counts["punct"]++
			continue
		}
		if unicode.IsSpace(r) {
			counts["space"]++
			continue
		}
		if unicode.IsLetter(r) {
			counts["letter"]++
			continue
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts["unknown"]++
	}
	return counts, invalid, nil
}
