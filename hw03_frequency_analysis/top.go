package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	words := strings.Fields(str)

	wordCounter := make(map[string]int)
	for _, w := range words {
		wordCounter[w]++
	}

	topKeys := make([]string, 0, len(wordCounter))
	for key := range wordCounter {
		topKeys = append(topKeys, key)
	}
	sort.Slice(topKeys, func(i, j int) bool {
		if wordCounter[topKeys[i]] == wordCounter[topKeys[j]] {
			return sort.StringsAreSorted([]string{topKeys[i], topKeys[j]})
		}
		return wordCounter[topKeys[i]] > wordCounter[topKeys[j]]
	})

	if len(topKeys) == 0 {
		return nil
	}

	maxSize := wordCounter[topKeys[0]]
	maxCountWord := 0
	for _, v := range topKeys {
		if wordCounter[v] == maxSize {
			maxCountWord++
		}
		if maxCountWord == 10 {
			sort.Strings(topKeys)
			break
		}
	}

	topWords := make([]string, 0, 10)
	for k, v := range topKeys {
		if k == 10 {
			break
		}
		topWords = append(topWords, v)
	}

	return topWords
}
