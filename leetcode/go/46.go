package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	permute(nums)
}

var result [][]int

// // temps 上一层传过来的，nums 过滤后的数组
// func dfs(temps []int, nums []int) {
// 	if len(nums) == 0 {
// 		result = append(result, temps)
// 		return
// 	}

// 	for index, v := range nums {
// 		res := make([]int, 0)
// 		res = append(res, temps...)
// 		res = append(res, v)
// 		mergedArr := make([]int, 0)
// 		start, end := nums[:index], nums[index+1:]
// 		mergedArr = append(mergedArr, start...)
// 		mergedArr = append(mergedArr, end...)
// 		dfs(res, mergedArr)
// 	}
// }

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	temp := make([]int, 0)
	if len(nums) == 1 {
		res = append(res, nums)
		return res
	}

	var result [][]int = make([][]int, 0)
	var dfs func(temps []int, nums []int)
	dfs = func(temps []int, nums []int) {
		if len(nums) == 0 {
			result = append(result, temps)
			return
		}

		for index, v := range nums {
			res := make([]int, 0)
			res = append(res, temps...)
			res = append(res, v)
			mergedArr := make([]int, 0)
			start, end := nums[:index], nums[index+1:]
			mergedArr = append(mergedArr, start...)
			mergedArr = append(mergedArr, end...)
			dfs(res, mergedArr)
		}
	}

	dfs(temp, nums)
	fmt.Println(result)
	return res
}
