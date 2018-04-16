package ex02

import (
	"unsafe"
	"reflect"
)

func IsCyclic(v interface{}) bool {
	seen := make(map[unsafe.Pointer]bool)
	return isCyclic(reflect.ValueOf(v), seen)
}

func isCyclic(x reflect.Value, seen map[unsafe.Pointer]bool) bool {
	switch x.Kind() {
	case reflect.Struct:
		for i := 0; i < x.NumField(); i++ {
			v := x.Field(i)
			if !v.CanAddr() {
				return false
			}
			ptr := unsafe.Pointer(v.UnsafeAddr())
			if _, ok := seen[ptr]; ok {
				return true
			}
			seen[ptr] = true
			return isCyclic(v, seen)
		}
	case reflect.Ptr:
		return isCyclic(x.Elem(), seen)
	default:
		return false
	}

	panic("unreachable")
}
