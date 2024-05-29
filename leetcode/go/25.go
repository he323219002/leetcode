package main

import "fmt"

func main() {
	a := ListNode{5, nil}
	b := ListNode{4, &a}
	c := ListNode{3, &b}
	d := ListNode{2, &c}
	e := ListNode{1, &d}

	res := reverseKGroup2(&e, 2)
	fmt.Println(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup2(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	//使用prevGroupTail记录上一个分组的尾部节点
	prevGroupTail := dummy

	for head != nil {
		// 检查剩余节点数量是否大于等于k
		curr := head
		count := 0
		for curr != nil && count < k {
			curr = curr.Next
			count++
		}

		// 剩余节点数量小于k，无需翻转，结束循环
		if count < k {
			break
		}

		// 进行当前分组的翻转
		groupTail := head
		prev := prevGroupTail
		for i := 0; i < k; i++ {
			next := head.Next
			head.Next = prev
			prev = head
			head = next
		}

		// 更新分组的前后连接关系
		prevGroupTail.Next = prev
		groupTail.Next = head
		prevGroupTail = groupTail
	}

	return dummy.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 使用双向队列
	// 如果队列满了，FILO，否则FIFO
	queue := make([]*ListNode, 0)

	dummy := ListNode{}
	curNode := &dummy
	for head != nil {
		curHead := head
		head = head.Next
		queue = append(queue, curHead)

		// 如果队列满了
		if len(queue) == k {
			for i := 0; i < k; i++ {
				lastNode := queue[len(queue)-1]
				lastNode.Next = nil
				queue = queue[:len(queue)-1]
				curNode.Next = lastNode
				curNode = curNode.Next
			}
			continue
		}

		// 如果队列未满
		if head == nil {
			curNode.Next = queue[0]
		}
	}
	return dummy.Next

}
