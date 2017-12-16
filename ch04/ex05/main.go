package main

import "fmt"

func main() {
	test := []string{"test", "test", "hoge", "fuga"}
	fmt.Println(test)
	res := uniq(test)
	fmt.Println(res)
	fmt.Println(test)
}

func uniq(s []string) []string {
	if len(s) == 1 {
		return s
	}
	buf := s[0]
	res := s[0:1]
	for i := 1; i < len(s); i++ {
		if buf != s[i] {
			res = append(res, s[i])
		}
		buf = s[i]
	}
	return res
}
