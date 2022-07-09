class Solution:
    @classmethod
    def longestPalindrome(self, s: str) -> str:
        # 动态规划, s[i...j]为回文串 , 如果s[i-1] = s[j+1] 则为回文串
        # 边界条件 1个数字为回文串，2个数字相同为回文串
        s_len = len(s)
        if s_len == 1:
            return s
        if s_len == 2:
            return s if s[0] == s[1] else s[0]
        # dp[i][j] 表示s[i...j]是否为回文串
        dp = [[False] * s_len for _ in range(s_len)]
        # 每一个字符必定为回文串 即dp[1][1],dp[2][2]
        for i in range(s_len):
            dp[i][i] = True
        # 遍历顺序应该遵循窗口滑动，窗口递增才能满足状态转移
        max_len = 0
        begin = 0
        for l in range(1, s_len):
            for i in range(s_len):
                j = i + l
                if j >= s_len:
                    break

                if i + 1 == j:
                    if s[i] == s[j]:
                        dp[i][j] = True
                        continue

                if s[i] == s[j]:
                    dp[i][j] = dp[i + 1][j - 1]

                if dp[i][j] and j - i > max_len:
                    max_len = j - i + 1
                    begin = i

        return s[begin:begin + max_len]


res = Solution.longestPalindrome("cbbd")
print(res)
