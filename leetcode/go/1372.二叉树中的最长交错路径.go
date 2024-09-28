package main

func main() {
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	// leetcode DFS,left right 为左右路径
	res := 0
	var dfs func(node *TreeNode, left int, right int)
	dfs = func(node *TreeNode, left int, right int) {
		res = max(res, left, right)
		// 向左遍历，下一个左路径 是当前右路径 + 1 ，左路径归零
		if node.Left != nil {
			dfs(node.Left, right+1, 0)
		}
		if node.Right != nil {
			dfs(node.Right, 0, left+1)
		}
	}
	dfs(root, 0, 0)
	return res
}
