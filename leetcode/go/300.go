package main

import "fmt"

func main() {
	// nums := []int{0, 1, 0, 3, 2, 3}
	// nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	// nums := []int{4, 10, 4, 3, 8, 9}
	nums := []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	res := lengthOfLIS(nums)
	fmt.Println(res)
}

func lengthOfLIS(nums []int) int {
	// 前面比他小的最大值+1
	// d[i]表示从0到i最长递增 d[0] = 1, d[i] = d[i-1]+1 if d[i] > d[i-1] else d[i] = d[n]+1 if d[i]>d[n]
	d := make([]int, len(nums))
	// 初始化所有为1
	for i := 0; i < len(nums); i++ {
		d[i] = 1
	}
	// 找到第一个就返回
	res := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				d[i] = max(d[i], d[j]+1)
				res = max(res, d[i])
			}
		}
	}
	return res
}
