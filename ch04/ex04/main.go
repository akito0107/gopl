package main

import "fmt"

func main() {
	test := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(test)
	res := rotate(test, 2)
	fmt.Println(res)
	fmt.Println(test)
}

func rotate(s []int, i int) []int {
	return append(s[i:], s[:i]...)
}
