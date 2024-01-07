package main

func main() {
	d := &ListNode{1, nil}
	c := &ListNode{2, d}
	b := &ListNode{2, c}
	a := &ListNode{1, b}

	isPalindrome(a)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// 将值放入列表和列表到序比较
	stack := make([]int, 0)
	for node := head; node != nil; node = node.Next {
		stack = append(stack, node.Val)
	}

	stackRev := make([]int, len(stack))
	for i, j := 0, len(stack)-1; i < len(stack); i, j = i+1, j-1 {
		stackRev[i] = stack[j]
	}

	for i := 0; i < len(stack); i++ {
		if stack[i] != stackRev[i] {
			return false
		}
	}
	return true
}
