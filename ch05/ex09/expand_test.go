package ex09

import "testing"

func TestExpand(t *testing.T) {
	cases := []struct {
		name string
		in   string
		out  string
		fn   func(string) string
	}{
		{
			name: "basic",
			in:   "$foo",
			out:  "hoge",
			fn: func(s string) string {
				if s == "foo" {
					return "hoge"
				}
				return ""
			},
		},
		{
			name: "multiple",
			in:   "$bar $foo",
			out:  "fuga hoge",
			fn: func(s string) string {
				if s == "foo" {
					return "hoge"
				} else if s == "bar" {
					return "fuga"
				}
				return ""
			},
		},
		{
			name: "mixed",
			in:   "$bar notmatch $foo",
			out:  "fuga notmatch hoge",
			fn: func(s string) string {
				if s == "foo" {
					return "hoge"
				} else if s == "bar" {
					return "fuga"
				}
				return ""
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if act := Expand(c.in, c.fn); c.out != act {
				t.Errorf("must be %s but got: %s \n", c.out, act)
			}
		})
	}
}
