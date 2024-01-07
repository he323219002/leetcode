package main

func main() {
	e := ListNode{0, nil}
	d := ListNode{4, &e}
	c := ListNode{3, &d}
	b := ListNode{5, &c}
	a := ListNode{-1, &b}
	sortList(&a)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode 自下而上归并排序
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	res := sortFunc(head, nil)
	return res
}

func sortFunc(head *ListNode, tail *ListNode) *ListNode {
	// 递归终止条件

}

func mergeFunc(node1 *ListNode, node2 *ListNode) *ListNode {
	dummy := ListNode{0, nil}
	temp, temp1, temp2 := dummy, node1, node2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = node2
			temp2 = temp2.Next
		}
		temp = *temp.Next
	}

	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}

	return dummy.Next

}
