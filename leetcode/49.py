from typing import *

class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        str_dic = dict()
        for index,str in enumerate(strs):
            new_str = "".join(sorted(str))
            index_arr = str_dic.get(new_str,[])
            index_arr.append(index)
            str_dic[new_str] = index_arr


        res = []
        for k,v in str_dic.items():
            res.append([strs[i] for i in v])

        print(res)


# strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
strs = [""]
Solution().groupAnagrams(strs)

