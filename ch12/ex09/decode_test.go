package sexpr

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestDecoder_Token(t *testing.T) {

	cases := []struct {
		in  string
		out []interface{}
	}{
		{
			in:  `(Title "Dr. Strangelove")`,
			out: []interface{}{&StartList{}, &Symbol{"Title"}, &String{"Dr. Strangelove"}, &EndList{}},
		},
		{
			in:  `(Year 1923)`,
			out: []interface{}{&StartList{}, &Symbol{"Year"}, &Int{1923}, &EndList{}},
		},
	}

	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			buf := bytes.NewBufferString(c.in)
			decoder := NewDecoder(buf)

			var tokens []interface{}
			for {
				tok, err := decoder.Token()
				if err == io.EOF {
					break
				} else if err != nil {
					t.Fatal(err)
				}
				tokens = append(tokens, tok)
			}

			if !reflect.DeepEqual(c.out, tokens) {
				t.Errorf("must be same expect: %+v, actual: %+v", c.out, tokens)
			}
		})
	}

}
