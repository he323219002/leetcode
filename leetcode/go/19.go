package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := ListNode{0, head}
	stack := make([]*ListNode, 0)
	for {
		if head == nil {
			break
		}
		stack = append(stack, head)
		head = head.Next
	}
	nNode := stack[len(stack)-n]
	stack = stack[:len(stack)-n]

	if len(stack) == 0 {
		return nNode.Next
	}

	preNnode := stack[len(stack)-1]
	preNnode.Next = nNode.Next
	return dummy.Next

}

func main() {
	e := ListNode{5, nil}
	d := ListNode{4, &e}
	c := ListNode{3, &d}
	b := ListNode{2, &c}
	a := ListNode{1, &b}

	removeNthFromEnd(&a, 2)
}
