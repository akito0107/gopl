package main

import "testing"

func Test_reverse(t *testing.T) {
	cases := []struct {
		in  []byte
		out []byte
	}{
		{
			in:  []byte("あいうえお"),
			out: []byte("おえういあ"),
		},
		{
			in:  []byte("abcあいうえお"),
			out: []byte("おえういあcba"),
		},
	}
	for _, c := range cases {
		if act := Reverse(c.in); !equals(act, c.out) {
			t.Errorf("expected: %s, but: %s", string(c.out), string(act))
		}
	}
}

func equals(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	return string(a) == string(b)
}
