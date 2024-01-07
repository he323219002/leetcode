package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	temp_left := maxSimiDiameter(root.Left) + 1

	temp_right := maxSimiDiameter(root.Right) + 1

	return temp_left + temp_right
}

// 最大半径
func maxSimiDiameter(node *TreeNode) int {
	if node == nil {
		return 0
	}

	temp_left := maxSimiDiameter(node.Left)
	temp_right := maxSimiDiameter(node.Right)

	return max(temp_left, temp_right) + 1
}
