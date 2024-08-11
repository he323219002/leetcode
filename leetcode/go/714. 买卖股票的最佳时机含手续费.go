package main

import "fmt"

func main() {
	// prices := []int{1, 3, 2, 8, 4, 9}
	prices := []int{1, 3, 7, 5, 10, 3}
	fee := 3
	res := maxProfit(prices, fee)
	fmt.Println(res)
}

func maxProfit(prices []int, fee int) int {
	// dp[i][0] 表示第i天不持有的最大利润
	// dp[i][1] 表示第i天持有的最大利润
	// dp[i][0] = max(dp[i-1][0],dp[i-1][1] + price[i] - 2)
	// dp[i][1] = max(dp[i-1][0] - price[i],dp[i-1][1])
	// dp[0][0] = 0
	// dp[0][1] = -price[0]

	// days := len(prices)
	dp := make([][]int, 2)

	for day, price := range prices {
		cur := make([]int, 2)
		if day == 0 {
			cur[0], cur[1] = 0, -price
			dp[1] = cur
			continue
		}
		dp[0] = dp[1]

		cur[0] = max(dp[0][0], dp[0][1]+price-fee)
		cur[1] = max(dp[0][0]-price, dp[0][1])

		dp[1] = cur
	}

	return max(dp[1][0], dp[1][1])
}
