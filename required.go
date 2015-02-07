package validator

import (
	"fmt"
	"reflect"
	"strings"
)

// Required Validator is very simple
type RequiredValidator struct{}

// IsValid determines if the given value satisfies a "Required" rule.
func (r *RequiredValidator) IsValid(value interface{}) (bool, []error) {

	if value == nil {
		return false, singleErrorSlice("nil is not a valid value for RequiredValidator")
	}

	switch reflect.TypeOf(value).Kind() {
	case reflect.String:
		return r.stringIsValid(value.(string))

	default:
		panic("unhandle value type")
		return false, singleErrorSlice("Cannot handle value(%v)", value)
	}
}

// stringIsValid ensures a string is valid (Not empty after trimming)
func (r *RequiredValidator) stringIsValid(value string) (bool, []error) {
	if strings.TrimSpace(value) != "" {
		return true, nil
	} else {
		return false, singleErrorSlice("String value(%v) doesnt not satisfy RequiredValidator", value)
	}
}

func singleErrorSlice(e string, args ...interface{}) []error {
	return append(make([]error, 0), fmt.Errorf(e, args...))
}
