package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// 考虑用双指针或者递归算法，不开辟额外空间
	// 先进后出用栈
	stack := make([]*ListNode, 0)
	for node := head; node != nil; node = node.Next {
		stack = append(stack, &ListNode{node.Val, nil})
	}
	dummy := &ListNode{}
	for node := dummy; len(stack) != 0; {
		node.Next = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Next
	}
	return dummy.Next
}

func main() {
	e := ListNode{5, nil}
	d := ListNode{4, &e}
	c := ListNode{3, &d}
	b := ListNode{2, &c}
	a := ListNode{1, &b}

	reverseList(&a)
}
