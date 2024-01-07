package main

func main() {

}

func twoSum(nums []int, target int) []int {
	// num为key 索引为val
	dict := make(map[int]int)
	for index, val := range nums {
		res, ok := dict[target-val]
		if ok {
			return []int{index, res}
		}
		dict[val] = index
	}
	return []int{0, 0}
}
