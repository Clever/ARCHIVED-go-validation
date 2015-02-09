package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	Value    interface{}
	Options  map[Option]bool
	Expected bool
}

var requiredCases = [...]testCase{
	testCase{
		Value:    "",
		Expected: false,
	},
	testCase{
		Value: "",
		Options: map[Option]bool{
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
		Value:    "    ",
		Options:  map[Option]bool{AllowEmptyString: true},
		Expected: true,
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
		Options: map[Option]bool{
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
		Options: map[Option]bool{
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
		Value:    map[int]bool{},
		Options:  map[Option]bool{AllowEmptyMap: true},
		Expected: true,
	},
	testCase{
		Value:    map[int]bool{1: true, 2: false},
		Expected: true,
	},
}

func TestRequired(t *testing.T) {
	rv := Required{}

	for _, c := range requiredCases {
		rv.SetOpts(c.Options)
		valid, errors := rv.IsValid(c.Value)
		assert.Equal(
			t,
			c.Expected,
			valid,
			"rv.IsValid(%#v) w/ options(%v) returned unexpected answer (%v)", c.Value, c.Options, c.Expected,
		)
		if !c.Expected {
			assert.Equal(t, 1, len(errors))
		}
	}
}
