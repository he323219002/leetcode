package main

func main() {
	a := []int{4, 4, 2, 1, 2}
	singleNumber(a)
}

func singleNumber(nums []int) int {
	numMap := make(map[int]int)
	for index, num := range nums {
		// 如果不在就放入，在就删除
		if _, ok := numMap[num]; ok {
			delete(numMap, num)
		} else {
			numMap[num] = index
		}
	}

	for res := range numMap {
		return res
	}
	return 0
}
