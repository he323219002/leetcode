package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// leetcode 双指针
	// 各自走完各自路线互换，如果相遇，则为相交节点
	for pA, pB := headA, headB; pA != nil && pB != nil; pA, pB = pA.Next, pB.Next {
		if pA.Next == nil {
			pA = headB
		}
		if pB.Next == nil {
			pB = headA
		}
		if pA == pB {
			return pA
		}
	}
	return nil
}

// func getIntersectionNode(headA, headB *ListNode) *ListNode {
// 	aMap := make(map[*ListNode]int)
// 	for node := headA; node == nil; node = node.Next {
// 		aMap[node] = node.Val
// 	}

// 	for node := headB; node == nil; node = node.Next {
// 		_, ok := aMap[node]
// 		if ok {
// 			return node
// 		}
// 	}
// 	return nil
// }

func main() {
	a := &ListNode{5, nil}
	b := &ListNode{4, a}
	c := &ListNode{8, b}
	d := &ListNode{1, c}
	e := &ListNode{4, d}

	h := &ListNode{1, c}
	i := &ListNode{6, h}
	j := &ListNode{5, i}

	getIntersectionNode(e, j)
}
