package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	d := TreeNode{2, nil, nil}
	c := TreeNode{4, nil, nil}
	b := TreeNode{1, nil, &d}
	a := TreeNode{3, &b, &c}

	//   3
	// 1    4
	//  2

	res := deleteNode(&a, 3)
	fmt.Println(res)
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	// 遍历
	preSort(&root, key)
	return root
}

func preSort(node **TreeNode, key int) {
	if *node == nil {
		return
	}
	if (*node).Val < key {
		preSort(&(*node).Right, key)
	} else if (*node).Val > key {
		preSort(&(*node).Left, key)
	} else {
		if (*node).Left == nil && (*node).Right == nil {
			*node = nil
			return
		}
		if (*node).Left == nil {
			*node = (*node).Right
			return
		}
		if (*node).Right == nil {
			*node = (*node).Left
			return
		}
		// 把左子树接上，右子树接到左子树的右子树底部
		leftTempNode := (*node).Left
		for leftTempNode.Right != nil {
			leftTempNode = leftTempNode.Right
		}
		leftTempNode.Right = (*node).Right
		*node = (*node).Left
	}
	fmt.Println(*node)
}
