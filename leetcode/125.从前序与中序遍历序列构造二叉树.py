from typing import *


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
        if not preorder:
            return None
        # pre 根、左、右
        # in 左、根、右
        if len(preorder) == 1:
            return TreeNode(preorder[0])

        # 根节点
        root = preorder[0]
        # 根的左边\右边子树节点个数
        root_p = inorder.index(root)
        left_tree_in_li = inorder[:root_p]
        right_tree_in_li = inorder[root_p + 1 :]

        left_tree_pre_li = preorder[1:1+len(left_tree_in_li)]
        right_tree_pre_li = preorder[1+len(left_tree_in_li):]

        left_tree = self.buildTree(left_tree_pre_li,left_tree_in_li)
        right_tree = self.buildTree(right_tree_pre_li,right_tree_in_li)

        tree = TreeNode(root,left_tree,right_tree)
        return tree


e = TreeNode(7,None,None)
d = TreeNode(15,None,None)
c = TreeNode(20,d,e)
b = TreeNode(9,None,None)
a = TreeNode(3,b,c)

# preorder = [3,9,20,15,7]
# inorder = [9,3,15,20,7]
preorder = [1,2]
inorder = [2,1]
res = Solution().buildTree(preorder,inorder)
print(res)
