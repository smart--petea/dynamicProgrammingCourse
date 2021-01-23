package fcc

import (
	"strings"
)

func CountConstruct(target string, wordsBank []string) int {
	if target == "" {
		return 1
	}

	var counts int
	for _, prefix := range wordsBank {
		if strings.HasPrefix(target, prefix) {
			newTarget := target[len(prefix):]
			counts = counts + CountConstruct(newTarget, wordsBank)
		}
	}

	return counts
}
