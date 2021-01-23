package fcc

import (
	"testing"
)

func TestCountConstruct(t *testing.T) {
	tests := []struct {
		wordsBank     []string
		target        string
		expectedCount int
	}{
		{
			wordsBank:     []string{"ab", "abc", "cd", "def", "abcd"},
			target:        "abcdef",
			expectedCount: 1,
		},
		{
			wordsBank:     []string{"purp", "p", "ur", "le", "purpl"},
			target:        "purple",
			expectedCount: 2,
		},
		{
			wordsBank:     []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"},
			target:        "skateboard",
			expectedCount: 0,
		},
		{
			wordsBank:     []string{"a", "p", "ent", "enter", "ot", "o", "t"},
			target:        "enterapotentpot",
			expectedCount: 4,
		},
		{
			wordsBank:     []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"},
			target:        "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef",
			expectedCount: 0,
		},
	}

	for _, test := range tests {
		outputCount := CountConstruct(test.target, test.wordsBank)
		if outputCount != test.expectedCount {
			t.Errorf("(%s, %+v)->%d, expected %d", test.target, test.wordsBank, outputCount, test.expectedCount)
		}
	}
}
