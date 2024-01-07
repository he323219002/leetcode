package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	rightValueDepth := make(map[int]int)
	maxDepth := -1

	nodeQueue := make([]*TreeNode, 0)
	depthQueue := make([]int, 0)

	nodeQueue = append(nodeQueue, root)
	depthQueue = append(depthQueue, 0)

	for len(nodeQueue) != 0 {
		node := nodeQueue[0]
		nodeQueue = nodeQueue[1:]

		depth := depthQueue[0]
		depthQueue = depthQueue[1:]

		if node != nil {
			maxDepth = max(maxDepth, depth)
			rightValueDepth[depth] = node.Val
			nodeQueue = append(nodeQueue, node.Left)
			depthQueue = append(depthQueue, depth+1)

			nodeQueue = append(nodeQueue, node.Right)
			depthQueue = append(depthQueue, depth+1)
		}
	}

	res := make([]int, 0)
	for i := 0; i < maxDepth + 1; i++ {
		res = append(res, rightValueDepth[i])
	}

	return res
}
