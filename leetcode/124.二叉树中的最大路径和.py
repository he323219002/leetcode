# Definition for a binary tree node.
from typing import *
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def maxPathSum(self, root: Optional[TreeNode]) -> int:
        """
        	// 递归
	// 每个节点的路径最大值 = max(左路径+val+右路径) 【左路径，右路径 > 0】
	// 左路径或右路径如何计算
	// max(node.left+val,val,node.right+val)
        """
        self.res = float('-inf')

        def searchNode(node:Optional[TreeNode]):
            if  node is None:
                return
            
            node.left_route,node.right_route = 0,0
            if node.left is None and node.right is None:
                self.res = max(self.res,node.val)
                return
            
            if node.left is not None:
                searchNode(node.left)
                node.left_route = max(node.left.left_route + node.left.val
                                    ,node.left.right_route + node.left.val
                                    ,node.left.val)
            if node.right is not None:
                searchNode(node.right)
                node.right_route = max(node.right.left_route + node.right.val
                                    ,node.right.right_route + node.right.val
                                    ,node.right.val)
          
            
            self.res = max(self.res,node.val + sum(filter(lambda x:x > 0,[node.left_route,node.right_route])))
        
        searchNode(root)
        return self.res


a = TreeNode(5)
b = TreeNode(4,a)
c = TreeNode(3,b)
d = TreeNode(2,c)
e = TreeNode(1,d)

Solution().maxPathSum(e)