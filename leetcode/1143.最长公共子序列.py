from typing import *

class Solution:
    def longestCommonSubsequence(self, text1: str, text2: str) -> int:
        # dp[i][j] 表示 text1[...i] 和 text2[...j]的最长公共子序列
        # dp[i][j] = dp[i-1][j-1] + 1  if text1[i] == text2[j]
        # else dp[i][j] = max[dp[i-1][j],dp[i][j-1]]
        