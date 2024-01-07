from typing import *

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def sortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        # leetCode自上而下归并排序
        

        def sortFunc(head: ListNode, tail: ListNode) -> ListNode:
            if not head:
                return head
            if head.next == tail:
                head.next = None
                return head
            slow = fast = head
            while fast != tail:
                slow = slow.next
                fast = fast.next
                if fast != tail:
                    fast = fast.next
            mid = slow
            return mergeFunc(sortFunc(head, mid), sortFunc(mid, tail))



        
        def mergeFunc(node1:ListNode,node2:ListNode) -> ListNode:
            dummy = ListNode(0)
            temp,temp1,temp2 = dummy,node1,node2
            while temp1 and temp2:
                if temp1.val > temp2.val:
                    temp.next = temp2
                    temp2 = temp2.next
                else:
                    temp.next = temp1
                    temp1 = temp1.next
                temp = temp.next
            if temp1:
                temp.next = temp1
            elif temp2:
                temp.next = temp2
            
            return dummy.next
        
        res = sortFunc(head, None)
        return res



a = ListNode(0,None)
b = ListNode(4,a)
c = ListNode(3,b)
d = ListNode(5,c)
e = ListNode(-1,d)
Solution().sortList(e)