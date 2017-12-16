package main

import "testing"

func TestIsAnagram(t *testing.T) {
	type in struct {
		s string
		t string
	}
	cases := []struct {
		in  in
		out bool
	}{
		{in{"abc", "abc"}, true},
		{in{"abc", "cba"}, true},
		{in{"abc", "ccba"}, false},
		{in{"あいうえお", "おえういあ"}, true},
		{in{"ああいうえお", "おえういあ"}, false},
	}
	for _, c := range cases {
		if act := IsAnagram(c.in.s, c.in.t); c.out != act {
			t.Errorf("expected: %s, but: %s", c.out, act)
		}
	}
}
