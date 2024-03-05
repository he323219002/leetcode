from typing import *


class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        # leetcode
        # dp[i][j] 表示 由word1 i位置到word2 j位置所需要编辑的步数
        # if word1[i] == word2[j] 则dp[i][j] = dp[i-1][j-1]
        # else dp[i][j] =
        # dp[0][0] = 0
        row, col = len(word1), len(word2)
        dp = [[0 for _ in range(col + 1)] for _ in range(row + 1)]

        # 第一行
        for per_col in range(1, col + 1):
            dp[0][per_col] = dp[0][per_col - 1] + 1

        # 第一列
        for per_row in range(1, row + 1):
            dp[per_row][0] = dp[per_row - 1][0] + 1

        for per_row in range(1,row + 1):
            for per_col in range(1,col + 1):
                if word1[per_row-1] == word2[per_col-1]:
                    dp[per_row][per_col] = dp[per_row - 1][per_col - 1]
                elif word1[per_row-1] != word2[per_col-1]:
                    dp[per_row][per_col] = (
                        min(
                            dp[per_row - 1][per_col - 1],
                            dp[per_row - 1][per_col],
                            dp[per_row][per_col - 1],
                        )
                        + 1
                    )

        return dp[-1][-1]


word1 = ""
word2 = ""
res = Solution().minDistance(word1, word2)
print(res)
