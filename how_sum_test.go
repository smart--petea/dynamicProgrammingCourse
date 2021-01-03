package fcc

import (
	"testing"
)

type TestEntryLL struct {
	InputTargetSum  int
	InputNumbers    []int
	ExpectedNumbers []int
}

func TestHowSum(t *testing.T) {
	testEntries := []TestEntryLL{
		TestEntryLL{
			InputTargetSum:  7,
			InputNumbers:    []int{2, 3},
			ExpectedNumbers: []int{3, 2, 2},
		},
		TestEntryLL{
			InputTargetSum:  7,
			InputNumbers:    []int{5, 3, 4, 7},
			ExpectedNumbers: []int{4, 3},
		},
		TestEntryLL{
			InputTargetSum:  7,
			InputNumbers:    []int{2, 4},
			ExpectedNumbers: nil,
		},
		TestEntryLL{
			InputTargetSum:  8,
			InputNumbers:    []int{2, 3, 5},
			ExpectedNumbers: []int{2, 2, 2, 2},
		},
		TestEntryLL{
			InputTargetSum:  300,
			InputNumbers:    []int{7, 14},
			ExpectedNumbers: nil,
		},
	}

	var outputNumbers []int
	for _, testEntry := range testEntries {
		outputNumbers = HowSum(testEntry.InputTargetSum, testEntry.InputNumbers)
		if len(outputNumbers) != len(testEntry.ExpectedNumbers) {
			t.Errorf("(%d, %v)->%v expected %v", testEntry.InputTargetSum, testEntry.InputNumbers, outputNumbers, testEntry.ExpectedNumbers)
		}

		for i := range testEntry.ExpectedNumbers {
			if testEntry.ExpectedNumbers[i] != outputNumbers[i] {
				t.Errorf("(%d, %v)->%v expected %v", testEntry.InputTargetSum, testEntry.InputNumbers, outputNumbers, testEntry.ExpectedNumbers)
			}
		}
	}
}
