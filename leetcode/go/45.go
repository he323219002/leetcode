package main

import "fmt"

func main() {
	nums := []int{2, 3, 1, 1, 4}
	res := jump(nums)
	fmt.Println(res)
}

func jump(nums []int) int {
	// leetcode
	// 不断更新最远距离
	end, res, maxPos := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		// 当前可以跳到的最大索引位置
		maxPos = max(maxPos, nums[i]+i)
		if i == end {
			end = maxPos
			res += 1
		}
	}
	return res
}
