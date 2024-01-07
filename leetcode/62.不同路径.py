from typing import *


class Solution:
    # def uniquePaths(self, m: int, n: int) -> int:
    #     # d[i][j] = d[i][j-1] + d[i-1][j]
    #     d = [[0] * n for _ in range(m)]
    #     # 沿着上和左的边都为1
    #     for i in range(m):
    #         for j in range(n):
    #             if i == 0:
    #                 d[i][j] = 1
    #                 continue
    #             if j == 0:
    #                 d[i][j] = 1
    #                 continue
    #             d[i][j] = d[i][j-1] + d[i-1][j]
    #     return d[m-1][j-1]

    # 只需要两行的值
    # def uniquePaths(self, m: int, n: int) -> int:
    #     pre = [1] * n
    #     cur = [1] * n
    #     for i in range(1, m):
    #         for j in range(1, n):
    #             cur[j] = pre[j] + cur[j-1]
    #         pre = cur[:]
    #     return pre[-1]

    # 左 + 上 前一个索引是左，当前索引是上一行，更新即可
    def uniquePaths(self, m: int, n: int) -> int:
        cur = [1] * n
        for i in range(1, m):
            for j in range(1, n):
                cur[j] += cur[j-1]
        return cur[-1]







Solution().uniquePaths(3,7)

