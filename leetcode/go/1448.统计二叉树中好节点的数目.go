package main

func main() {
	d := TreeNode{2, nil, nil}
	c := TreeNode{4, nil, nil}
	b := TreeNode{1, nil, &d}
	a := TreeNode{3, &b, &c}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	// 用一个标记遍历路径 记录最大值
	res := 0

	var bfs func(node *TreeNode, tempMax int)
	bfs = func(node *TreeNode, tempMax int) {
		if node.Val >= tempMax {
			res += 1
			tempMax = node.Val
		}
		if node.Left != nil {
			bfs(node.Left, tempMax)
		}
		if node.Right != nil {
			bfs(node.Right, tempMax)
		}
	}

	bfs(root, root.Val)
	return res
}
