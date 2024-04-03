from typing import *
class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        n = len(matrix)
        
        x_shaft = n // 2
        for i in range(x_shaft):
            for y in range(n):
                # 上下翻转
                matrix[i][y],matrix[n-i-1][y] = matrix[n-i-1][y],matrix[i][y]
        # 按对角线轴对称交换元素
        for i in range(n):
            for j in range(i+1, n):
                matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]

            
        print(matrix)

matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
# matrix = [[1,2,3],[4,5,6],[7,8,9]]
Solution().rotate(matrix)