from typing import *

class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        # leetcode答案
        
        # 什么时候回溯终止：1.满足条件  2.超过target（此情况要将temp_res剔除）
        def traceback(temp_res:List[int],start:int,res:List[List[int]],target:int,can:List[int]):
            if target == 0:
                # list相当于深拷贝
                res.append(list(temp_res))
                return
      
            # 依次遍历节点加入临时结果集，继续向下
            for i in range(start,len(candidates)):
                if target - candidates[i] < 0:
                    break

                # 思考这么做的思路
                temp_res.append(candidates[i])
                traceback(temp_res,i,res,target-candidates[i],can)
                temp_res.pop()

        temp_res = []
        res = []
        # 为什么要排序：剪支需要，较大的数可能会由前面较小的组合而成，如果不排序遍历到超过target的数直接剪支会遗漏
        candidates.sort()
        traceback(temp_res,0,res,target,candidates)
        return res
    


candidates = [100,7,2,6,3]
target = 7
Solution().combinationSum(candidates,target)
            

