package main

import (
	"fmt"
)

func main() {
	// nums := []int{1, 2, 3, 4, 5}
	nums := []int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2}
	res := increasingTriplet(nums)
	fmt.Println(res)
}

func increasingTriplet(nums []int) bool {
	// 遍历记录 一个small 和 一个mid
	// if a < samll ，update small
	// 	else a < mid , update mid
	// 即使更新了 small , 也隐含 存在一个n（更新前的small），使得 small < n < mid
	// 更新了mid 不影响结果，如果存在大于mid 则返回

	res := false
	if len(nums) < 3 {
		return res
	}

	small, mid, init := nums[0], nums[0], false
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] && !init {
			continue
		} else if !init {
			small, mid, init = nums[i-1], nums[i], true
		} else if nums[i] < small {
			small = nums[i]
		} else if nums[i] < mid && nums[i] != small {
			mid = nums[i]
		} else if nums[i] > mid {
			return true
		}
	}
	return res
}
