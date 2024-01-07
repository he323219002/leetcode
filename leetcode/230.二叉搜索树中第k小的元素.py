from typing import *

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    # leetcode答案见go解法
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        li = []
        self.add_node(li,root)
        li.sort(key=lambda x:x.val)

    
        return li[k-1].val
    

    def add_node(self,li:List[TreeNode],node:TreeNode):
        li.append(node)
        if node.left:
            self.add_node(li,node.left)
        if node.right:
            self.add_node(li,node.right)

e = TreeNode(2,None,None)
d = TreeNode(1,None,e)
c = TreeNode(4,None,None)
b = TreeNode(3,d,c)
# a = TreeNode(2,None,None)

Solution().kthSmallest(b,2)
