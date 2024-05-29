package main

import "fmt"

func main() {
	// nums := []int{5, 7, 7, 8, 8, 10}
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 3, 4, 4, 4, 4, 5, 5, 6, 6, 6, 8, 10, 10}
	res := searchRange(nums, 4)
	fmt.Println(res)
}

func searchRange(nums []int, target int) []int {
	start, end := -1, -1
	if len(nums) == 0 {
		return []int{start, end}
	}

	if nums[0] > target || nums[len(nums)-1] < target {
		return []int{start, end}
	}

	// 归并
	return bfs(&nums, 0, len(nums)-1, target)
}

func bfs(nums *[]int, left int, right int, target int) []int {
	if (*nums)[left] > target || (*nums)[right] < target {
		return []int{-1, -1}
	}

	if left == right {
		return []int{left, right}
	}

	mid := (left + right) / 2
	leftRes := bfs(nums, left, mid, target)
	rightRes := bfs(nums, mid+1, right, target)

	if leftRes[0] == -1 && rightRes[0] == -1 {
		return []int{-1, -1}
	}
	if leftRes[0] == -1 {
		return rightRes
	}
	if rightRes[0] == -1 {
		return leftRes
	}
	return []int{leftRes[0], rightRes[1]}

}
