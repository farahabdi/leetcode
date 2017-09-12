package buySellStock

func maxProfit(prices []int) int {
	//nums := []int{3, 4, 1, 8, 3, 5}

	maxDiff := 0

	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {

			if prices[j]-prices[i] > maxDiff {
				maxDiff = prices[j] - prices[i]
			}

		}
	}

	return maxDiff
}
