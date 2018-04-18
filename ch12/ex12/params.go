package params

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/akito0107/gopl/ch12/ex12/validator"
)

func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	fields := make(map[string]reflect.Value)
	validators := make(map[string]string)

	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		validatorInfo := tag.Get("validate")
		if validatorInfo != "" {
			validators[name] = validatorInfo
			//if !res {
			//	return fmt.Errorf("validate failed filed: %s, rule: %s", fieldInfo.Name, validatorInfo)
			//}
		}
		fields[name] = v.Field(i)
	}

	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue
		}

		for _, value := range values {
			if validateInfo, ok := validators[name]; ok {
				if !validator.Validate(validateInfo, value) {
					return fmt.Errorf("validate failed filed: %s, rule: %s", name, validateInfo)
				}
			}
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s, %v", name, err)
				}
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 110, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}
