package main

import "log"

func min(args ...int) int {
	if len(args) == 0 {
		log.Fatal("at least 1 arg") // or return 0
	}
	var m int
	for _, i := range args {
		if i < m {
			m = i
		}
	}
	return m
}

func max(args ...int) int {
	if len(args) == 0 {
		log.Fatal("at least 1 arg") // or return 0
	}
	var m int
	for _, i := range args {
		if i > m {
			m = i
		}
	}
	return m
}
