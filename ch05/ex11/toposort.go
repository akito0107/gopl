package main

import (
	"fmt"
	"log"
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
	courses, err := topoSort(prereqs)
	if err != nil {
		log.Fatal(err)
	}
	for i, course := range courses {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)

	var visitAll func(items []string) error

	visitAll = func(items []string) error {
		for _, item := range items {
			s, ok := seen[item]

			if !s && ok {
				return fmt.Errorf("cyclic error %s", item)
			}

			if !ok {
				seen[item] = false
				if err := visitAll(m[item]); err != nil {
					return err
				}

				seen[item] = true
				order = append(order, item)
			}
		}
		return nil
	}

	for k := range m {
		if err := visitAll([]string{k}); err != nil {
			return nil, err
		}
	}

	return order, nil
}
