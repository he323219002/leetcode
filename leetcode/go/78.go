package main

func main() {
	nums := []int{1, 2, 3}
	subsets(nums)
}

func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	tempPath := make([]int, 0)
	dfs(&res, tempPath, &nums, 0)
	return res
}

func dfs(res *[][]int, tempPath []int, nums *[]int, deepth int) {
	for i := deepth; i < len(*nums); i++ {
		tempPath = append(tempPath, (*nums)[i])
		newTempPath := make([]int, len(tempPath))
		copy(newTempPath, tempPath)
		(*res) = append((*res), newTempPath)
		dfs(res, tempPath, nums, i+1)
		tempPath = tempPath[:len(tempPath)-1]
	}
}
