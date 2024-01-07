from typing import *

# 广度遍历用go
class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        # 遍历每个点
        row = len(grid)
        col = len(grid[0])
        # dfs
        def dfs(grid: List[List[str]],per_row:int,per_col:int):
            if grid[per_row][per_col] == '0':
                return
            # 陆地变成水
            grid[per_row][per_col] = '0'
            # 不越界的情况继续dfs
            if per_row - 1>=0:
                dfs(grid,per_row-1,per_col)
            if per_row + 1<= row-1:
                dfs(grid,per_row+1,per_col)
            if per_col - 1 >= 0:
                dfs(grid,per_row,per_col-1)
            if per_col + 1 <= col-1:
                dfs(grid,per_row,per_col+1)
            
        count = 0
        for i in range(row):
            for j in range(col):
                # 0 跳过
                if  grid[i][j] == '0': continue
                # dfs 最外面首次调用dfs完成后 count+1 陆地变成水
                dfs(grid,i,j)
                count += 1
        return count

grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
Solution().numIslands(grid)
