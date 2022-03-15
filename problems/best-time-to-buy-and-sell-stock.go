package problems

// 从前向后遍历 保留最小的数字，并记录走到当前的数值的时候，记录最大的利润

func maxProfit(prices []int) int {

	maxProfit := 0
	minPrice := prices[0]

	for _, val := range prices {
		if val < minPrice {
			minPrice = val
		} else if maxProfit < val-minPrice {
			maxProfit = val - minPrice
		}
	}

	return maxProfit
}
