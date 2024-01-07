from typing import * 
import collections

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        # 简历一个hash，数字转为str作为key，出现次数作为值
        dic = dict()
        for num in nums:
            count = dic.get(str(num))
            if not count:
                dic[str(num)] = 1
            else:
                dic[str(num)] = count + 1
        temp_res = collections.OrderedDict(sorted(dic.items(),key=lambda x:x[1]))
        res = []
        for i in range(k):
            tup = temp_res.popitem()
            res.append(int(tup[0]))
        
        return res
    

nums = [1,1,1,2,2,3]
k = 2
Solution().topKFrequent(nums,k)