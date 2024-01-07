class Solution:
    @classmethod
    def convert(self, s: str, numRows: int) -> str:
        # 只有1行或2行直接返回
        if len(s) == 1 or len(s) == 2 or numRows == 1:
            return s

        # row有几行就有几个数组
        res = ["" for _ in range(numRows)]
        # flag 为遍历添加字母的顺序，group为数组
        flag = 1
        group = 0
        for i in s:
            res[group] += i
            # revert时机
            if group >= numRows - 1:
                flag = -1
            elif group == 0:
                flag = 1
            group += flag
        return ''.join(res)

print(Solution.convert('PAYPALISHIRING',3))
# PAHNAPLSIIGYIR