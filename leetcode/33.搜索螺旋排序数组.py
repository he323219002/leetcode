from typing import *

class Solution:
	def search(self, nums: List[int], target: int) -> int:
		def binary_search(nums: List[int], target: int, start: int, end: int) -> int:
			while start <= end:
				middle = (start + end) // 2
				if nums[middle] == target:
					return middle
				elif nums[middle] < target:
					start = middle + 1
				else:
					end = middle - 1
			return -1

		index = 0
		# 找出旋转节点
		if len(nums) == 1:
			return 0 if nums[0] == target else -1
		for i in range(1, len(nums)):
			if nums[i] - nums[i - 1] < 0:
				index = i - 1
				break
		# 判断target在哪个区间
		if target <= nums[index] and target >= nums[0]:
			res = binary_search(nums, target, 0, index)
		else :
			res = binary_search(nums, target, index+1, len(nums) - 1)

		return res
	
nums = [3,1]
Solution().search(nums,1)