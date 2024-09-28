package main

import "fmt"

func main() {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	res := findPeakElement(nums)
	fmt.Println(res)
}

func findPeakElement(nums []int) int {
	// 如果不考虑时间复杂度要求，可以达到On
	// 1.标记值：前两个元素是否递增，2.当前值是否小于前一个值

	// leetcode
	// 最左和最右题设为负无穷，则沿最左一定存在峰值
	// 双指针 l,r 若n[m] > n[m+1]则左侧存在峰值，反之为右侧
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
