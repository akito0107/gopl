package validator

import (
	"regexp"
)

var emailRegex = regexp.MustCompile(`^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)

func init() {
	Register("email", EmailValidator)
}

func EmailValidator(v interface{}) bool {
	value, ok := v.(string)
	if !ok {
		return false
	}
	return emailRegex.Match([]byte(value))
}
