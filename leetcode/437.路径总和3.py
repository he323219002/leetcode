from typing import *
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def pathSum(self, root: Optional[TreeNode], targetSum: int) -> int:
        # leetcode
        # 定义一个从根节点开始的方法，递归遍历每个节点
        def search_root(root:TreeNode,target:int)->int:
            if not root:
                return 0
            temp = 0
            if root.val == target:
                temp += 1

            temp += search_root(root.left,target-root.val)
            temp += search_root(root.right,target-root.val)
            return temp
        
        res = 0
        if not root:
            return res
        
        res += search_root(root,targetSum)
        res += self.pathSum(root.left,targetSum)
        res += self.pathSum(root.right,targetSum)
        return res
        