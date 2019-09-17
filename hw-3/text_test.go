package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCaseType struct {
	description string
	testData string
	expectedResult []string
}

func TestFindWords(t *testing.T) {
	assert := assert.New(t)
	var testCases = []TestCaseType {
		{
			description: "it should work well with usual text",
			testData: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. In ornare, mi nec iaculis semper, augue arcu auctor turpis, hendrerit faucibus tellus nulla non odio. Sed accumsan, urna iaculis interdum pretium, ligula leo feugiat ipsum, in ultricies risus nibh ac purus. Donec scelerisque lacus mi, nec finibus lorem blandit a.`,
			expectedResult: []string{"nec", "iaculis", "Lorem", "sit", "amet,", "consectetur", "adipiscing", "elit.", "In", "ornare,"},
		},
		{
			description: "it should work well with empty string",
			testData: "",
			expectedResult: nil,
		},
		{
			description: "it should work well with small text",
			testData: "one two three four two five",
			expectedResult: []string{"two", "one", "three", "four", "five"},
		},
	}
	for _, testCase := range testCases {
		assert.Equal(testCase.expectedResult, FindWords(testCase.testData), testCase.description)
	}
}
