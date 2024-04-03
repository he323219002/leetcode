from typing import *


class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class Solution:
    def lowestCommonAncestor(
        self, root: "TreeNode", p: "TreeNode", q: "TreeNode"
    ) -> "TreeNode":
        # 递归遍历左右，根据返回值left，right
        # 1. left = none & right = none 表示无公共节点 (题设不存在)
        # 2. left != none & right != none 表示分布在此节点两侧 返回当前节点
        # 3. left != none & right = none 节点为left
        # 4. left = none & rigth != none 节点为right

        # 终止条件当前节点为none 无法继续遍历 2，找到p、 q 有对应返回

        if root == None:
            return root
        if root.val == p.val:
            return p
        if root.val == q.val:
            return q
        left = self.lowestCommonAncestor(root.left,p,q)
        right = self.lowestCommonAncestor(root.right,p,q)


        if left != None and right != None:
            return root
        if left != None and right == None:
            return left
        if left == None and right != None:
            return right
        
Solution().lowestCommonAncestor(None,None,None)