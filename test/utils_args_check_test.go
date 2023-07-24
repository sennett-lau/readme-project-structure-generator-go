package tests

import (
	"sennett-lau/rpsg/utils"
	"testing"
)

type argsCheckTestCase struct {
	input    string
	expected bool
}

func TestArgIsValidExtendIgnoreList(t *testing.T) {
	testCases := []argsCheckTestCase{
		// Valid test cases
		{input: "--extend-ignore-list=image.jpg", expected: true},
		{input: "--extend-ignore-list=image.jpg,test.js", expected: true},
		{input: "--extend-ignore-list=image.jpg,test.js,react.jsx", expected: true},

		// Invalid test cases
		{input: "--extend-ignore-list=invalid$", expected: false},
		{input: "--extend-ignore-list=", expected: false},
		{input: "--extend-ignore-list=value,invalid$", expected: false},
		{input: "--extend-ignore-list=value1/value2,invalid$", expected: false},
		{input: "--extend-ignore-list=value1/invalid$", expected: false},
		{input: "--extend-ignore-list=value1, value2", expected: false},
		{input: "key=value1", expected: false},
	}

	for _, tc := range testCases {
		result := utils.ArgIsValidExtendIgnoreList(tc.input)
		if result != tc.expected {
			t.Errorf("Input: %s, Expected: %v, Got: %v", tc.input, tc.expected, result)
		}
	}
}
