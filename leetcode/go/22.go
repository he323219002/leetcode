package main

import "fmt"

func main() {
	generateParenthesis(3)
}

type Node struct {
	left   int
	right  int
	curStr string
}

func bfs(node *Node, queue *[]Node, res *[]string) {
	if node.left == 0 && node.right == 0 {
		*res = append(*res, node.curStr)
		return
	}

	if node.left > node.right {
		return
	}

	// 加左括号
	if node.left > 0 {
		leftNode := Node{node.left - 1, node.right, node.curStr + "("}
		*queue = append(*queue, leftNode)
	}

	// 加右括号
	if node.right > 0 {
		rightNode := Node{node.left, node.right - 1, node.curStr + ")"}
		*queue = append(*queue, rightNode)
	}

}

func generateParenthesis(n int) []string {
	queue := make([]Node, 0)
	res := make([]string, 0)
	// 初始化第一个
	root := Node{n, n, ""}
	queue = append(queue, root)
	for len(queue) != 0 {
		curNode := queue[0]
		queue = queue[1:]
		bfs(&curNode, &queue, &res)
	}
	fmt.Println(res)
	return res
}

// // leetcode
// func dfs(curStr string, left int, right int, res *[]string) {
// 	if left == 0 && right == 0 {
// 		*res = append(*res, curStr)
// 		return
// 	}

// 	// 用的右括号多，剩余左括号大于右括号 剪枝
// 	if left > right {
// 		return
// 	}

// 	if left > 0 {
// 		dfs(curStr+"(", left-1, right, res)
// 	}

// 	if right > 0 {
// 		dfs(curStr+")", left, right-1, res)
// 	}
// }

// func generateParenthesis(n int) []string {
// 	res := make([]string, 0)
// 	if n == 0 {
// 		return res
// 	}

// 	dfs("", n, n, &res)
// 	return res
// }
