from typing import *

class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0

        # leetcode
        nums_map = {val:index for index,val in enumerate(nums)}
        # 遍历nums，如果当前num-1在map中则跳过
        # 如果则判断num+1 是否有 如果无更新长度
        max_len = 1
        for num in nums_map.keys():
            if nums_map.get(num-1) != None:
                continue
            temp_len = 1
            temp_num = num + 1
            # 向上找
            while nums_map.get(temp_num) != None:
                temp_len += 1
                max_len = max(max_len,temp_len)
                temp_num += 1
        return max_len


# nums = [0,1,2,4,8,5,6,7,9,3,55,88,77,99,999999999]
nums = [100,4,200,1,3,2]
# nums = [0,3,7,2,5,8,4,6,0,1]
Solution().longestConsecutive(nums)

        
