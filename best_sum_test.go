package fcc

import (
	"testing"
)

func TestBestSum(t *testing.T) {
	tests := []struct{
		targetSum int 
		candidates []int
		expectedLen int
	}{
		{
			targetSum: 7 ,
			candidates: []int{5, 3, 4, 7},
			expectedLen: 1,
		},
		{
			targetSum: 8 ,
			candidates: []int{2, 3, 5},
			expectedLen: 2,
		},
		{
			targetSum: 8 ,
			candidates: []int{1, 4, 5},
			expectedLen: 2,
		},
		{
			targetSum: 100 ,
			candidates: []int{1, 2, 5, 25},
			expectedLen: 4,
		},
	}

	for _, test := range tests {
		output := BestSum(test.targetSum, test.candidates)

		var outputLen int
		if output == nil {
			outputLen = 0
		} else {
			outputLen = len(output)
		}

		if outputLen != test.expectedLen {
			t.Errorf("(%v,%v) got %v with len %v, expected %v", test.targetSum, test.candidates, output, outputLen, test.expectedLen)
		}
	}
}
