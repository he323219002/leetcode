package main

import "fmt"

func main() {
	d := ListNode{-4, nil}
	c := ListNode{0, &d}
	b := ListNode{2, &c}
	a := ListNode{3, &b}
	d.Next = &b

	res := hasCycle(&a)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// 快慢指针
	f := head
	s := head
	for f != nil {
		fNext := f.Next
		if fNext == nil {
			return false
		}
		f = fNext.Next
		s = s.Next
		if f == s {
			return true
		}
	}
	return false
}
