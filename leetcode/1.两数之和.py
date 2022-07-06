# 给定一个整数数组nums和一个整数目标值target，请你在该数组中找出和为目标值target的那两个整数，并返回它们的数组下标。
# 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
# 你可以按任意顺序返回答案。

if __name__ == '__main__':

    class Solution:
        @classmethod
        # 双指针算法
        def twoSum(self, nums: list[int], target: int) -> list[int]:
            i = 0
            while True:
                for j in range(len(nums) - 1, i, -1):
                    if nums[i] + nums[j] == target:
                        return [i, j]

                i += 1

        # 官方算法
        def twoSum2(self, nums: list[int], target: int) -> list[int]:
            n = len(nums)
            for i in range(n):
                for j in range(i + 1, n):
                    if nums[i] + nums[j] == target:
                        return [i, j]

            return []
