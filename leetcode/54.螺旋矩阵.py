from typing import *

class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        # 此题没想到巧解
        if not matrix: return []
        l,r,b,t,res = 0,len(matrix[0])-1,len(matrix)-1,0,[]

        while 1:
            # 从左到右
            for i in range(l,r+1):
                res.append(matrix[t][i])
            # 到达右边界后，上边界+1，如果上边界和下边界重合则停止
            t += 1
            if t > b: break

            # 从上到下
            for i in range(t,b+1):
                res.append(matrix[i][r])
            # 到达下边界后，有边界-1，如果左右边界重合则停止
            r -= 1
            if l>r: break

            # 从右到左
            for i in range(r,l-1,-1):
                res.append(matrix[b][i])
            # 到达左边界后，下边界-1，如果上边界和下边界重合则停止
            b -= 1
            if t > b: break

            # 从下到上
            for i in range(b,t-1,-1):
                res.append(matrix[i][l])
            # 到达上边界后，左边界+1，如果左右边界重合则停止
            l += 1
            if l>r: break

        return res




matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Solution().spiralOrder(matrix)
