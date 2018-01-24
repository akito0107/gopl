package ex01

import (
	"fmt"
	"testing"
)

func TestLineCounter_Write(t *testing.T) {
	cases := []struct {
		name string
		in   string
		out  int
	}{
		{
			name: "single line",
			in:   "hoge",
			out:  1,
		},
		{
			name: "multi line",
			in: `hoge
hoge
hoge
`,
			out: 3,
		},
		{
			name: "blank line",
			in: `hoge
hoge

hoge
`,
			out: 4,
		},
		{
			name: "no lines",
			in:   ``,
			out:  0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var l LineCounter
			fmt.Fprint(&l, c.in)
			if int(l) != c.out {
				t.Errorf("line count must be %d, but %d\n", c.out, int(l))
			}
		})
	}
}

func TestWordCounter_Write(t *testing.T) {
	cases := []struct {
		name string
		in   string
		out  int
	}{
		{
			name: "single word",
			in:   "hoge",
			out:  1,
		},
		{
			name: "single line/w multi words",
			in:   "hoge huga piyo hoge",
			out:  4,
		},
		{
			name: "multi line",
			in: `hoge hoge
hoge fuga piyo
hoge hoe hoge
`,
			out: 8,
		},
		{
			name: "no words",
			in:   ``,
			out:  0,
		},
		{
			name: "multibyte words",
			in: `あいうえお　かきくけこ
さしすせそ
たちつてと`,
			out: 4,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var w WordCounter
			fmt.Fprint(&w, c.in)
			if int(w) != c.out {
				t.Errorf("line count must be %d, but %d\n", c.out, int(w))
			}
		})
	}
}
