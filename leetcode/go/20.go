package main

import (
	"fmt"
)

func main() {
	// s := "(]"
	// s := "()[]{}"
	s := "({[()]})"
	// s := "([)]"
	res := isValid(s)
	fmt.Println(res)
}

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, bracket := range s {
		if len(stack) == 0 || bracket == '(' || bracket == '[' || bracket == '{' {
			stack = append(stack, bracket)
			continue
		}
		leftBracket := stack[len(stack)-1]
		if bracket == ')' && leftBracket != '(' {
			return false
		} else if bracket == ']' && leftBracket != '[' {
			return false
		} else if bracket == '}' && leftBracket != '{' {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
