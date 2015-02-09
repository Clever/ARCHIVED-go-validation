package validation

import (
	"fmt"
	"reflect"
	"strings"
)

type Option uint8

const (
	// Allows empty string, array or map. Default is false
	AllowEmptyString Option = iota
	AllowEmptyArray
	AllowEmptyMap
	AllowEmptySlice
)

// Required Validator is very simple
type Required struct {
	opts map[Option]bool
}

func (r *Required) SetOpts(opts map[Option]bool) {
	r.opts = opts
}

// IsValid determines if the given value satisfies a "Required" rule.
func (r *Required) IsValid(value interface{}) (bool, []error) {
	if value == nil {
		return false, singleErrorSlice("nil is not a valid value for Required")
	}

	kind := reflect.TypeOf(value).Kind()
	switch kind {
	case reflect.Struct, reflect.Bool, reflect.Int, reflect.Float32, reflect.Float64:
		return true, nil
	case reflect.String:
		return r.stringIsValid(value.(string))
	case reflect.Map, reflect.Array, reflect.Slice:
		var override bool
		switch kind {
		case reflect.Map:
			override = r.opts[AllowEmptyMap]
		case reflect.Array:
			override = r.opts[AllowEmptyArray]
		case reflect.Slice:
			override = r.opts[AllowEmptySlice]
		}
		return r.collectionLengthOK(kind.String(), reflect.ValueOf(value).Len(), override)
	}

	// Default case
	return false, singleErrorSlice("Cannot handle value(%v)", value)
}

// stringIsValid ensures a string is valid (Not empty after trimming)
func (r *Required) stringIsValid(value string) (bool, []error) {
	if strings.TrimSpace(value) != "" || r.opts[AllowEmptyString] {
		return true, nil
	} else {
		return false, singleErrorSlice("String value(%v) doesnt not satisfy Required", value)
	}
}

// collectionLengthOK returns true if it contains a value (ie, nonempty). False if empty, unless the option is set.
func (r *Required) collectionLengthOK(colType string, length int, override bool) (bool, []error) {
	if length > 0 || override {
		return true, nil
	} else {
		return false, singleErrorSlice("%s length(%d) does not satisfy Required(override=%t)", colType, length, override)
	}
}

// singleErrorSlice is a helper to match the interface []error
func singleErrorSlice(e string, args ...interface{}) []error {
	return append(make([]error, 0), fmt.Errorf(e, args...))
}
