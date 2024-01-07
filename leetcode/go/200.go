package main

// gpt
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	count := 0

	// 定义方向数组，用于上、下、左、右四个方向的移动
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// 定义递归函数进行广度优先搜索
	var bfs func(row, col int)
	bfs = func(row, col int) {
		// 检查边界条件和当前格子的值是否为 '1'
		if row < 0 || row >= rows || col < 0 || col >= cols || grid[row][col] != '1' {
			return
		}

		// 将当前格子标记为已访问
		grid[row][col] = '0'

		// 递归搜索上、下、左、右四个方向的相邻格子
		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]
			bfs(newRow, newCol)
		}
	}

	// 遍历整个网格
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 当遇到陆地 '1' 时，进行广度优先搜索
			if grid[i][j] == '1' {
				count++
				bfs(i, j)
			}
		}
	}

	return count
}

// type Queue [][]int

// func (q *Queue) Enqueue(item []int) {
// 	*q = append(*q, item)
// }

// func (q *Queue) Dequeue() []int {
// 	if len(*q) == 0 {
// 		return nil
// 	}
// 	item := (*q)[0]
// 	*q = (*q)[1:]
// 	return item
// }

// func bfs(grid *[][]byte, queue *Queue, row int, col int) {
// 	item := queue.Dequeue()
// 	if item == nil {
// 		return
// 	}
// 	perRow, perCol := item[0], item[1]

// 	// 陆地变水
// 	(*grid)[perRow][perCol] = '0'
// 	// 上下左右不越界则添加到队列
// 	if perRow-1 >= 0 && (*grid)[perRow-1][perCol] == '1' {
// 		(*grid)[perRow-1][perCol] = '0'
// 		queue.Enqueue([]int{perRow - 1, perCol})
// 	}
// 	if perRow+1 <= row-1 && (*grid)[perRow+1][perCol] == '1' {
// 		(*grid)[perRow+1][perCol] = '0'
// 		queue.Enqueue([]int{perRow + 1, perCol})
// 	}
// 	if perCol-1 >= 0 && (*grid)[perRow][perCol-1] == '1' {
// 		(*grid)[perRow][perCol-1] = '1'
// 		queue.Enqueue([]int{perRow, perCol - 1})
// 	}
// 	if perCol+1 <= col-1 && (*grid)[perRow][perCol+1] == '1' {
// 		(*grid)[perRow][perCol+1] = '0'
// 		queue.Enqueue([]int{perRow, perCol + 1})
// 	}
// 	bfs(grid, queue, row, col)
// }

// func numIslands(grid [][]byte) int {
// 	row := len(grid)
// 	col := len(grid[0])
// 	count := 0
// 	// 遍历每个点
// 	for i := 0; i < row; i++ {
// 		for j := 0; j < col; j++ {
// 			if grid[i][j] == '0' {
// 				continue
// 			}
// 			// bfs
// 			p := Queue{}
// 			p = append(p, []int{i, j})
// 			bfs(&grid, &p, row, col)
// 			count += 1
// 		}
// 	}
// 	return count
// }

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	numIslands(grid)
}
