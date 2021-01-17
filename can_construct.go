package fcc

import (
	"strings"
)

func CanConstruct(target string, wordBank []string) bool {
	memo := make(map[string]bool)
	return canConstruct(target, wordBank, memo)
}

func canConstruct(target string, wordBank []string, memo map[string]bool) bool {
	if len(target) == 0 {
		return true
	}

	if val, ok := memo[target]; ok {
		return val
	}


	result := false
	for _, prefix := range wordBank {
		if strings.HasPrefix(target, prefix) {
			targetWithoutPrefix := target[len(prefix):]
			if canConstruct(targetWithoutPrefix, wordBank, memo) {
				result = true
			}
		}
	}

	memo[target] = result
	return result
}
