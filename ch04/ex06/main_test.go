package main

import "testing"

func Test_compact(t *testing.T) {
	cases := []struct {
		in  []byte
		out []byte
	}{
		{
			in:  []byte("test"),
			out: []byte("test"),
		},
		{
			in:  []byte("tes t"),
			out: []byte("tes t"),
		},
		{
			in:  []byte("tes  t"),
			out: []byte("tes t"),
		},
		{
			in:  []byte("日 本  語"),
			out: []byte("日 本 語"),
		},
		{
			in:  []byte("  先頭"),
			out: []byte(" 先頭"),
		},
		{
			in:  []byte("  先頭 と  末尾   "),
			out: []byte(" 先頭 と 末尾 "),
		},
		{
			in:  []byte("  先頭 と  末尾  a  "),
			out: []byte(" 先頭 と 末尾 a "),
		},
	}
	for _, c := range cases {
		if act := compact(c.in); !equals(act, c.out) {
			t.Errorf("expected: %v, but: %v", c.out, act)
		}
	}
}

func equals(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	return string(a) == string(b)
}
