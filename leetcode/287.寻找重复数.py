from typing import *

class Solution:
	def findDuplicate(self, nums: List[int]) -> int:
		# 题设：nums.length == n + 1
		# 题设：1 <= nums[i] <= n
		# 转换为链表是否存在环，索引对应的值视为下一个节点的指针

		# 如何证明和环的思路一样
		# 索引对应的值永远不会指向数组之外，一直在数组内，所以必定是有环的链表
		slow = nums[0]
		fast = nums[nums[0]]
		if fast == slow:
			return slow

		while slow != fast:
			slow = nums[slow]
			fast = nums[nums[fast]]

		# 快慢指针相等时，从起点开始设置一个慢指针，相遇则为环的入口，即重复数字
		start = nums[0]
		slow = nums[slow]
		while start != slow:
			start = nums[start]
			slow = nums[slow]

		return slow
	
nums = [1,1]
# nums = [1,3,4,2,2]
# nums = [3,1,3,4,2]
res = Solution().findDuplicate(nums)
print(res)