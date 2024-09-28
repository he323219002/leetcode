package main

import "fmt"

func main() {
	e := ListNode{5, nil}
	d := ListNode{4, &e}
	c := ListNode{3, &d}
	b := ListNode{2, &c}
	a := ListNode{1, &b}
	res := oddEvenList(&a)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 双指针

	cur, nodeA := head, head
	for nodeA != nil {
		tempNode := nodeA
		if nodeA.Next == nil {
			cur.Next = tempNode
			return cur
		}
		nodeB := nodeA.Next
		tempNode.Next = nodeB
		cur.Next = tempNode

		tempNode = tempNode.Next
		cur = tempNode

		nodeA = nodeB.Next
		if nodeA != nil {
			nodeB = nodeA.Next
		}
	}
	return cur
}
