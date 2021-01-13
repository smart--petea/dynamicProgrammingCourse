package fcc

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

	var shortestCombination []int
	for _, candidate := range candidates {
		result := bestSum(targetSum - candidate, candidates, memo)

		if result == nil {
			continue
		}

		result = append(result, targetSum)
		if shortestCombination == nil || len(result) < len(shortestCombination) {
			shortestCombination = result
		}
	}

	memo[targetSum] = shortestCombination
	return shortestCombination
}
