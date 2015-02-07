package validator

import (
	//"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	Value    interface{}
	Expected bool
}

var requiredCases = [...]testCase{
	testCase{
		Value:    "",
		Expected: false,
	},
	testCase{
		Value:    struct{}{},
		Expected: false,
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
		Value:    [...]interface{}{},
		Expected: false,
	},
	testCase{
		Value: [...]interface{}{
			struct{ Name string }{Name: "Yes"},
			struct{ Name string }{Name: "No"},
		},
		Expected: true,
	},
}

func TestRequiredValidator(t *testing.T) {
	rv := RequiredValidator{}

	for _, c := range requiredCases {
		valid, errors := rv.IsValid(c.Value)
		assert.Equal(t, c.Expected, valid)
		assert.Equal(t, 1, len(errors))
	}
}
