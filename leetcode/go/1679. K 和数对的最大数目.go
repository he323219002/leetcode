package main

import "fmt"

func main() {
	res := maxOperations([]int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2}, 3)
	fmt.Println(res)

	// 已解答 答案可用双指针 排序时间复杂度 nlogn 遍历n
}

func maxOperations(nums []int, k int) int {
	// 新建一个map，key是数字，val是出现的次数
	numsMap := make(map[int]int, 0)
	for _, val := range nums {
		count := numsMap[val]
		numsMap[val] = count + 1
	}

	res := 0
	for num, _ := range numsMap {
		target := k - num
		targetCount := numsMap[target]
		if targetCount == 0 {
			continue
		}
		// target 和 num 相同
		if target == num {
			for numsMap[num] >= 2 {
				numsMap[num] -= 2
				res += 1
			}
			continue
		}
		for numsMap[num] > 0 && numsMap[target] > 0 {
			numsMap[num] -= 1
			numsMap[target] -= 1
			res += 1
		}
	}
	return res
}
