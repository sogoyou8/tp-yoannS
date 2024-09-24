package exo

func Ft_profit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		profit := prices[i] - minPrice
		maxProfit = max(maxProfit, profit)
	}

	return maxProfit
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
