package tests

import (
	"sennett-lau/rpsg/utils"
	"testing"
)

func TestArgIsValidExtendIgnoreList(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		// Valid test cases
		{input: "key=image.jpg", expected: true},
		{input: "key=image.jpg,test.js", expected: true},
		{input: "key=image.jpg,test.js,react.jsx", expected: true},

		// Invalid test cases
		{input: "key=invalid$", expected: false},
		{input: "key=", expected: false},
		{input: "key=value,invalid$", expected: false},
		{input: "key=value1/value2,invalid$", expected: false},
		{input: "key=value1/invalid$", expected: false},
		{input: "key=value1, value2", expected: false},
	}

	for _, tc := range testCases {
		result := utils.ArgIsValidExtendIgnoreList(tc.input)
		if result != tc.expected {
			t.Errorf("Input: %s, Expected: %v, Got: %v", tc.input, tc.expected, result)
		}
	}
}
