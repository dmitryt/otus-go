package string

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpackString(t *testing.T) {
	assert := assert.New(t)
	getResult := func(value string) string {
		result, _ := UnpackString(value)
		return result
	}
	getError := func(value string) error {
		_, err := UnpackString(value)
		return err
	}
	assert.Equal("", getResult(""), "it should work well with empty string")
	assert.Equal("abcd", getResult("abcd"), "it should unpack simple string correctly")
	assert.Equal("aabbbcccc", getResult("a2b3c4"), "it should unpack basic string correctly")
	assert.Equal("аабббвввв", getResult("а2б3в4"), "it should unpack basic string with unicode symbols correctly")
	assert.Equal("abbcd", getResult("ab2cd"), "it should unpack symbols without number correctly")
	assert.Equal(errors.New("Incorrect input string"), getError("45"), "it should throw error, when string is incorrect")

	assert.Equal("qwe44444", getResult(`qwe\45`), "it should support escape sequence in basic string")
	assert.Equal(`qwe\\\\\`, getResult(`qwe\\5`), "it should support escape sequence in basic string")
}
