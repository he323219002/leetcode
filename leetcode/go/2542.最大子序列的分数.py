from typing import *
from heapq import heapify, heapreplace


class Solution(object):
    def maxScore(self, nums1, nums2, k):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :type k: int
        :rtype: int
        """
        # leetcode
        # 用zip将数组结合，按nums2倒序排列
        # 用容量为k的堆去装nums1，因为要求最小值的最大可能，遍历过程中 num2是越来越小，对应的m1如果比堆内最小的还小
        # 则num1*num2更小，不是最大可能，
        repo = sorted(zip(nums1, nums2), key=lambda x: x[-1], reverse=True)
        h = [i for i, _ in repo[:k]]
        heapify(h)
        h_sum = sum(h)
        ans = h_sum * repo[k - 1][1]

        for x, y in repo[k:]:
            # if x > h[0]:
            #     h_sum += x - heapreplace(h,x)
            #     ans = max(ans, h_sum * y)
            temp = x - heapreplace(h, x)
            if temp > 0:
                h_sum = h_sum + temp
                ans = max(ans, h_sum * y)
        return ans


nums1 = [4, 2, 3, 1, 1]
nums2 = [7, 5, 10, 9, 6]
k = 1
res = Solution().maxScore(nums1, nums2, k)
print(res)
