package main

func main() {

}

func moveZeroes(nums []int) {
	// leetcode 双指针
	a := 0
	b := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[b] = nums[i]
			b += 1
		}
		a += 1
	}
	for i := b; i < len(nums); i++ {
		nums[i] = 0
	}
}
