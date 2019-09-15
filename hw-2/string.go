package string

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

// UnpackString - unpacks the string by repeating particular symbols N times
// Example: UnpackString("a2b3c4") => aabbbcccc
func UnpackString(str string) (string, error) {
	var result string
	var lastSymbol string
	var counter int
	var isEscaped bool
	for _, charCode := range str {
		// escape symbol
		if charCode == 92 && !isEscaped {
			result += lastSymbol
			isEscaped = true
			continue
		}
		if isEscaped {
			lastSymbol = string(charCode)
			isEscaped = false
			continue
		}
		if unicode.IsDigit(charCode) {
			counter, _ = strconv.Atoi(string(charCode))
			if lastSymbol == "" {
				return "", errors.New("Incorrect input string")
			}
		} else {
			result += lastSymbol
			lastSymbol = string(charCode)
		}
		if counter > 0 {
			result += strings.Repeat(lastSymbol, counter)
			lastSymbol = ""
			counter = 0
		}
	}
	return result + lastSymbol, nil
}
