package leetcode

func MaxProfit(prices []int) int {
	min := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		current := prices[i]
		if current < min {
			min = current
		} else {
			profit := current - min
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}

	return maxProfit
}
