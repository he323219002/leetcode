package main

import "fmt"

func main() {
	res := partition("cdd")
	fmt.Println(res)
}

func isLoopStr(s []byte, left int, right int) bool {
	// 判断是否是回文串
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left += 1
		right -= 1
	}
	return true
}

func dfs(s []byte, totalDeep int, curDeep int, path *[]string, res *[][]string) {
	if totalDeep == curDeep {
		*res = append(*res, append([]string{}, *path...))
		return
	}

	// 开始截取
	for i := 1; i < len(s)+1; i++ {
		if !isLoopStr(s, 0, i-1) {
			continue
		}

		*path = append(*path, string(s[:i]))
		dfs(s[i:], totalDeep, curDeep+i, path, res)
		// 回溯
		*path = (*path)[:len(*path)-1]
	}
}

func partition(s string) [][]string {
	// leetcode
	res := make([][]string, 0)
	// 路径容器
	path := make([]string, 0)

	dfs([]byte(s), len(s), 0, &path, &res)
	return res
}
