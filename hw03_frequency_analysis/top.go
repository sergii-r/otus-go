package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"regexp"
	"sort"
	"strings"
)

var reg = regexp.MustCompile(`[\p{L}\d]+-?[\p{L}\d]*`)

func Top10(text string) []string {
	textWords := reg.FindAllString(text, -1)

	wordStatMap := map[string]int{}
	wordStatSlice := make([]string, 0)

	for _, word := range textWords {
		word = strings.ToLower(word)
		if _, ok := wordStatMap[word]; !ok {
			wordStatSlice = append(wordStatSlice, word)
		}
		wordStatMap[word]++
	}

	sort.Slice(wordStatSlice, func(i, j int) bool {
		return wordStatMap[wordStatSlice[i]] > wordStatMap[wordStatSlice[j]]
	})

	if len(wordStatSlice) > 10 {
		wordStatSlice = wordStatSlice[:10]
	}

	return wordStatSlice
}
