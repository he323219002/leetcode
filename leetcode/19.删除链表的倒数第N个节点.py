# Definition for singly-linked list.

from typing import * 
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        dummy = ListNode(0,head)
        # 用栈操作
        stack = []
        while head:
            stack.append(head)
            head = head.next
        for i in range(n):
            n_node = stack.pop()
        if not stack:
            # 如果是第一个节点
            return n_node.next
        
        pre_n_node = stack.pop()
        pre_n_node.next = n_node.next
        return dummy.next
        
e = ListNode(5,None)
d = ListNode(4,e)
c = ListNode(3,d)
b = ListNode(2,c)
a = ListNode(1,b)

Solution().removeNthFromEnd(a,2)

            


            
        
