package main

func main() {
	// firstMissingPositive([]int{1, 2, 3})
	firstMissingPositive([]int{1})
}

func firstMissingPositive(nums []int) int {
	numsMap := make(map[int]int)
	for index, v := range nums {
		numsMap[v] = index
	}

	// [3,4,1,-1]
	for i := 1; i < len(nums)+2; i++ {
		_, ok := numsMap[i]
		if !ok {
			return i
		}
	}
	//[1,2,3]
	return len(nums)

}

// func firstMissingPositive(nums []int) int {
// 	// 第一遍遍历，找到最大的正整数
// 	var maxNum int
// 	for _, v := range nums {
// 		maxNum = max(v, maxNum)
// 	}
// 	// 初始化坑位
// 	tempLi := make([]int, maxNum)

// 	// 第二遍遍历 小于等于0的跳过
// 	for _, v := range nums {
// 		if v <= 0 {
// 			continue
// 		}
// 		tempLi[v-1] = 1
// 	}

// 	// 第三遍遍历
// 	for index, v := range tempLi {
// 		if v == 0 {
// 			return index + 1
// 		}
// 	}

// 	// fmt.Println(tempLi)
// 	return maxNum + 1
// }
