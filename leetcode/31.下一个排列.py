from typing import *

class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        # 找到最靠右的升序的索引
        len_nums = len(nums)
        if len_nums == 1:
            return
        
        i_index = None
        j_index = None 
        temp_j = None
        for i in range(len_nums-1,0,-1):
            if nums[i] > nums[i-1]:
                i_index,j_index = i,i-1
                temp_j = nums[i-1]
                break
            else:
                continue
        
        # 全部降序则重新排序
        if not i_index and not j_index:
            nums.sort()
            return

        # 在j_index右边找到比temp_j大的值，大的越少越优先
        swap_val = nums[i_index]
        swap_index = i_index

        for i in range(i_index,len_nums):
            if nums[i] < swap_val and nums[i] > temp_j:
                swap_val = nums[i]
                swap_index = i
        
        # 交换位置
        nums[j_index],nums[swap_index] = nums[swap_index],nums[j_index]
        temp_li = sorted(nums[i_index:])
        nums[i_index:] = temp_li
        print(nums)

nums = [3,2,1]
# nums = [1,5,1]
# nums = [1,2,3,8,5,7,6,4]
Solution().nextPermutation(nums)


    
        
            