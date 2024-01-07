package main

func main() {
	nums := []int{1, 3, 6, 5}
	nextPermutation(nums)
}

func nextPermutation(nums []int) {
	// leetcode
	if len(nums) <= 1 {
		return
	}

	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	// 找到赠序
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i >= 0 {
		// 为什么不找比nums[i]大最小的，比如nums[i] = 5，不是应该找6吗
		// 因为后面是降序排列，越往前越大
		for nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}

	// 双指针，降序变升序
	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
