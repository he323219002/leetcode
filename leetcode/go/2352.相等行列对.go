package main

import "fmt"

func main() {
	grid := [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}
	res := equalPairs(grid)
	fmt.Println(res)
}

func equalPairs(grid [][]int) int {
	// 构造key为行数组，val为次数的字典
	row, col := len(grid), len(grid[0])
	listMap := make(map[string]int)

	// 遍历行
	for i := 0; i < row; i++ {
		tempLi := make([]int, row)
		for j := 0; j < col; j++ {
			tempLi[j] = grid[i][j]
		}
		listMap[fmt.Sprintf("%v", tempLi)] += 1
	}

	// 遍历列
	res := 0
	for j := 0; j < col; j++ {
		tempLi := make([]int, row)
		for i := 0; i < row; i++ {
			tempLi[i] = grid[i][j]
		}
		keyString := fmt.Sprintf("%v", tempLi)
		res = res + listMap[keyString]
	}
	return res
}
