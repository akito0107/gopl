package ex11

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Pack(ptr interface{}) (string, error) {
	var result []string
	type typeset struct {
		filedName string
		value     reflect.Value
	}
	var fields []typeset
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields = append(fields, typeset{
			filedName: name,
			value:     v.Field(i),
		})
	}

	for _, field := range fields {
		paramsStr := toParam(field.value)
		if paramsStr != "" {
			result = append(result, fmt.Sprintf("%s=%s", field.filedName, paramsStr))
		}
	}
	return strings.Join(result, "&"), nil
}

func toParam(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return v.String()
	case reflect.Slice, reflect.Array:
		var buf strings.Builder
		buf.WriteString("[")
		for i := 0; i < v.Len(); i++ {
			if i != 0 {
				buf.WriteString(",")
			}
			buf.WriteString(toParam(v.Index(i)))
		}
		buf.WriteString("]")
		return buf.String()
	default:
		return ""
	}

	return ""
}
