package main

import "testing"

func Test_comma(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{"123456.123", "123,456.123"},
		{"1234567.123456", "1,234,567.123456"},
		{"123.123", "123.123"},
		{"0.123", "0.123"},
	}
	for _, c := range cases {
		if act := comma(c.in); c.out != act {
			t.Errorf("expected: %s, but: %s", c.out, act)
		}
	}
}
