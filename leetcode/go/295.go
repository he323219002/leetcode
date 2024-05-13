package main

import (
	"fmt"
	"sort"
)

func main() {
	// 用一个数组去接收所有add的数字
	finder := Constructor()
	finder.AddNum(1)
	finder.AddNum(2)
	finder.AddNum(3)
	res := finder.FindMedian()
	fmt.Println(res)
}

type MedianFinder struct {
	Arr []int
}

func Constructor() MedianFinder {
	initArr := make([]int, 0)
	return MedianFinder{initArr}
}

func (this *MedianFinder) AddNum(num int) {
	binarySearchInsert(&this.Arr, num)
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.Arr)%2 == 1 {
		return float64(this.Arr[len(this.Arr)/2])
	} else {
		pos := len(this.Arr) / 2
		return float64((this.Arr[pos] + this.Arr[pos-1])) / 2
	}
}

func binarySearchInsert(nums *[]int, target int) *[]int {
	index := sort.SearchInts(*nums, target)  // 使用二分查找找到插入位置
	*nums = append(*nums, 0)                 // 扩展切片
	copy((*nums)[index+1:], (*nums)[index:]) // 向后移动元素
	(*nums)[index] = target                  // 插入目标元素
	return nums
}
