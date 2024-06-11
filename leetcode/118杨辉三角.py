from typing import *


class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        # 初始化数组
        li = [[0 * j for j in range(i + 1)] for i in range(numRows)]

        for row in range(numRows):
            li[row][0],li[row][-1] = 1,1
            if row == 0 or row == 1:
                continue
            for index in range(1,row):
                li[row][index] = li[row-1][index-1] + li[row-1][index]
            

        return li




Solution().generate(1)
