package encode

import (
	"bytes"
	"fmt"
	"reflect"
)

func encode(buf *bytes.Buffer, depth int, v reflect.Value) error {
	if IsZero(v) {
		return nil
	}
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())
	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())
	case reflect.Ptr:
		return encode(buf, depth, v.Elem())
	case reflect.Bool:
		if v.Bool() {
			buf.WriteString("t")
		} else {
			buf.WriteString("nil")
		}
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%G", v.Float())
	case reflect.Complex64, reflect.Complex128:
		fmt.Fprintf(buf, "#C(%G %G)", real(v.Complex()), imag(v.Complex()))
	case reflect.Array, reflect.Slice:
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if i != 0 {
				fmt.Fprintf(buf, "%*s", depth, "")
			}
			if err := encode(buf, depth, v.Index(i)); err != nil {
				return err
			}
			if i != v.Len()-1 {
				buf.WriteString("\n")
			}
		}
		buf.WriteString(")")
	case reflect.Struct:
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if IsZero(v.Field(i)) {
				continue
			}
			if i > 0 {
				buf.WriteByte(' ')
			}
			tag := v.Type().Field(i).Tag.Get("sexpr")
			if tag != "" {
				fmt.Fprintf(buf, "(%s ", tag)
			} else {
				fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			}
			d := depth + len(v.Type().Field(i).Name) + 3
			if err := encode(buf, d, v.Field(i)); err != nil {
				return err
			}
			buf.WriteString(")")
			if i != v.NumField()-1 {
				buf.WriteString("\n")
			}
		}
		buf.WriteString(")")

	case reflect.Map:
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if i == 0 {
				buf.WriteByte('(')
			} else {
				fmt.Fprintf(buf, "%*s%s", depth, "", "(")
			}
			if err := encode(buf, depth, key); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, depth, v.MapIndex(key)); err != nil {
				return err
			}
			buf.WriteString(")")
			if i != len(v.MapKeys())-1 {
				buf.WriteString("\n")
			}
		}
		buf.WriteString(")")

	case reflect.Interface:
		buf.WriteByte('(')
		fmt.Fprintf(buf, "\"%s\" ", v.Elem().Type().String())
		encode(buf, depth, v.Elem())
		buf.WriteString(")\n")

	default:
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func IsZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && IsZero(v.Index(i))
		}
		return z
	case reflect.Struct:
		z := true
		for i := 0; i < v.NumField(); i++ {
			z = z && IsZero(v.Field(i))
		}
		return z
	case reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Ptr:
		return v.IsNil() || IsZero(reflect.Indirect(v))
	default:
		if !v.CanInterface() {
			return true
		}
		z := reflect.Zero(v.Type())
		return v.Interface() == z.Interface()
	}
}

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, 0, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
