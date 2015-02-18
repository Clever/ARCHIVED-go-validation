package validation

import (
	"fmt"
	"reflect"
	"strings"
)

// Required Validator ensures a field is "present", certain "falsey" values are handled by
// the AllowEmpty option .
type Required struct {
	AllowEmpty bool
}

// IsValid determines if the given value satisfies a "Required" rule.
func (r *Required) IsValid(value interface{}) (bool, error) {
	if value == nil {
		return false, singleError("nil", value)
	}

	rvalue := reflect.ValueOf(value)
	switch rvalue.Kind() {
	case reflect.Struct, reflect.Bool, reflect.Int, reflect.Float32, reflect.Float64:
		return true, nil
	case reflect.String:
		return r.stringIsValid(value.(string))
	case reflect.Map, reflect.Array, reflect.Slice:
		return r.collectionIsValid(rvalue)
	}

	// Default case
	return false, singleError("Unknown", value)
}

// stringIsValid ensures a string is valid (Not empty after trimming)
func (r *Required) stringIsValid(value string) (bool, error) {
	if strings.TrimSpace(value) != "" || r.AllowEmpty {
		return true, nil
	} else {
		return false, singleError("String", value)
	}
}

// collectionIsValid returns true if it contains a value (ie, nonempty). False if empty, unless the option is set.
func (r *Required) collectionIsValid(value reflect.Value) (bool, error) {
	// Decide if we'll allow an empty collection
	var override bool
	switch value.Kind() {
	case reflect.Map:
		override = r.AllowEmpty
	case reflect.Array:
		override = r.AllowEmpty
	case reflect.Slice:
		override = r.AllowEmpty
	}

	if value.Len() > 0 || override {
		return true, nil
	} else {
		return false, singleError(value.Kind(), value)
	}
}

// singleError is a helper to match the interface error
func singleError(args ...interface{}) error {
	return fmt.Errorf("%s value(%v) is not a valid value for Required", args...)
}
