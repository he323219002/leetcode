from typing import *

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode()
        cache = 0
        node = dummy
        while l1 and l2:
            temp_val = l1.val + l2.val + cache
            node_val = temp_val -10
            if node_val >= 0:
                cache = 1
                temp_node = ListNode(temp_val - 10)
            else:
                cache = 0
                temp_node = ListNode(temp_val)
            node.next = temp_node
            node = node.next
            l1,l2 = l1.next,l2.next
            if l1 and not l2:
                l2 = ListNode(0)
            if l2 and not l1:
                l1 = ListNode(0)
        if cache:
            node.next = ListNode(1)


        print(dummy)
        return dummy.next


a = ListNode(3)
b = ListNode(4,a)
c = ListNode(2,b)
d = ListNode(4)
e = ListNode(6,d)
f = ListNode(5,e)
Solution().addTwoNumbers(c,f)