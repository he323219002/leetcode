package main

import (
	"strconv"
	"strings"
)

func main() {
	// leetcode
}

func decodeString(s string) string {
	stack := make([][2]interface{}, 0)
	res := ""
	multi := 0

	for _, c := range s {
		if c == '[' {
			stack = append(stack, [2]interface{}{multi, res})
			res, multi = "", 0
		} else if c == ']' {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			curMulti := top[0].(int)
			lastRes := top[1].(string)

			res = lastRes + strings.Repeat(res, curMulti)
		} else if c >= '0' && c <= '9' {
			num, _ := strconv.Atoi(string(c))
			multi = multi*10 + num
		} else {
			res += string(c)
		}
	}

	return res
}
