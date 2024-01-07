from typing import *

class Solution:
    def numSquares(self, n: int) -> int:
        # leetcode思路
        # dp[n] 表示n可以由多少个完全平方数构成，dp[0] = 0,dp[i] = min(dp[i],dp[i-j^2] + 1)
        dp = [i for i in range(n+1)]
        for i in range(1,n+1):
            for j in range(1,i):
                temp = i - j*j
                if temp < 0:
                    break
                dp[i] = min(dp[i],dp[temp]+1)
        return dp[-1]
        # print(dp)

Solution().numSquares(14)

    
