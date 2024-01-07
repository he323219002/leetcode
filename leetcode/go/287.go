package main

import "fmt"

func main() {
	res := findDuplicate([]int{1, 3, 4, 2, 2})
	fmt.Println(res)
}

func findDuplicate(nums []int) int {
	// leetcode
	i := 0
	for i < len(nums) {
		if nums[i] == i {
			i += 1
			continue
		}

		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		swap(i, nums[i], &nums)
	}
	return -1
}

func swap(a int, b int, nums *[]int) {
	temp := (*nums)[a]
	(*nums)[a] = (*nums)[b]
	(*nums)[b] = temp
}
