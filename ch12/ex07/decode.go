package ex07

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"reflect"
	"text/scanner"
)

type Decoder struct {
	reader *bufio.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	reader := bufio.NewReader(r)
	return &Decoder{reader: reader}
}

func (d *Decoder) Decode(i interface{}) (err error) {
	line, _, err := d.reader.ReadLine()
	if err != nil {
		return err
	}
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader(line))
	lex.next()
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	if err := read(lex, reflect.ValueOf(i).Elem()); err == io.EOF {
		return nil
	} else if err != nil {
		return err
	}
	return nil
}
