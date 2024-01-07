package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉搜索树-中序遍历-递增,*TreeNode 是指向TreeNode的指针，而**TreeNode是指向*TreeNode的指针
func dfs(node *TreeNode, count *int, res **TreeNode) {
	if node == nil {
		return
	}

	// 遍历左
	dfs(node.Left, count, res)
	if *res == nil {
		// 根节点
		*count -= 1
		if *count == 0 {
			// 遍历到第k小
			*res = node
			return
		}
	}

	// 遍历右
	dfs(node.Right, count, res)
	return
}

func kthSmallest(root *TreeNode, k int) int {
	var res *TreeNode
	dfs(root, &k, &res)
	return res.Val
}
func main() {
	d := TreeNode{2, nil, nil}
	c := TreeNode{4, nil, nil}
	b := TreeNode{1, nil, &d}
	a := TreeNode{3, &b, &c}

	kthSmallest(&a, 2)

}
