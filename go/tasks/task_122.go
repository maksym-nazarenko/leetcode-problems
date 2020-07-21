package tasks

func maxProfit(prices []int) int {
	var buy int = 0
	var sell int = 0
	var result int = 0

	n := len(prices)
	for {
		for buy < n-1 && prices[buy] >= prices[buy+1] {
			buy++
		}
		sell = buy + 1
		for sell < n-1 && prices[sell] < prices[sell+1] {
			sell++
		}

		if buy >= n || sell >= n {
			return result
		}

		result += prices[sell] - prices[buy]
		buy = sell + 1
	}
}
