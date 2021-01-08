package fcc

import (
	"sort"
	"fmt"
)

func BestSum(targetSum int, candidates []int) []int {
	if targetSum < 0 {
		return nil
	}

	if targetSum == 0 {
		return []int{}
	}

	results := make(map[int][]int)
	lenResults := []int{}
	for _, candidate := range candidates {
		result := BestSum(targetSum - candidate, candidates)

		if result == nil {
			continue
		}

		result = append(result, targetSum)
		lenResult := len(result)
		results[lenResult] = result
		lenResults = append(lenResults, lenResult)
	}

	if len(results) == 0 {
		return nil
	}

	sort.Ints(lenResults)

	return results[lenResults[0]] 
}
