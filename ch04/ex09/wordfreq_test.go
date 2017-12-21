package main

import (
	"strings"
	"testing"
)

type tester func(map[string]int) bool

func TestWordfreq(t *testing.T) {
	cases := []struct {
		name   string
		in     string
		tester tester
	}{
		{
			name: "ascii",
			in: `abc abc
cde efs asdf asdf`,
			tester: func(in map[string]int) bool {
				return in["abc"] == 2 && in["asdf"] == 2
			},
		},
		{
			name: "japanese",
			in: ` hoge hoge ほげ
hoge fuga ほげ`,
			tester: func(in map[string]int) bool {
				return in["hoge"] == 3 && in["ほげ"] == 2
			},
		},
	}

	for _, c := range cases {
		src := strings.NewReader(c.in)
		count := Wordfreq(src)
		if !c.tester(count) {
			t.Errorf("%s: actual %+v", c.name, count)
		}
	}
}
