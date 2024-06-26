package main

import "fmt"

func main() {
	nums, target := []int{1, 3, 5, 6}, 8
	res := searchInsert(nums, target)
	fmt.Println(res)
}

func searchInsert(nums []int, target int) int {
	// 二分查找
	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return start
}
