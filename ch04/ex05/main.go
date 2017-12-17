package main

import "fmt"

func main() {
	test := []string{"test", "hoge", "hoge", "fuga"}
	fmt.Println(test)
	res := uniq(test)
	fmt.Println(res)
	fmt.Println(test)
}

func uniq(strs []string) []string {
	if len(strs) == 1 {
		return strs[:]
	}
	i := 0
	for _, s := range strs {
		if s == strs[i] {
			continue
		}
		i++
		strs[i] = s
	}
	return strs[:i+1]
}
