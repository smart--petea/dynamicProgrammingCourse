package fcc

import (
	"strings"
)

func CanConstruct(target string, wordBank []string) bool {
	if len(target) == 0 {
		return true
	}

	for _, prefix := range wordBank {
		if strings.HasPrefix(target, prefix) {
			targetWithoutPrefix := target[len(prefix):]
			if CanConstruct(targetWithoutPrefix, wordBank) {
				return true
			}
		}
	}

	return false
}
