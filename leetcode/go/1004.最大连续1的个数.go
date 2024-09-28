package main

import "fmt"

func main() {
	// nums, k := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2
	nums, k := []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3
	res := longestOnes(nums, k)
	fmt.Println(res)

}

func longestOnes(nums []int, k int) int {
	res, zeros, left := 0, 0, 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeros += 1
		}
		for zeros > k {
			if nums[left] == 0 {
				zeros -= 1
			}
			left += 1
		}
		res = max(res, right-left+1)
	}
	return res
}
