package main

func main() {

}

func majorityElement(nums []int) int {
	// 空间复杂度On，空间复杂度O1
	// 票数count 元素major

	var count int = 0
	var major int
	for i := 0; i < len(nums); i++ {
		if count == 0 {
			count = 1
			major = nums[i]
			continue
		}
		if major == nums[i] {
			count += 1
			continue
		}
		count -= 1

	}
	return major

}
