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
	}

	for _, test := range tests {
		outputCount := CountConstruct(test.target, test.wordsBank)
		if outputCount != test.expectedCount {
			t.Errorf("(%s, %+v)->%d, expected %d", test.target, test.wordsBank, outputCount, test.expectedCount)
		}
	}
}
