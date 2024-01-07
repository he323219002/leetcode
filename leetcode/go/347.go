package main

import "sort"

func main() {

}

func topKFrequent(nums []int, k int) []int {
	// gpt给出的解法
	// 1.构造hash
	hashmap := make(map[int]int)
	for _, num := range nums {
		hashmap[num]++
	}

	// 2.构造队列，存储{num,freq}，按freq对队列进行排序
	pairs := make([][2]int, 0, len(hashmap))
	for num, freq := range hashmap {
		pairs = append(pairs, [2]int{num, freq})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] > pairs[j][1]
	})

	// 3.构造结果返回
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = pairs[i][0]
	}

	return result

}
