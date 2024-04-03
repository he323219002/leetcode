package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

func mergeTwoList(node1 *ListNode, node2 *ListNode) *ListNode {
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

func merge(lists []*ListNode, left int, right int) *ListNode {
	if left == right {
		return lists[left]
	}

	mid := (left + right) / 2
	leftList := merge(lists, left, mid)
	rightList := merge(lists, mid+1, right)
	return mergeTwoList(leftList, rightList)
}
