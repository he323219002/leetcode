from typing import *
class Solution:
    def trap(self, height: List[int]) -> int:
        if len(height) == 0 or len(height) == 1:
            return 0
        # 双指针 低的动，可以获取左右各自走过的最大值 最后相加
        left,right = 0,len(height) - 1
        temp_left_height = height[left]
        temp_right_height = height[right]
        result = 0
        while right > left:
            if height[left] >= height[right]:
                right -= 1
                if right == left:
                    return result
                if height[right] < temp_right_height:
                    result += temp_right_height - height[right]
                else:
                    temp_right_height = height[right]
            else:
                left += 1
                if right == left:
                    return result
                if height[left] < temp_left_height:
                    result += temp_left_height - height[left]
                else:
                    temp_left_height = height[left]
        return result

height = [4,2,0,3,2,5]
res = Solution().trap(height)
print(res)

                
