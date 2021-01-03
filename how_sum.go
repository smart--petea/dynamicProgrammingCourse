package fcc

func HowSum(targetSum int, numbers []int) []int {
	memo := make(map[int][]int)
	return howSum(targetSum, numbers, memo)
}

func howSum(targetSum int, numbers []int, memo map[int][]int) []int {
	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nil
	}

	if _, ok := memo[targetSum]; ok {
		return memo[targetSum]
	}

	var remainder int
	var remainderResult []int
	memo[targetSum] = nil
	for _, num := range numbers {
		remainder = targetSum - num
		remainderResult = howSum(remainder, numbers, memo)
		if remainderResult != nil {
			remainderResult = append(remainderResult, num)
			memo[targetSum] = remainderResult

			return remainderResult
		}
	}

	return nil
}
