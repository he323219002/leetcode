if __name__ == '__main__':
    class Solution:
        @classmethod
        def lengthOfLongestSubstring(self, s: str) -> int:
            if not s: return 0
            # 定义一个窗口，记录最大长度
            window = list()
            max_len = 0
            for val in s:
                while val in window:
                    window.pop(0)
                window.append(val)
                if len(window) >= max_len: max_len = len(window)
            return max_len

res = Solution.lengthOfLongestSubstring("bbbbbbbb")
print(res)