package buySellStocks

// sliding window solution O(N)
func maxProfit(prices []int) int {
	profit := 0
	buy := prices[0]
	for i := 1; i < len(prices); i++ {
		price := prices[i]
		// if the price gives us a profit, check to see if it's the best possible profit and save if it is
		if price > buy {
			profit = max(profit, price-buy)
		} else {
			// new low price found. save it
			buy = price
		}
	}

	return profit
}

// brute force solution O(N^2)
func maxProfitBrute(prices []int) int {
	max := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i; j < len(prices); j++ {
			diff := prices[j] - prices[i]
			if diff > max {
				max = diff
			}
		}
	}

	return max
}
