package main

import "fmt"

func main() {
	// rooms := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	rooms := [][]int{{1}, {2}, {3}, {}}
	res := canVisitAllRooms(rooms)
	fmt.Println(res)
}

func canVisitAllRooms(rooms [][]int) bool {
	// dfs
	path := make(map[int]bool, len(rooms))
	firstRoom := rooms[0]

	// var pathPtr *map[int]bool = &path
	path[0] = true
	for i := 0; i < len(firstRoom); i++ {
		dfs(&path, firstRoom[i], &rooms)
	}

	if len(path) != len(rooms) {
		return false
	}
	return true
}

func dfs(path *map[int]bool, key int, rooms *[][]int) {
	// 如果已经进入过，则跳过
	if (*path)[key] {
		return
	}

	(*path)[key] = true
	nextRoomKeys := (*rooms)[key]
	for i := 0; i < len(nextRoomKeys); i++ {
		dfs(path, nextRoomKeys[i], rooms)
	}
}
