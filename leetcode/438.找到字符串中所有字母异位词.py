from typing import *


class Solution:
    def findAnagrams(self, s: str, p: str) -> List[int]:
        # 滑动窗口
        import collections

        def match(coll: collections.Counter):
            for k, v in coll.items():
                if p_dic.get(k) == None:
                    return False
                if p_dic.get(k) != v:
                    return False
            return True

        p_dic = collections.Counter(p)

        # 窗口大小
        win_len = len(p)

        res = []
        for i in range(len(s) - win_len + 1):
            temp_dic = collections.Counter(s[i:i+win_len])
            if match(temp_dic):
                res.append(i)

        print(res)
        return res


s = "cbaebabacd"
p = "abc"
Solution().findAnagrams(s, p)
