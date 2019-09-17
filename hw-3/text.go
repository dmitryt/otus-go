package text

import (
	"regexp"
	"sort"
)

// FindWords - find 10 the most used words in incoming text
func FindWords(txt string) []string {
	var result []string
	if txt == "" {
		return nil
	}
	words := regexp.MustCompile(`\s+`).Split(txt, -1)
	hash := make(map[string]int)
	for _, word := range words {
		hash[word]++
		if hash[word] == 1 {
			result = append(result, word)
		}
	}
	sort.Slice(result, func (i, j int) bool {
		return hash[result[i]] > hash[result[j]]
	})
	if len(result) > 10 {
		return result[:10]
	}
	return result
}
