package validator

import "log"

type Validator func(v interface{}) bool

var validators = map[string]Validator{}

func Register(name string, validator Validator) {
	validators[name] = validator
}

func Validate(name string, value interface{}) bool {
	validator, ok := validators[name]
	if !ok {
		log.Printf("unknown validator type: %s\n", name)
		return false
	}
	return validator(value)
}
