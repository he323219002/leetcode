from typing import *


class Solution:
    def canJump(self, nums: List[int]) -> bool:
        if len(nums) == 1:
            return True

        # 每次跳跃以第一步步长为基准，获取一个列表，
        # 处理列表，减去下标得到相对起点的距离
        # 取列表中最大步幅作为起点，重复上一步
        start = 0
        while start < len(nums):
            end = start + nums[start]
            if end >= len(nums) - 1:
                return True
            temp_li = []
            for i,v in enumerate(nums[start : end + 1]):
                temp_li.append(v+i)
  
            max_index = self.get_first_max_index(temp_li)
            # 不会前进
            if max_index == 0:
                return False
            start += max_index
        return True
            
    def get_first_max_index(self,lst):
        max_value = max(lst)
        first_max_index = lst.index(max_value)
        return first_max_index


nums = [1,1,1,0]
res = Solution().canJump(nums)
print(res)
