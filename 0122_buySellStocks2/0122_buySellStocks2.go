package buySellStocks2

// sliding window solution O(N) time complexity
// Strategy is to be greedy on sells if the price ever falls below a local maxima then we should sell on the day prior
// and capture all the gains up until that point and then set the buy price to the day after the local maxima. Repeat
// until the end of the array and make sure to capture the last available profit by checking if sell > buy after the
// loop exits.
func maxProfit(prices []int) int {
	profit := 0
	buy := prices[0]
	sell := 0
	for i := 1; i < len(prices); i++ {
		price := prices[i]
		// price is higher. we should grab the new local maximum
		if price > sell {
			sell = price
		}
		// price is below the sell price so the sell we saved is a local maximum and we should capture profits
		if price < sell {
			// only capture if the sell is greater than the buy price
			if sell > buy {
				profit += sell - buy
			}
			// new buy price is the first price after the local maximum captured and erase the sell price
			buy = price
			sell = 0
		}

		// better buy price. use it
		if price < buy {
			buy = price
		}
	}

	// after the loop exits capture the final potential profit
	if sell > buy {
		profit += sell - buy
	}

	return profit
}
