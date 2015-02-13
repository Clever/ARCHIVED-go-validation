package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	Validator *Required
	Value     interface{}
	Expected  bool
}

var requiredCases = [...]testCase{
	testCase{
		Value:    "",
		Expected: false,
	},
	testCase{
		Value: "",
		Validator: &Required{
			AllowEmptyString: true,
		},
		Expected: true,
	},
	testCase{
		Value:    struct{}{},
		Expected: true,
	},
	testCase{
		Value:    nil,
		Expected: false,
	},
	testCase{
		Value:    "    ",
		Expected: false,
	},
	testCase{
		Value:     "    ",
		Validator: &Required{AllowEmptyString: true},
		Expected:  true,
	},
	testCase{
		Value:    "testing   ",
		Expected: true,
	},
	testCase{
		Value:    false,
		Expected: true,
	},
	testCase{
		Value:    true,
		Expected: true,
	},
	testCase{
		Value:    392323,
		Expected: true,
	},
	testCase{
		Value:    392323.9324,
		Expected: true,
	},
	testCase{
		Value:    struct{ Name string }{Name: "Yes"},
		Expected: true,
	},
	testCase{
		// Empty Array
		Value:    [...]interface{}{},
		Expected: false,
	},
	testCase{
		// Empty Array, with option
		Value: [...]interface{}{},
		Validator: &Required{
			AllowEmptyArray: true,
		},
		Expected: true,
	},
	testCase{
		// Array w/ 3 empty structs
		Value: [...]interface{}{
			struct{}{},
			struct{}{},
			struct{}{},
		},
		Expected: true,
	},
	testCase{
		Value: [...]interface{}{
			struct{ Name string }{Name: "Yes"},
			struct{ Name string }{Name: "No"},
		},
		Expected: true,
	},
	testCase{
		// Empty Slice
		Value:    []interface{}{},
		Expected: false,
	},
	testCase{
		// Empty Slice, w/ option
		Value: []interface{}{},
		Validator: &Required{
			AllowEmptySlice: true,
		},
		Expected: true,
	},
	testCase{
		// Full slice
		Value:    []uint{1, 3, 4, 63434},
		Expected: true,
	},
	testCase{
		// empty map literal
		Value:    map[int]bool{},
		Expected: false,
	},
	testCase{
		// empty map literal, with option
		Value:     map[int]bool{},
		Validator: &Required{AllowEmptyMap: true},
		Expected:  true,
	},
	testCase{
		Value:    map[int]bool{1: true, 2: false},
		Expected: true,
	},
}

func TestRequired(t *testing.T) {

	for _, c := range requiredCases {
		if c.Validator == nil {
			c.Validator = &Required{}
		}
		valid, errors := c.Validator.IsValid(c.Value)
		assert.Equal(
			t,
			c.Expected,
			valid,
			"c.Validator.IsValid(%#v) w/ returned unexpected answer (%v)", c.Value, c.Expected,
		)
		if !c.Expected {
			assert.Equal(t, 1, len(errors))
		}
	}
}
