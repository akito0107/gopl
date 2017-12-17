package main

import (
	"strings"
	"testing"
)

type tester func(map[string]int) bool

func TestCount(t *testing.T) {
	cases := []struct {
		name   string
		in     string
		tester tester
	}{
		{
			name: "test number",
			in:   "012١٢٣",
			tester: func(in map[string]int) bool {
				return in["number"] == 6
			},
		},
		{
			name: "test control",
			in:   "\u0000\u0001", // NULL / Start of Heading
			tester: func(in map[string]int) bool {
				return in["control"] == 2
			},
		},
		{
			name: "test mark",
			in:   "҈⃠⃣",
			tester: func(in map[string]int) bool {
				return in["mark"] == 3
			},
		},
		{
			name: "test symbol",
			in: "۞৺	௵",
			tester: func(in map[string]int) bool {
				return in["symbol"] == 3
			},
		},
		{
			name: "test letter",
			in:   "ǋabcdefgk",
			tester: func(in map[string]int) bool {
				return in["letter"] == 9
			},
		},
	}

	for _, c := range cases {
		src := strings.NewReader(c.in)
		count, invalid, err := Count(src)
		if err != nil {
			t.Error(err)
		}
		if invalid != 0 {
			t.Errorf("expected no invalid letters %d", invalid)
		}
		if !c.tester(count) {
			t.Errorf("%s: actual %+v", c.name, count)
		}
	}
}
