package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpackString(t *testing.T) {
	var expected []string
	assert := assert.New(t)
	expected = []string{"nec", "iaculis", "Lorem", "sit", "amet,", "consectetur", "adipiscing", "elit.", "In", "ornare,"}
	const text1 = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. In ornare, mi nec iaculis semper, augue arcu auctor turpis, hendrerit faucibus tellus nulla non odio. Sed accumsan, urna iaculis interdum pretium, ligula leo feugiat ipsum, in ultricies risus nibh ac purus. Donec scelerisque lacus mi, nec finibus lorem blandit a.`;
	assert.Equal(expected, FindWords(text1), "it should work well with usual text")
	assert.Equal([]string{}, FindWords(""), "it should work well with empty string")
	assert.Equal([]string{"two", "one", "three", "four", "five"}, FindWords("one two three four two five"), "it should work well with small text")
}
