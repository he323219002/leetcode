def maxArea(self, height: list[int]) -> int:
    if len(height) < 2:
        return 0
    # 定义左右指针
    left, right = 0, len(height) - 1
    max_size = 0
    while left < right:
        wide = right - left
        l = height[left] - height[right]
        max_size = max(max_size, wide * min(height[left], height[right]))
        if l <= 0:
            left += 1
        else:
            right -= 1

    return max_size
