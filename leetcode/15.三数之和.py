from typing import *


class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        # 排序
        nums.sort()
        res = []
        # k指针从左到右
        for k in range(len(nums) - 2):
            # 重复则过滤
            if k > 0 and nums[k] == nums[k - 1]:
                continue
            if nums[k] > 0:
                return res
            i, j = k + 1, len(nums) - 1
            while i < j:
                if nums[k] + nums[i] + nums[j] == 0:
                    res.append([nums[k], nums[i], nums[j]])
                    i,j = i+1,j-1
                    while nums[i] == nums[i-1] and i < j:
                        i += 1
                    while nums[j] == nums[j+1] and i < j:
                        j -= 1
                    
                # 左指针右移
                elif nums[k] + nums[i] + nums[j] < 0:
                    step_i = 1
                    while i + step_i < j and nums[i] == nums[i + step_i]:
                        step_i += 1
                    i = i + step_i
                # 右指针左移
                elif nums[k] + nums[i] + nums[j] > 0:
                    step_j = 1
                    while j - step_j > i and nums[j] == nums[j - step_j]:
                        step_j += 1
                    j = j-step_j
        return res

nums = [-1,0,1,2,-1,-4]
Solution().threeSum(nums)
