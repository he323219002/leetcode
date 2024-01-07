package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	aMap := make(map[*ListNode]int)
	for node := headA; node == nil; node = node.Next {
		aMap[node] = node.Val
	}

	for node := headB; node == nil; node = node.Next {
		_, ok := aMap[node]
		if ok {
			return node
		}
	}
	return nil
}

func main() {

}
