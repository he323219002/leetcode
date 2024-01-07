from typing import *

class Solution:
    def subarraySum(self, nums: List[int], k: int) -> int:
        # 前缀和 当前前缀和curr 和之前 curr-k 存在则+1
        # map 记录 curr-k 的个数
        res = 0
        res_map = {0:1}
        curr = 0
        for num in nums:
            curr += num
            res += res_map.get(curr-k,0)

            temp = res_map.get(curr,0)
            res_map[curr] = temp + 1

        return res
    
# nums = [1,2,3]
nums = [1,1,1]
Solution().subarraySum(nums,2)