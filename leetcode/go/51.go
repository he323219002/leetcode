package main

import "strings"

func main() {
	n := 4
	solveNQueens(n)
}

func solveNQueens(n int) [][]string {
	// leetcode 答案提示
	// 初始化一个数组存坐标
	store := make([][]int, n)
	for i := 0; i < n; i++ {
		store[i] = make([]int, n)
		for j := 0; j < n; j++ {
			store[i][j] = 0
		}
	}

	res := make([][]string, 0)
	init := strings.Repeat(".", n)

	// dfs
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			temp := make([]string, 0)
			dfs(&res, n, i, j, &store, &temp, init)
		}
	}
	return res
}

func dfs(res *[][]string, n int, curRow int, curCol int, store *[][]int, temp *[]string, init string) {
	if curRow >= n || curCol >= n {
		return
	}
	if findInRow(store, curCol) || findInLeftLean(store, curRow, curCol) || findInRightLean(store, curRow, curCol) {
		return
	}

	(*store)[curRow][curCol] = 1
	prefix := init[:curCol]
	suffix := init[curCol+1:]
	tempRes := prefix + "Q" + suffix
	*temp = append(*temp, tempRes)
	if len(*temp) == n {
		*res = append(*res, *temp)
	}

	dfs(res, n, curRow+1, curCol, store, temp, init)
	dfs(res, n, curRow, curCol+1, store, temp, init)

}

// 不能在已存在的列上
func findInRow(store *[][]int, target int) bool {
	for _, perRows := range *store {
		if perRows[target] != 0 {
			return true
		}
	}
	return false
}

// 是否存在右斜线
func findInRightLean(store *[][]int, curRow int, curCol int) bool {
	// x - y 是个定值
	c := curRow - curCol
	for _, perRows := range *store {
		for perCol, perRow := range perRows {
			if (*store)[perRow][perCol] == 0 {
				continue
			}
			if perRow-perCol == c {
				return true
			}
		}
	}
	return false
}

// 是否存在左斜线
func findInLeftLean(store *[][]int, curRow int, curCol int) bool {
	// x + y 是个定值
	c := curRow + curCol
	for _, perRows := range *store {
		for perCol, perRow := range perRows {
			if (*store)[perRow][perCol] == 0 {
				continue
			}
			if perRow+perCol == c {
				return true
			}
		}
	}
	return false
}
