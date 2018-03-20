package ex05_test

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	cases := []struct {
		s    string
		sep  string
		want int
	}{
		{
			"a:b:c",
			":",
			3,
		},
		{
			"a b c",
			" ",
			3,
		},
		{
			"a/b/c/de",
			"/",
			4,
		},
	}

	for _, c := range cases {
		words := strings.Split(c.s, c.sep)
		if got := len(words); c.want != got {
			t.Errorf("Split(%q, %q) returned %d words, want %d", c.s, c.sep, got, c.want)
		}
	}
}
