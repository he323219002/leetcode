from typing import *


class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        # 合并成一个数组
        total_li = [num for row in matrix for num in row]
        # 用二分法判断是否在数组中
        start,end = 0,len(total_li)-1
        
        while start <= end:
            mid = (start + end) // 2
            if total_li[mid] == target:
                return True
            elif total_li[mid] < target:
                start = mid + 1
            elif total_li[mid] > target:
                end = mid - 1
        
        return False
            



# matrix = [[1, 3, 5, 7], [10, 11, 16, 20], [23, 30, 34, 60]]
# matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]]
# matrix = [[1]]
    
target = 3
res = Solution().searchMatrix(matrix, target)
print(res)