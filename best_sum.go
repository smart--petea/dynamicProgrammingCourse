package fcc

import (
	"sort"
)

func BestSum(targetSum int, candidates []int) []int {
	memo := make(map[int][]int)
	return bestSum(targetSum, candidates, memo)
}

func bestSum(targetSum int, candidates []int, memo map[int][]int) []int {
	if targetSum < 0 {
		return nil
	}

	if targetSum == 0 {
		return []int{}
	}

	if val, ok := memo[targetSum]; ok {
		return val
	}

	results := make(map[int][]int)
	lenResults := []int{}
	for _, candidate := range candidates {
		result := bestSum(targetSum - candidate, candidates, memo)

		if result == nil {
			continue
		}

		result = append(result, targetSum)
		lenResult := len(result)
		results[lenResult] = result
		lenResults = append(lenResults, lenResult)
	}

	if len(results) == 0 {
		memo[targetSum] = nil
		return nil
	}

	sort.Ints(lenResults)

	memo[targetSum] = results[lenResults[0]]
	return memo[targetSum]
}
