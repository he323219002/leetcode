from typing import *
class Solution:
    def minWindow(self, s: str, t: str) -> str:
        import collections
        # 双指针+哈希
        i,j = 0,0
        target_dict = collections.Counter(t)
        res = (0,len(s)-1)
        s_match = False

        temp_dic = {}

        while j < len(s):
            temp_s = s[i:j]
            temp_dic = collections.Counter(temp_s)
            # 当满足匹配 左指针 + 1
            while self.match(temp_dic,target_dict):
                # 记录
                if res[1] - res[0] > j - i:
                    res = (i,j)
                i += 1
                s_match = True
                break
            while not self.match(temp_dic,target_dict):
                j += 1
                break
            
        if s_match:
            print(res)
            # print(s[res[i]:res[j]])
            # return s[res[i]:res[j]+1]
            
    
    def match(self,temp:dict,t_dict:dict) -> bool:
        for s,count in t_dict.items():
            if temp.get(s,0) < count:
                return False
        return True


s = "ADOBECODEBANC"
t = "ABC"
Solution().minWindow(s,t)

