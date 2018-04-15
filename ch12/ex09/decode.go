package ex09

import (
	"fmt"
	"io"
	"reflect"
	"strconv"
	"text/scanner"
)

type Decoder struct {
	lex *lexer
	r   io.Reader
}

type Symbol struct {
	Name string
}

type String struct {
	Val string
}

type Int struct {
	Val int
}

type StartList struct{}
type EndList struct{}

func NewDecoder(r io.Reader) *Decoder {
	lex := lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(r)
	return &Decoder{lex: &lex, r: r}
}

func (d *Decoder) Decode(v interface{}) (err error) {
	d.lex.next()
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", d.lex.scan.Position, x)
		}
	}()
	read(d.lex, reflect.ValueOf(v).Elem())
	return nil
}

func (d *Decoder) Token() (interface{}, error) {
	d.lex.next()
	switch d.lex.token {
	case scanner.Ident:
		return &Symbol{d.lex.text()}, nil
	case scanner.String:
		s, _ := strconv.Unquote(d.lex.text())
		return &String{s}, nil
	case scanner.Int:
		i, _ := strconv.Atoi(d.lex.text())
		return &Int{i}, nil
	case '(':
		return &StartList{}, nil
	case ')':
		return &EndList{}, nil
	case scanner.EOF:
		return nil, io.EOF
	}
	panic(fmt.Sprintf("unexpected token %q", d.lex.text()))
}
