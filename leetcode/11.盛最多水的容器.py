from typing import *
class Solution:
    # def maxArea(self, height: List[int]) -> int:
    #     # d[i][j]标识 由i和j构成的最大值
    #     d = [[0] * len(height) for i in range(len(height))]
    #     max_area = 0
    #     for i in range(len(height)):
    #         for j in range(i,len(height)):
    #             if i == j:
    #                 continue
    #             else:
    #                 d[i][j] = min(height[i],height[j])*(j-i)
    #             max_area = max(max_area,d[i][j])

    #     return max_area
    def maxArea(self, height: List[int]) -> int:
        # 双指针
        left,right = 0,len(height) - 1
        max_area = 0
        # 面积由较短的决定，所以较短的指针位移
        while left < right:
            max_area = max(min(height[left],height[right])*(right-left),max_area)
            if height[left] < height[right]:
                left += 1
            else:
                right -= 1

        return max_area

height = [1,8,6,2,5,4,8,3,7]
res = Solution().maxArea(height)