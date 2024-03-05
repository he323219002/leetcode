package main

func main() {
	nums := []int{1, 0}
	sortColors(nums)
}

func sortColors(nums []int) {
	// leetcoden
	p0, p1 := 0, 0
	for i, c := range nums {
		if c == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			// 如果p1 领先于 p0，则再交换
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p1 += 1
			p0 += 1
		} else if c == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1 += 1
		}
	}
}
