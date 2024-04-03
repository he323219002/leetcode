from typing import *


class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        res = []
        self.dfs(n,n,"",res)
        print(res)
        return res

    def dfs(self, left: int, right: int, curStr: str, res: List[str]):
        # 括号用完
        if left == 0 and right == 0:
            res.append(curStr)
            return

        # 左括号剩余比右括号多
        if left > right:
            return

        if left > 0:
            self.dfs(left - 1, right, curStr + "(", res)

        if right > 0:
            self.dfs(left, right - 1, curStr + ")", res)


Solution().generateParenthesis(3)