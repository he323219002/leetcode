# Definition for a binary tree node.
from typing import *


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def inorderTraversal(self, root: Optional[TreeNode]) -> List[int]:
        def track(li: List[int], root: Optional[TreeNode]):
            if not root:
                return
            track(li, root.left)
            li.append(root.val)
            track(li, root.right)

        li = list()
        track(li, root)
        return li


a = TreeNode(val=3)
b = TreeNode(val=2, left=a)
c = TreeNode(val=1, right=b)

d = TreeNode(val=1)

test = Solution()
test.inorderTraversal(d)
