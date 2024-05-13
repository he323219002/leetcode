package main

import (
	"math/rand"
	"time"
)

func main() {
	// nums := []int{3, 2, 1, 5, 6, 4}
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	findKthLargest(nums, 4)
}

func findKthLargest(nums []int, k int) int {
	// leetcode 答案
	res := quickSelect(&nums, k)
	return res
}

func quickSelect(nums *[]int, k int) int {
	// 随机选择一个数
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	random := (*nums)[seed.Int()%len(*nums)]

	// 分为3组
	lt, eq, gt := make([]int, 0), make([]int, 0), make([]int, 0)
	for _, val := range *nums {
		if val < random {
			lt = append(lt, val)
		} else if val == random {
			eq = append(eq, val)
		} else {
			gt = append(gt, val)
		}
	}
	// 如果在大于的组中
	if k <= len(gt) {
		return quickSelect(&gt, k)
	} else if k > len(gt)+len(eq) {
		return quickSelect(&lt, k-len(eq)-len(gt))
	} else {
		return random
	}
}
