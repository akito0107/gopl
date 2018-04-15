package ex10

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"text/scanner"
)

func Unmarshal(data []byte, out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader(data))
	lex.next()
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}

func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		} else if lex.text() == "true" {
			v.SetBool(true)
			lex.next()
			return
		} else if lex.text() == "false" {
			v.SetBool(false)
			lex.next()
			return
		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text())
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text())
		if isSignedInt(v) { // Exercise 12.10
			v.SetInt(int64(i))
		} else {
			v.SetUint(uint64(i)) // Exercise 12.10
		}
		lex.next()
		return
	case scanner.Float:
		f, _ := strconv.ParseFloat(lex.text(), 64)
		v.SetFloat(f)
		lex.next()
		return
	case '(':
		lex.next()
		readList(lex, v)
		lex.next()
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array:
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}
	case reflect.Slice:
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}
	case reflect.Struct:
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want filed name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}
	case reflect.Map:
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}
	case reflect.Interface:
		t, _ := strconv.Unquote(lex.text())
		value := reflect.New(typeOf(t)).Elem()
		lex.next()
		read(lex, value)
		v.Set(value)

	default:
		panic(fmt.Sprintf("cannot decode list int %v", v.Type()))
	}
}

func isSignedInt(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return true

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return false

	default:
		panic(fmt.Sprintf("isSignedInt: v.Kind(%d) not supported", v.Kind()))
	}
}

var maps = map[string]reflect.Type{
	"int":    reflect.TypeOf(int(0)),
	"int8":   reflect.TypeOf(int8(0)),
	"int16":  reflect.TypeOf(int16(0)),
	"int32":  reflect.TypeOf(int32(0)),
	"int64":  reflect.TypeOf(int64(0)),
	"uint":   reflect.TypeOf(uint(0)),
	"uint8":  reflect.TypeOf(uint8(0)),
	"uint16": reflect.TypeOf(uint16(0)),
	"uint32": reflect.TypeOf(uint32(0)),
	"uint64": reflect.TypeOf(uint64(0)),
	"bool":   reflect.TypeOf(false),
	"string": reflect.TypeOf(""),
}

func typeOf(tName string) reflect.Type {
	t, ok := maps[tName]
	if ok {
		return t
	}

	if strings.HasPrefix(tName, "[]") {
		return reflect.SliceOf(typeOf(tName[2:]))
	}

	if tName[0] == '[' {
		i := strings.Index(tName, "]")
		if i > 0 {
			len, _ := strconv.Atoi(tName[1:i])
			return reflect.ArrayOf(len, typeOf(tName[i+1:]))
		}
	}

	if strings.HasPrefix(tName, "map") {
		i := strings.Index(tName, "]")
		if i > 0 {
			return reflect.MapOf(typeOf(tName[4:i]), typeOf(tName[i+1:]))
		}
	}

	panic(fmt.Sprintf("%s not supported yet\n", tName))
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}
