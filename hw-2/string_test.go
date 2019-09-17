package string

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseType struct {
	description string
	testData string
	expectedResult string
	expectedError error
}

func TestUnpackString(t *testing.T) {
	assert := assert.New(t)
	var testCases = []TestCaseType {
		{
			description: "it should work well with empty string",
			testData: "",
			expectedResult: "",
		},
		{
			description: "it should unpack simple string correctly",
			testData: "abcd",
			expectedResult: "abcd",
		},
		{
			description: "it should unpack basic string correctly",
			testData: "a2b3c4",
			expectedResult: "aabbbcccc",
		},
		{
			description: "it should unpack basic string with unicode symbols correctly",
			testData: "а2б3в4",
			expectedResult: "аабббвввв",
		},
		{
			description: "it should unpack symbols without number correctly",
			testData: "ab2cd",
			expectedResult: "abbcd",
		},
		{
			description: "it should throw error, when string is incorrect",
			testData: "45",
			expectedError: errors.New("incorrect input string"),
		},
		{
			description: "it should support escaping sequence in basic string",
			testData: `qwe\45`,
			expectedResult: "qwe44444",
		},
		{
			description: "it should support escaping escape symbol in basic string",
			testData: `qwe\\5`,
			expectedResult: `qwe\\\\\`,
		},
	}
	for _, testCase := range testCases {
		result, err := UnpackString(testCase.testData)
		if testCase.expectedError != nil {
			assert.Equal(testCase.expectedError, err, testCase.description)
		} else {
			assert.Equal(testCase.expectedResult, result, testCase.description)
		}
	}
}
