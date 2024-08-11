package main

import "fmt"

func main() {
	maze := [][]byte{
		{'+', '+', '.', '+'},
		{'.', '.', '.', '+'},
		{'+', '+', '+', '.'},
	}
	entrance := []int{1, 2}
	res := nearestExit(maze, entrance)
	fmt.Println(res)
}

func nearestExit(maze [][]byte, entrance []int) int {
	// bfs 走过的地方变成墙
	path := make([][]int, 0)
	path = append(path, entrance)
	for len(path) != 0 {
		bfs(maze, entrance, 0, &path)
	}
	return bfs(maze, entrance, 0, &path)
}

func bfs(maze [][]byte, current []int, step int, path *[][]int) int {
	if current[0] == 0 || current[1] == 0 {
		if step != 0 {
			return step
		}
	}

	// 当前变成墙，四个方向分别放入队列
	maze[current[0]][current[1]] = '+'

	if current[0] > 0 && maze[current[0]-1][current[1]] == '.' {
		left := []int{current[0] - 1, current[1]}
		(*path) = append(*path, left)
	}

	if current[0] < len(maze[0])-1 && maze[current[0]+1][current[1]] == '.' {
		right := []int{current[0] + 1, current[1]}
		(*path) = append((*path), right)
	}

	if current[1] > 0 && maze[current[0]][current[1]-1] == '.' {
		up := []int{current[0], current[1] - 1}
		(*path) = append((*path), up)
	}

	if current[1] < len(maze)-1 && (maze)[current[0]][current[1]+1] == '.' {
		down := []int{current[0], current[1] + 1}
		(*path) = append((*path), down)
	}

	if len((*path)) == 0 {
		return -1
	}

	for len((*path)) != 0 {
		current = (*path)[0]
		(*path) = (*path)[1:]
		return bfs(maze, current, step+1, path)
	}

	return -1
}
