package ex02

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	cases := []struct {
		name string
		in   string
		out  int64
	}{
		{
			name: "basic",
			in:   "test",
			out:  4,
		},
		{
			name: "multibyte",
			in:   "日本語",
			out:  9,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			writer, size := CountingWriter(ioutil.Discard)
			fmt.Fprint(writer, c.in)
			if *size != c.out {
				t.Errorf("size must be %d, but %d \n", *size, c.out)
			}
		})
	}
}
