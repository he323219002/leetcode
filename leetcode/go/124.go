package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// 递归
	// 每个节点的路径最大值 = max(左路径+val+右路径) 【左路径，右路径 > 0】
	// 左路径或右路径如何计算
	// max(node.left+val,val,node.right+val)

}

func maxPathSum(root *TreeNode) int {
	// leetcode
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算左右最大贡献值
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		priceNewPath := node.Val + leftGain + rightGain
		maxSum = max(maxSum, priceNewPath)

		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}
