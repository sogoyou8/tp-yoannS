package exo

var memo map[int]int

func Ft_coin(coins []int, amount int) int {
	memo = make(map[int]int)

	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	if val, ok := memo[amount]; ok {
		return val
	}

	result := int(^uint(0) >> 1)

	for _, coin := range coins {
		subproblem := Ft_coin(coins, amount-coin)
		if subproblem != -1 {
			result = min(result, 1+subproblem)
		}
	}

	if result == int(^uint(0)>>1) {
		return -1
	}

	memo[amount] = result

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
