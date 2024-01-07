from typing import *

class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        # leetcode 
        # max_res是最终结果，imax和imin是连续数组得到的结果，不会跳过中间元素运算
        max_res = float('-inf')
        imax = 1
        imin = 1
        for i in range(len(nums)):
            if nums[i] < 0:
                temp = imax
                imax = imin
                imin = temp
            imax = max(imax*nums[i],nums[i])
            imin = min(imin*nums[i],nums[i])
            max_res = max(max_res,imax)

        return max_res


nums = [2,3,-2,0,4]
Solution().maxProduct(nums)
