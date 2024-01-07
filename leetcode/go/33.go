package main

func main() {

}

func search(nums []int, target int) int {
	if nums[0] == target {
		return 0
	}
	index := 0
	if len(nums) == 1 {
		return -1
	}

	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] < 0 {
			index = i - 1
			break
		}
	}

	var res int
	if target <= nums[index] && target >= nums[0] {

		res = binarySearch(&nums, target, 0, index)
	} else {
		res = binarySearch(&nums, target, index+1, len(nums)-1)
	}

	return res

}

func binarySearch(nums *[]int, target int, start int, end int) int {
	for {
		if start > end {
			break
		}
		middle := (start + end) / 2
		if (*nums)[middle] == target {
			return middle
		} else if (*nums)[middle] < target {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return -1
}
