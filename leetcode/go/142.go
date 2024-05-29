package main

import "fmt"

func main() {
	d := ListNode{-4, nil}
	c := ListNode{0, &d}
	b := ListNode{2, &c}
	a := ListNode{3, &b}
	d.Next = &b

	res := detectCycle(&a)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	// 路径放入字典
	nodeMap := make(map[*ListNode]bool)
	for node := head; node != nil; node = node.Next {
		if true == nodeMap[node] {
			return node
		}
		nodeMap[node] = true
	}
	return nil
}
