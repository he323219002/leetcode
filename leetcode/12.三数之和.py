def threeSum(self, nums: list[int]) -> list[list[int]]:
    if len(nums) < 3:
        return []

    # 排序新数组
    nums.sort()
    len_nums = len(nums)

    # 遍历，双指针
    res = set()
    for i in range(len_nums):
        left = i + 1
        right = len_nums - 1
        two_sum = 0 - nums[i]
        while left < right:
            if nums[left] + nums[right] < two_sum:
                left += 1
            elif nums[left] + nums[right] > two_sum:
                right -= 1
            else:
                res.add((nums[i], nums[left], nums[right]))
                left += 1
    new_res = []
    for val in res:
        new_res.append(list(val))
    print(new_res)
    return new_res