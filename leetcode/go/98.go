package main

import "fmt"

func main() {
	a := TreeNode{1, nil, nil}
	b := TreeNode{3, nil, nil}
	c := TreeNode{2, &a, &b}

	res := isValidBST(&c)

	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	li := make([]int, 0)
	middleSearch(root, &li)
	return checkOrder(li)
}

// 检查数组是否严格递增
func checkOrder(li []int) bool {
	if len(li) == 1 {
		return true
	}
	for i := 1; i < len(li); i++ {
		if li[i] <= li[i-1] {
			return false
		}
	}

	return true
}

// 中序遍历
func middleSearch(node *TreeNode, li *[]int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		middleSearch(node.Left, li)
	}
	*li = append(*li, node.Val)
	if node.Right != nil {
		middleSearch(node.Right, li)
	}
}
