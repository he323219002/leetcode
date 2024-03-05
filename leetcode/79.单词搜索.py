from typing import *
import collections
class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        # leetcode
        def dfs(row,col,index):
            if board[row][col] != word[index]:
                return False
            if index == len(word)-1:
                return True

            temp = board[row][col]
            # 防止路径重复
            board[row][col] = ''
            # 上
            if row > 0:
                if dfs(row-1,col,index+1):
                    return True
            # 右
            if col <len(board[0])-1:
                if dfs(row,col+1,index+1):
                    return True
                
            # 下 
            if row < len(board) -1:
                if dfs(row+1,col,index+1):
                    return True
            
            # 左
            if col > 0:
                if dfs(row,col-1,index+1):
                    return True
            
            # 还原
            board[row][col] = temp

            return False
        
        # dfs
        for row in range(len(board)):
            for col in range(len(board[0])):
                if dfs(row,col,0):
                    return True
        return False
                    

# board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]
board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]
# word = "ABCCED"
# word = "SEE"
word = "ABCB"
res = Solution().exist(board,word)
print(res)
