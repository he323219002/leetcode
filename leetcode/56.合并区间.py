from typing import *
class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        # 先排序
        sort_li = sorted(intervals,key=lambda x:x[0])
        merge_li = []

        for per_li in sort_li:
            # 没有元素直接加入
            if len(merge_li) == 0:
                merge_li.append(per_li)
                continue
            # 左边元素和最后一个右边比较
            per_li_left,per_li_right = per_li[0],per_li[-1]
            last_num = merge_li[-1][-1]

            if last_num < per_li_left:
                # 不能合并
                merge_li.append(per_li)
                continue
            
            if last_num < per_li_right:
                # 需要合并
                merge_li[-1][-1] = per_li_right
                continue
        
        return merge_li

intervals = [[1,3],[2,6],[8,10],[15,18]]
Solution().merge(intervals)
