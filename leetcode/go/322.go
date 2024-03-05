package main

import (
	"fmt"
	"math"
)

func main() {
	coins := []int{2}
	amount := 11
	res := coinChange(coins, amount)
	fmt.Println(res)
}

func coinChange(coins []int, amount int) int {
	// leetcode
	d := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		d[i] = math.MaxInt32
	}
	d[0] = 0

	for i := 1; i < amount+1; i++ {
		for _, coin := range coins {
			if coin <= i {
				d[i] = min(d[i], d[i-coin]+1)
			}
		}
	}

	if d[amount] == math.MaxInt32 {
		return -1
	}

	return d[amount]
}
