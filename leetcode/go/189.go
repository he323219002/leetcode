package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
}

func rotate(nums []int, k int) {
	// 反转所有数组
	if k == 0 {
		return
	}
	realK := k % len(nums)
	rever(&nums, 0, len(nums)-1)
	// 反转 [0,realK]
	rever(&nums, 0, realK-1)
	rever(&nums, realK, len(nums)-1)
	fmt.Println(nums)

}

func rever(nums *[]int, start int, end int) {
	left, right := start, end
	for left < right {
		temp := (*nums)[left]
		(*nums)[left] = (*nums)[right]
		(*nums)[right] = temp

		left += 1
		right -= 1
	}
}

// func rotate(nums []int, k int) {
// 	fmt.Printf("切片的内存地址：%p\n", &nums)
// 	fmt.Println(nums)
// 	if k == 0 {
// 		return
// 	}

// 	// 用队列过渡
// 	queue := make([]int, 0)
// 	// 弹出次数
// 	popTime := k % len(nums)
// 	for i := 0; i < popTime; i++ {
// 		item := nums[len(nums)-1]
// 		queue = append([]int{item}, queue...)
// 		nums = nums[:len(nums)-1]
// 	}

// 	// 放到原数组首位
// 	for i := 0; i < popTime; i++ {
// 		item := queue[len(queue)-1]
// 		queue = queue[:len(queue)-1]
// 		nums = append([]int{item}, nums...)
// 	}

// 	fmt.Printf("切片的内存地址：%p\n", &nums)
// 	fmt.Println(nums)

// }
