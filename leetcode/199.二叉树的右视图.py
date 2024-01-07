from typing import *
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
        if not root:
            return []
        res_map = {}
        # dfs
        def dfs(node:TreeNode,deep:int):
            if not node:
                return
            res = res_map.get(deep)
            if  res == None:
                res_map[deep] = node.val
            # 向右递归
            dfs(node.right,deep+1)
            # 右边结束后,判断左节点是否存在
            left_node = node.left
            if not left_node:
                return
            dfs(node.left,deep+1)
            
        dfs(root,1)
        res = list(map(lambda k: res_map[k],sorted(res_map.keys())))

        return res
            
e = TreeNode(4)
d = TreeNode(3,right=e)
c = TreeNode(5)
b = TreeNode(2,right=c)
a = TreeNode(1,b,d)

Solution().rightSideView(a)