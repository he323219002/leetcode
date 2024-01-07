from typing import *
class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        # 思考用分治方法
        # 动态规划 d[i] 表示num[0..i]最大值 （必须包含nums[i]）
        # if d[i-1] < 0  d[i] = num[i]
        # if d[i-1] >= 0 d[i] = d[i-1] + num[i]
        if len(nums) == 1:
            return nums[0]
        temp = nums[0]
        result = temp
        for i in range(1,len(nums)):
            if temp < 0:
                temp = nums[i]
            else:
                temp += nums[i]
            result = max(temp,result)
        return result



# nums = [-2,1,-3,4,-1,2,1,-5,4]
nums = [5,4,-1,7,8]
Solution().maxSubArray(nums)
            
