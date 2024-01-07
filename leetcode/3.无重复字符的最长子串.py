from typing import * 
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if len(s) == 0:
            return 0
        start,end = 0,len(s)
        win_len = 1
        max_len = 0
        while start + win_len <= end:
            temp_str = s[start:start+win_len]
            # 无重复，win_len + 1
            if len(list(temp_str)) == len(set(temp_str)):
                max_len = max(max_len,win_len)
                win_len += 1
                continue
            start += 1
        return max_len
        # print(max_len)


# s = "abcabcbb"
s = "pwwkew"
Solution().lengthOfLongestSubstring(s)
