package validation

import (
	"fmt"
	"reflect"
	"strings"
)

// Required Validator ensures a field is "present", certain "falsey" values are handled by
// boolean options (eg AllowEmptyArray). They all default to deny (eg, deny empty strings)
type Required struct {
	AllowEmptyString bool
	AllowEmptyArray  bool
	AllowEmptyMap    bool
	AllowEmptySlice  bool
}

// IsValid determines if the given value satisfies a "Required" rule.
func (r *Required) IsValid(value interface{}) (bool, []error) {
	if value == nil {
		return false, singleErrorSlice("nil", value)
	}

	kind := reflect.TypeOf(value).Kind()
	switch kind {
	case reflect.Struct, reflect.Bool, reflect.Int, reflect.Float32, reflect.Float64:
		return true, nil
	case reflect.String:
		return r.stringIsValid(value.(string))
	case reflect.Map, reflect.Array, reflect.Slice:

		// Decide if we'll allow an empty collection
		var override bool
		switch kind {
		case reflect.Map:
			override = r.AllowEmptyMap
		case reflect.Array:
			override = r.AllowEmptyArray
		case reflect.Slice:
			override = r.AllowEmptySlice
		}

		return r.collectionLengthValid(kind.String(), reflect.ValueOf(value), override)
	}

	// Default case
	return false, singleErrorSlice("Unknown", value)
}

// stringIsValid ensures a string is valid (Not empty after trimming)
func (r *Required) stringIsValid(value string) (bool, []error) {
	if strings.TrimSpace(value) != "" || r.AllowEmptyString {
		return true, nil
	} else {
		return false, singleErrorSlice("String", value)
	}
}

// collectionLengthValid returns true if it contains a value (ie, nonempty). False if empty, unless the option is set.
func (r *Required) collectionLengthValid(colType string, value reflect.Value, override bool) (bool, []error) {
	if value.Len() > 0 || override {
		return true, nil
	} else {
		return false, singleErrorSlice(colType, value)
	}
}

// singleErrorSlice is a helper to match the interface []error
func singleErrorSlice(args ...interface{}) []error {
	return []error{fmt.Errorf("%s value(%v) is not a valid value for Required", args...)}
}
