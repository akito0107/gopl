package main

import "sort"

type RuneSlice []rune

func (r RuneSlice) Len() int {
	return len(r)
}
func (r RuneSlice) Less(i, j int) bool {
	return r[i] < r[j]
}
func (r RuneSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func IsAnagram(src string, tar string) bool {
	s := []rune(src)
	sort.Sort(RuneSlice(s))
	t := []rune(tar)
	sort.Sort(RuneSlice(t))
	return string(s) == string(t)
}
