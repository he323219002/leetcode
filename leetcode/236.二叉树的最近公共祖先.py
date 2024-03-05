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
        # 前序遍历 找到其中一个值，以那个节点为开始继续遍历，否则回溯上一个节点，直到都找到
        self.res_node = None
        self.p_res, self.q_res = False, False

        def pre_search(node: TreeNode) -> TreeNode:
            if not node:
                return None

            if node.val == p:
                self.p_res = True
            elif node.val == q:
                self.q_res = True

            if self.p_res and self.q_res:
                return node

            if node.left:
                temp_left_res = pre_search(node.left)
                if self.res_node != None:
                    return node
                if temp_left_res:
                    self.res_node = node
                    return node

            if node.right:
                temp_right_res = pre_search(node.right)
                if self.res_node:
                    return node
                if temp_right_res:
                    self.res_node = node
                    return node

        res = pre_search(root)
        print(res)


root = TreeNode(1)

# 创建左子树
root.left = TreeNode(2)
root.left.left = TreeNode(3)
root.left.right = TreeNode(4)

# 创建右子树
root.right = TreeNode(5)
root.right.left = TreeNode(6)
root.right.right = TreeNode(7)
Solution().lowestCommonAncestor(root, 6, 2)
