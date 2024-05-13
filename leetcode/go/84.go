package main

import (
	"fmt"
)

func main() {
	// heights := []int{4, 2}
	heights := []int{2, 1, 5, 6, 2, 3}
	res := largestRectangleArea(heights)
	fmt.Println(res)
}

func largestRectangleArea(heights []int) int {
	// leetcode
	// 头尾部新增哨兵
	newHeights := []int{-1}
	newHeights = append(newHeights, heights...)
	newHeights = append(newHeights, -1)
	// 向右遍历 遇到严格小的可以确定矩形面积
	stack := make([]int, 0)

	res := 0
	for i := 0; i < len(newHeights); i++ {

		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}

		// 栈中最后一个元素高度
		lastVal := newHeights[stack[len(stack)-1]]
		// 元素大于等于栈最后一个元素 直接入栈
		if newHeights[i] >= lastVal {
			stack = append(stack, i)
			continue
		}

		// 计算左边可以确定面积的矩形
		for newHeights[i] < lastVal {
			// 只有左边严格小于当前 才可以确定面积
			for len(stack) > 0 {
				stack = stack[:len(stack)-1]
				nextLastVal := newHeights[stack[len(stack)-1]]
				if lastVal > nextLastVal {
					board := i - stack[len(stack)-1] - 1
					res = max(res, board*lastVal)
					lastVal = nextLastVal
					break
				}
			}
		}
		stack = append(stack, i)
	}
	return res
}
