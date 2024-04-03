# Definition for a binary tree node.
from typing import *
import collections


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        if not root: return []
        res, queue = [], collections.deque()
        queue.append(root)
        while queue:
            tmp = []
            for _ in range(len(queue)):
                node = queue.popleft()
                tmp.append(node.val)
                if node.left: queue.append(node.left)
                if node.right: queue.append(node.right)
            res.append(tmp)
        return res


    # def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
    #     if root == None:
    #         return []
    #     # leetcode BFS
    #     res, queue = [], []
    #     queue.append(root)

    #     while len(queue) != 0:
    #         self.handle_node(queue, res)

    #     return res

    # def handle_node(self, queue: List[TreeNode], res: List):
    #     temp_li = []
    #     while len(queue) != 0:
    #         cur_node = queue.pop(0)
    #         temp_li.append(cur_node)


    #     filter_li = list(filter(lambda x:x is not None,temp_li))
    #     temp_li_val = [i.val for i in filter_li]
    #     if len(temp_li_val) != 0:
    #         res.append(temp_li_val)
    #     for per_node in filter_li:
    #         queue.append(per_node.left)
    #         queue.append(per_node.right)

a = TreeNode(9,None,None)
b = TreeNode(15,None,None)
c = TreeNode(7,None,None)
d = TreeNode(20,b,c)
e = TreeNode(3,a,d)
res = Solution().levelOrder(e)
print(res)