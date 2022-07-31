package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var separator = regexp.MustCompile(`[\s\t\n.,;:!?"]+`)

func Top10(originalText string) []string {
	if len(originalText) == 0 {
		return nil
	}

	words := separator.Split(originalText, -1)

	countedWords, words := countWords(words)
	sort.Slice(words, func(i, j int) bool {
		if countedWords[words[i]] == countedWords[words[j]] {
			return words[i] < words[j]
		} else {
			return countedWords[words[i]] > countedWords[words[j]]
		}
	})

	return words[0:10]
}

func countWords(words []string) (map[string]int, []string) {
	var newWordsSlice []string
	result := map[string]int{}

	for _, word := range words {
		word = strings.ToLower(word)
		// учтем короткое и длинное тире
		if word == "-" || word == "—" {
			continue
		}

		if _, exist := result[word]; exist {
			result[word] = result[word] + 1
		} else {
			result[word] = 1
			newWordsSlice = append(newWordsSlice, word)
		}
	}

	return result, newWordsSlice
}
