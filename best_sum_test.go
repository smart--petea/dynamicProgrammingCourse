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
			targetSum: 8 ,
			candidates: []int{2, 3, 5},
			expectedLen: 2,
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
