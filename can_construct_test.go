package fcc

import (
	"testing"
)

func TestCanConstruct(t *testing.T) {
	tests := []struct{
		target string
		wordBank []string
		expected bool
	}{
		{
			target: "abcdef",
			wordBank: []string{"ab", "abc", "cd", "def", "abcd"},
			expected: true,
		},
		{
			target: "skateboard",
			wordBank: []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"},
			expected: false,
		},
		{
			target: "enterapotentpot",
			wordBank: []string{"a", "p", "ent", "enter", "ot", "o", "t"},
			expected: true,
		},
		{
			target: "eeeeeeeeeeeeeeeeeeeeeeeeeeeeef",
			wordBank: []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"},
			expected: false,
		},
	}

	for _, test := range tests {
		result := CanConstruct(test.target, test.wordBank)
		if result != test.expected {
			t.Errorf("(%+v,%+v)->%+v, expected +%v", test.target, test.wordBank, result, test.expected)
		}
	}
}
