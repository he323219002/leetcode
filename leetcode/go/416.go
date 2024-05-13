package main

import (
	"fmt"
)

func main() {
	nums := []int{14, 9, 8, 4, 3, 2}
	res := canPartition(nums)
	fmt.Println(res)
}

func sumArray(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func canPartition(nums []int) bool {
	// gpt
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// 如果总和为奇数，无法分割成两个相等的子集
	if totalSum%2 != 0 {
		return false
	}

	targetSum := totalSum / 2
	dp := make([]bool, targetSum+1)
	dp[0] = true

	for _, num := range nums {
		for j := targetSum; j >= num; j-- {
			if dp[targetSum] {
				return true
			}
			// 上一层循环的结果如果为true则跳过
			dp[j] = dp[j] || dp[j-num]
		}
	}

	return dp[targetSum]
}
