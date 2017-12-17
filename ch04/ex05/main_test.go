package main

import (
	"fmt"
	"testing"
)

func Test_reverse(t *testing.T) {
	cases := []struct {
		in  []string
		out []string
	}{
		{
			in:  []string{"test", "test", "fuga", "hoge"},
			out: []string{"test", "fuga", "hoge"},
		},
		{
			in:  []string{"test", "test", "test", "hoge"},
			out: []string{"test", "hoge"},
		},
		{
			in:  []string{"hoge", "test", "test", "test", "hoge"},
			out: []string{"hoge", "test", "hoge"},
		},
	}
	for _, c := range cases {
		fmt.Printf("case: %v \n", c)
		if act := uniq(c.in); !equals(act, c.out) {
			t.Errorf("expected: %v, but: %v", c.out, act)
		}
	}
}

func equals(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
