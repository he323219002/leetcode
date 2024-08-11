package main

import "math"

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

	maxLevelSum(&a)
}

func maxLevelSum(root *TreeNode) int {
	nodeMap := make(map[int]int, 0)
	nodeMap[0] = root.Val
	dfs(root, &nodeMap, 0)
	res := findMaxValueKey(nodeMap)
	return res
}

func dfs(node *TreeNode, resultMap *map[int]int, deep int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		(*resultMap)[deep+1] += node.Left.Val
		dfs(node.Left, resultMap, deep+1)
	}
	if node.Right != nil {
		(*resultMap)[deep+1] += node.Right.Val
		dfs(node.Right, resultMap, deep+1)
	}
}

func findMaxValueKey(m map[int]int) int {
	val := math.MinInt64
	res := 0
	for k, v := range m {
		if v > val {
			res = k
			val = v
		} else if v == val {
			res = min(res, k)
		}
	}
	return res
}
