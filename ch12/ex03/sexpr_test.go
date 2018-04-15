package ex03

import (
	"bytes"
	"reflect"
	"testing"
)

func Test_encode(t *testing.T) {

	testStruct := struct {
		v interface{}
	}{
		v: []int{1, 2, 3},
	}

	cases := []struct {
		name string
		in   interface{}
		out  string
	}{
		{
			name: "boolean(true)",
			in:   true,
			out:  "t",
		},
		{
			name: "boolean(false)",
			in:   false,
			out:  "nil",
		},
		{
			name: "float",
			in:   12.345,
			out:  "12.345",
		},
		{
			name: "complex",
			in:   complex(1, 1),
			out:  "#C(1 1)",
		},
		{
			name: "complex",
			in:   complex(1.2345, 1.2345),
			out:  "#C(1.2345 1.2345)",
		},
		{
			name: "interface with struct",
			in:   testStruct,
			out:  `((v("[]int" (1 2 3))))`,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var b bytes.Buffer
			encode(&b, reflect.ValueOf(c.in))

			if act := b.String(); act != c.out {
				t.Errorf("actual: %s, but expect: %s", act, c.out)
			}
		})
	}
}
