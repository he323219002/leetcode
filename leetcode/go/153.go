package main

import "fmt"

func main() {
	nums := []int{3, 1, 2}
	res := findMin(nums)
	fmt.Println(res)
}

// func findMin(nums []int) int {
// 	left, right := 0, len(nums)-1
// 	for left <= right {
// 		middle := left + (right-left)/2
// 		if nums[middle] < nums[right] {
// 			right = middle
// 		} else if nums[middle] > nums[right] {
// 			left = middle + 1
// 		} else {
// 			right -= 1
// 		}
// 	}
// 	return nums[left]
// }

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] < nums[right] {
			right = middle
		} else if nums[middle] > nums[right] {
			left = middle + 1
		} else {
			return nums[middle]
		}
	}
	return nums[left]
}
