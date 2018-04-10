package ex01

import (
	"reflect"
	"strconv"
	"strings"
)

func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {

	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	case reflect.Struct:
		var buf strings.Builder
		buf.WriteString(v.Type().String())
		buf.WriteString("{")
		for i := 0; i < v.NumField(); i++ {
			buf.WriteString(v.Type().Field(i).Name)
			buf.WriteString(": ")
			field := v.FieldByName(v.Type().Field(i).Name)
			buf.WriteString(formatAtom(field))
			if i != v.NumField()-1 {
				buf.WriteString(", ")
			}
		}
		buf.WriteString("}")

		return buf.String()
	case reflect.Array:
		var buf strings.Builder
		buf.WriteString(v.Type().String())
		buf.WriteString("{")
		for i := 0; i < v.Len(); i++ {
			buf.WriteString(formatAtom(v.Index(i)))
			if i != v.Len()-1 {
				buf.WriteString(", ")
			}
		}
		buf.WriteString("}")

		return buf.String()
	default:
		return v.Type().String() + " value"
	}
}
