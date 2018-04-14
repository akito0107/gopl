package sexpr

import (
	"fmt"
	"io"
	"reflect"
	"text/scanner"
)

type Decoder struct {
	lex *lexer
	r   io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	lex := lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	return &Decoder{lex: &lex, r: r}
}

func (d *Decoder) Decode(v interface{}) (err error) {
	d.lex.scan.Init(d.r)
	d.lex.next()
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", d.lex.scan.Position, x)
		}
	}()
	read(d.lex, reflect.ValueOf(v).Elem())
	return nil
}
