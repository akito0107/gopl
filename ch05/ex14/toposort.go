package main

import (
	"fmt"

	"github.com/akito0107/gopl/ch05/links"
)

var prereqs = map[string][]string{
	"algorithm":      {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	result := run(prereqs)

	for i, course := range result {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

var result []string

func run(items map[string][]string) []string {
	//for k := range items {
	//	keys = append(keys, k)
	//}
	links.BreadthFirst(topoSort(), []string{"programming languages"})

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func topoSort() func(item string) []string {
	return func(item string) []string {
		result = append(result, item)
		return prereqs[item]
	}
}
