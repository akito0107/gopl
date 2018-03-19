package main

import (
	"bytes"
	"reflect"
	"testing"
	"unicode/utf8"
)

func TestRun(t *testing.T) {
	cases := []struct {
		name    string
		in      string
		counts  map[rune]int
		utflen  [utf8.UTFMax + 1]int
		invalid int
	}{
		{
			name:    "single word",
			in:      "test",
			counts:  map[rune]int{'t': 2, 'e': 1, 's': 1},
			utflen:  [utf8.UTFMax + 1]int{0, 4, 0, 0, 0},
			invalid: 0,
		},
		{
			name:    "日本語",
			in:      "こんにちは世界",
			counts:  map[rune]int{'こ': 1, 'ん': 1, 'に': 1, 'ち': 1, 'は': 1, '世': 1, '界': 1},
			utflen:  [utf8.UTFMax + 1]int{0, 0, 0, 7, 0},
			invalid: 0,
		},
		{
			name:    "4byte文字",
			in:      "𠮷野屋で𩸽頼んで𠮟られる",
			counts:  map[rune]int{'𠮷': 1, '野': 1, '屋': 1, 'で': 2, '𩸽': 1, '頼': 1, 'ん': 1, '𠮟': 1, 'ら': 1, 'れ': 1, 'る': 1},
			utflen:  [utf8.UTFMax + 1]int{0, 0, 0, 9, 3},
			invalid: 0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			counts, utflen, invalid := Run(bytes.NewBufferString(c.in))
			if !reflect.DeepEqual(counts, c.counts) {
				t.Errorf("counts must be same: expect: %+v, actual: %+v", c.counts, counts)
			}
			if !reflect.DeepEqual(utflen, c.utflen) {
				t.Errorf("utflen must be same: expect: %+v, actual: %+v", c.utflen, utflen)
			}
			if invalid != c.invalid {
				t.Errorf("invalid must be same: expect: %d, actual: %d", c.invalid, invalid)
			}
		})
	}
}
