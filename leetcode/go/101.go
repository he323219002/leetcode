package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return bfs(root.Left, root.Right)
}

func bfs(leftNode *TreeNode, rightNode *TreeNode) bool {
	if leftNode == nil && rightNode == nil {
		return true
	}
	if leftNode == nil || rightNode == nil {
		return false
	}
	if leftNode.Val != rightNode.Val {
		return false
	}
	leftRes := bfs(leftNode.Left, rightNode.Right)
	rightRes := bfs(leftNode.Right, rightNode.Left)
	return leftRes && rightRes
}
