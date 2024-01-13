package curl2struct

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewArgs(t *testing.T) {
	testCases := map[string]struct {
		Given    []string
		Expected Args
	}{
		"should return empty args when given args is nil": {
			Given:    nil,
			Expected: Args{},
		},
		"should return empty args when given args is empty": {
			Given:    []string{},
			Expected: Args{},
		},
		"should return args with given values when given args is not empty or nil": {
			Given:    []string{"a", "b", "\n", "", "c"},
			Expected: Args{"a", "b", "\n", "", "c"},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.Expected, NewArgs(tc.Given))
		})
	}
}

func TestArgs_format(t *testing.T) {
	testCases := map[string]struct {
		Given    Args
		Expected Args
	}{
		"should remove line break when arg contains line break char": {
			Given:    Args{"a", "b", "c\n", "\nd"},
			Expected: Args{"a", "b", "c", "d"},
		},
		"should remove leading and trailing white space when arg starts with space": {
			Given:    Args{"a", "  b", "  c", "d"},
			Expected: Args{"a", "b", "c", "d"},
		},
		"should remove $ prefix when arg starts with dollar sign": {
			Given:    Args{"$a", "%b", "$c"},
			Expected: Args{"a", "%b", "c"},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			tc.Given.format()
			assert.Equal(t, tc.Expected, tc.Given)
		})
	}
}

func TestArg_IsURL(t *testing.T) {
	testCases := map[string]struct {
		Given    Arg
		Expected bool
	}{
		"should return true when given starts with http://": {
			Given:    "http://localhost:0000",
			Expected: true,
		},
		"should return true when given starts with https://": {
			Given:    "https://localhost:0000",
			Expected: true,
		},
		"should return false when given is not starts with http:// or https://": {
			Given:    "localhost:0000",
			Expected: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.Expected, tc.Given.IsURL())
		})
	}
}

func TestArg_ParseHeader(t *testing.T) {
	expectedKey, expectedValue := "a", "b"
	given := Arg("a:b")
	actualKey, actualValue := given.ParseHeader()
	assert.Equal(t, expectedKey, actualKey)
	assert.Equal(t, expectedValue, actualValue)
}

func TestArg_String(t *testing.T) {
	testCases := map[string]struct {
		Given    *Arg
		Expected string
	}{
		"should return empty string when given arg is nil": {
			Given:    nil,
			Expected: "",
		},
		"should return empty string when given arg is empty string": {
			Given:    newArg(""),
			Expected: "",
		},
		"should return abcd when given arg is abcd": {
			Given:    newArg("abcd"),
			Expected: "abcd",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.Expected, tc.Given.String())
		})
	}
}

func newArg(s string) *Arg {
	arg := Arg(s)
	return &arg
}
