package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int)
	dup := make(map[string]bool)

	files := os.Args[1:]
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		counts[arg] = make(map[string]int)
		dup[arg] = checkDuplicateLines(f, counts[arg])
		f.Close()
	}
	for filename, needPrint := range dup {
		if !needPrint {
			continue
		}
		fmt.Printf("%s\n", filename)
		for line, n := range counts[filename] {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}

func checkDuplicateLines(f *os.File, counts map[string]int) bool {
	res := false
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			res = true
		}
	}
	return res
}
