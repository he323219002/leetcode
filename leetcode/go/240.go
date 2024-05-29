package main

import "fmt"

func main() {

	matrix := [][]int{
		{-1, 3},
	}

	res := searchMatrix(matrix, 1)
	fmt.Println(res)
}

func searchMatrix(matrix [][]int, target int) bool {
	// 二分查找，每一次横竖查找后可以排除范围
	rows := len(matrix)
	cols := len(matrix[0])

	return recursion(&matrix, 0, cols-1, 0, rows-1, target)

}

func recursion(matrix *[][]int, left int, right int, up int, down int, target int) bool {

	if (*matrix)[up][left] > target || (*matrix)[down][right] < target {
		return false
	}

	if left == right && up == down {
		return (*matrix)[up][left] == target
	}

	initLeft := left
	// 横向查找
	nextLeft, nextRight, nextUp, nextDown := left, right, up, down

	for right >= left {
		mid := (left + right) / 2
		if (*matrix)[up][mid] == target {
			return true
		} else if (*matrix)[up][mid] > target {
			// 排除范围
			nextRight = mid
			right = mid - 1
		} else if (*matrix)[up][mid] < target {
			left = mid + 1
		}
	}
	// 横向没找到
	nextUp += 1
	if nextUp == len(*matrix) {
		return false
	}

	for down >= up {
		mid := (up + down) / 2
		if (*matrix)[mid][initLeft] == target {
			return true
		} else if (*matrix)[mid][initLeft] > target {
			// 排除范围
			nextDown = mid
			down = mid - 1
		} else if (*matrix)[mid][initLeft] < target {
			up = mid + 1
		}
	}
	// 纵向没找到
	nextLeft += 1
	if nextLeft == len((*matrix)[0]) {
		return false
	}

	return recursion(matrix, nextLeft, nextRight, nextUp, nextDown, target)

}
