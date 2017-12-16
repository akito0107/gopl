package main

import "fmt"

const size = 10

func main() {
	test := [size]int{1, 2, 3, 4, 5, 6}
	fmt.Println(test)
	reverse(&test)
	fmt.Println(test)
}

func reverse(s *[size]int) {
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
