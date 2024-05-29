package main

import "fmt"

func main() {
	nums := []int{1, 3}
	// nums := []int{-10, -3, 0, 5, 9}
	res := sortedArrayToBST(nums)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	// 二分找根 中序递归构建树
	return buildTree(&nums, 0, len(nums)-1)
}

func buildTree(nums *[]int, start int, end int) *TreeNode {
	if start == end {
		return &TreeNode{(*nums)[start], nil, nil}
	}
	mid := (start + end) / 2
	if mid == start {
		node := buildTree(nums, mid, mid)
		rightTree := buildTree(nums, mid+1, end)
		node.Right = rightTree
		return node
	}
	leftTree := buildTree(nums, start, mid-1)
	node := buildTree(nums, mid, mid)
	node.Left = leftTree
	if mid+1 <= end {
		rightTree := buildTree(nums, mid+1, end)
		node.Right = rightTree
	}
	return node
}
