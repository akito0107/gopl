package ex02

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

var writer io.Writer = os.Stdout

const LIMIT = 10

func Display(name string, x interface{}) {
	fmt.Fprintf(writer, "Display %s (%T): \n", name, x)
	display(writer, name, reflect.ValueOf(x), 0)
}

func display(w io.Writer, path string, v reflect.Value, limit int) {
	if limit >= LIMIT {
		fmt.Fprintf(w, "%s... too deep", path)
		return
	}
	limit += 1
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Fprintf(w, "%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(writer, fmt.Sprintf("%s[%d]", path, i), v.Index(i), limit)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(writer, fieldPath, v.Field(i), limit)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(writer, fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key), limit)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Fprintf(writer, "%s = nil\n", path)
		} else {
			display(writer, fmt.Sprintf("(*%s)", path), v.Elem(), limit)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Fprintf(writer, "%s = nil\n", path)
		} else {
			fmt.Fprintf(writer, "%s.type = %s\n", path, v.Elem().Type())
			display(writer, path+".value", v.Elem(), limit)
		}
	default:
		fmt.Fprintf(writer, "%s = %s\n", path, formatAtom(v))
	}
}
