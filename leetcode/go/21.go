package main

import "fmt"

func main() {
	a := ListNode{4, nil}
	b := ListNode{2, &a}
	c := ListNode{1, &b}

	d := ListNode{4, nil}
	e := ListNode{3, &d}
	f := ListNode{1, &e}

	res := mergeTwoLists(&c, &f)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curNode := dummy

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			curNode.Next = list2
			list2 = list2.Next
		} else {
			curNode.Next = list1
			list1 = list1.Next
		}
		curNode = curNode.Next
	}
	if list1 == nil {
		curNode.Next = list2
	}
	if list2 == nil {
		curNode.Next = list1
	}
	return dummy.Next
}
