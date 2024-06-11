package main

import (
	"fmt"
	"math"
)

func main() {
	// prices := []int{7, 1, 5, 3, 6, 4}
	prices := []int{7, 6, 4, 3, 1}
	res := maxProfit(prices)
	fmt.Println(res)
}

func maxProfit(prices []int) int {
	// leetcode
	cost, result := math.MaxInt64, 0
	for _, price := range prices {
		if price < cost {
			cost = price
			continue
		}
		result = max(result, price-cost)
	}
	return result
}
