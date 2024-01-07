from typing import *

class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        # 用入度表的方法再写一次
        # dfs 本次自身遍历的点记为1，非自身遍历的点记为-1，未遍历0
        # 构造 {课程：遍历情况}
        search_map = {i:0 for i in range(numCourses)}
        # print(search_map)

        # 构造 {
        #   必修:[选修],
        # },表示选修和必修的关系
        course_map = {}
        for cur,pre in prerequisites:
            temp_cur = course_map.get(cur,[])
            temp_cur.append(pre)
            course_map[cur] = temp_cur
        # print(course_map)

        def dfs(course:int):
            # 无其他选修，则该节点置为-1，此次递归完成，退出
            temp_course = course_map.get(course)
            if not temp_course:
                search_map[course] = -1
                return False
        
        
            for per_course in temp_course:
                # 如果是其他节点遍历过的，说明无环，返回false
                if search_map.get(per_course) == -1:
                    continue
                # 如果是自身遍历过，说明有环，返回True
                elif search_map.get(per_course) == 1:
                    return True
                # 未遍历过，置为1,继续深度遍历
                else:
                    search_map[per_course] = 1
                    if dfs(per_course):
                        return True
                    search_map[per_course] = -1
                    continue
            return False
                 

        for per_course in range(numCourses):
            res = dfs(per_course)
            if res:
                return False
        return True
    



numCourses = 8
prerequisites =[[1,0],[2,6],[1,7],[6,4],[7,0],[0,5]]
# prerequisites =[[2,0],[1,0],[3,1],[3,2],[1,3]]
# prerequisites =[[1,4],[2,4],[3,1],[3,2]]
# prerequisites =[[1,0],[0,1]]
res = Solution().canFinish(numCourses,prerequisites)
print(res)

