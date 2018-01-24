package ex10

import "sort"

func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		j := s.Len() - i - 1
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}
