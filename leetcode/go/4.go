package main

import "fmt"

func main() {
	// nums1, nums2 := []int{1, 3}, []int{2, 7}
	// nums1, nums2 := []int{1, 3}, []int{2}
	// nums1, nums2 := []int{1, 2}, []int{3, 4}
	nums1, nums2 := []int{4, 5, 6, 8, 9}, []int{}
	res := findMedianSortedArrays(nums1, nums2)
	fmt.Println(res)

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 双指针合并数组
	resList := make([]int, 0)
	if len(nums1) == 0 {
		resList = nums2
	} else if len(nums2) == 0 {
		resList = nums1
	} else {
		nums1Pointer, nums2Pointer := 0, 0
		index := 0
		for nums1Pointer != len(nums1) && nums2Pointer != len(nums2) {
			if nums1[nums1Pointer] > nums2[nums2Pointer] {
				resList = append(resList, nums2[nums2Pointer])

				nums2Pointer += 1
			} else {
				resList = append(resList, nums1[nums1Pointer])
				// resList[index] = nums1[nums1Pointer]
				nums1Pointer += 1
			}
			index += 1
		}

		if nums1Pointer == len(nums1) {
			resList = append(resList, nums2[nums2Pointer:]...)
		}

		if nums2Pointer == len(nums2) {
			resList = append(resList, nums1[nums1Pointer:]...)
		}

	}

	// 如果是奇数个，返回中间
	if len(resList)%2 == 1 {
		return float64(resList[len(resList)/2])
	} else {
		temp := len(resList) / 2
		return float64(resList[temp]+resList[temp-1]) / 2
	}

}
