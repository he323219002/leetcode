class Solution:
    @classmethod
    def isMatch(self, s: str, p: str) -> bool:
        # 动态规划,dp[i][j] 表示s[..i]可以匹配p[..j]
        # 初始条件边界为 dp[0][0] 表示都为空字符串可以匹配
        # 状态转移 1.如果p[j]为字母，s[i]==p[j],dp[i][j]=dp[i-1][j-1]
        #         2.如果p[j]==*，dp[i][j]=dp[i][j-1]=dp[i][j-2],s[i]最后匹配0个，
        #           dp[i][j]=dp[i-1][j-1]=dp[i-1][j-2]=dp[i-1][j-3],s[i]最后匹配1个，依次类推
        #         3.如果p[j]=. 则可以匹配任何字母 dp[i][j]=dp[i][j-1]=dp[i-1][j-1]
        # dp = [[False] * len(p) for _ in range(len(s))]
        # dp[0][0] = True
        #
        m, n = len(s), len(p)

        def matches(i: int, j: int) -> bool:
            if i == 0:
                return False
            if p[j - 1] == '.':
                return True
            return s[i - 1] == p[j - 1]

        f = [[False] * (n + 1) for _ in range(m + 1)]
        f[0][0] = True
        for i in range(m + 1):
            for j in range(1, n + 1):
                if p[j - 1] == '*':
                    f[i][j] |= f[i][j - 2]
                    if matches(i, j - 1):
                        f[i][j] |= f[i - 1][j]
                else:
                    if matches(i, j):
                        f[i][j] |= f[i - 1][j - 1]
        print(f)
        return f[m][n]



print(Solution.isMatch('abc','ab.'))

