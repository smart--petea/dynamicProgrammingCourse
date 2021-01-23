package fcc

import (
	"strings"
)

func CountConstruct(target string, wordsBank []string) int {
	memo := make(map[string]int)

	return countConstruct(target, wordsBank, memo)
}

func countConstruct(target string, wordsBank []string, memo map[string]int) int {
	if target == "" {
		return 1
	}

	if _, ok := memo[target]; ok {
		return memo[target]
	}

	var counts int
	for _, prefix := range wordsBank {
		if strings.HasPrefix(target, prefix) {
			newTarget := target[len(prefix):]
			counts = counts + countConstruct(newTarget, wordsBank, memo)
		}
	}

	memo[target] = counts
	return counts
}
