package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// dfs

}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	curDeep := 1
	res := dfs(root, curDeep)
	return res
}

func dfs(node *TreeNode, curDeep int) int {
	if node.Left == nil && node.Right == nil {
		return curDeep
	}

	if node.Left == nil {
		return dfs(node.Right, curDeep+1)
	} else if node.Right == nil {
		return dfs(node.Left, curDeep+1)
	}

	return max(dfs(node.Left, curDeep+1), dfs(node.Right, curDeep+1))
}
