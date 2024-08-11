package main

func main() {
	a := ListNode{4, nil}
	// b := ListNode{3, &a}
	// c := ListNode{2, &b}
	// d := ListNode{1, &c}
	// e := ListNode{0, &d}

	deleteMiddle(&a)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	// 遍历2遍，第一遍得倒总数，第二遍操作
	origin := head
	count := 0
	for head != nil {
		count += 1
		head = head.Next
	}
	if count == 1 {
		return head
	}

	target := count / 2

	curNode := origin
	for i := 0; i < count; i++ {
		nextNode := curNode.Next
		if i == target-1 {
			if nextNode != nil {
				tempNode := nextNode.Next
				nextNode.Next = nil
				curNode.Next = tempNode
			}
			return origin
		}
		curNode = curNode.Next
	}
	return origin
}
