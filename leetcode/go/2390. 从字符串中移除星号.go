package main

import "fmt"

func main() {
	s := "erase*****"
	// s := "leet**cod*e"
	res := removeStars(s)
	fmt.Println(res)
}

func removeStars(s string) string {
	// 入栈 碰到* 出栈
	stack := make([]rune, 0)
	for _, str := range s {
		if str != '*' {
			stack = append(stack, str)
			continue
		}
		if len(stack) == 0 {
			continue
		}
		stack = stack[:len(stack)-1]
	}
	return string(stack)
}
