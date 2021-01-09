package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"regexp"
	"sort"
	"strings"
)

func Top10(text string) []string {
	reg := regexp.MustCompile(`[!"',. \t\v\r\n\f]+`)
	textWords := reg.Split(text, -1)

	wordStatMap := map[string]int{}

	for _, word := range textWords {
		if word == "" || word == "-" {
			continue
		}
		wordStatMap[strings.ToLower(word)]++
	}

	wordStatSlice := make([]string, 0, len(wordStatMap))

	for word := range wordStatMap {
		wordStatSlice = append(wordStatSlice, word)
	}

	sort.Slice(wordStatSlice, func(i, j int) bool {
		return wordStatMap[wordStatSlice[i]] > wordStatMap[wordStatSlice[j]]
	})

	if len(wordStatSlice) > 10 {
		wordStatSlice = wordStatSlice[:10]
	}

	return wordStatSlice
}
