package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	middleSearch(root, &res)
	return res
}

func middleSearch(node *TreeNode, res *[]int) {
	if node.Left != nil {
		middleSearch(node.Left, res)
	}

	*res = append(*res, node.Val)

	if node.Right != nil {
		middleSearch(node.Right, res)
	}
}
