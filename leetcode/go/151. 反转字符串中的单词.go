package main

import (
	"fmt"
	"strings"
)

func main() {
	// word := "the sky is blue"
	// word := "  hello world  "
	word := "a good   example"
	res := reverseWords(word)
	fmt.Println(res)

	// 已解答，答案用双指针法
}

func reverseWords(s string) string {
	// 队列和栈，队列存储每个单词，栈来存储队列
	stack := make([][]rune, 0)
	// 去除多余空格
	lastSpace := true

	queue := make([]rune, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && lastSpace {
			// 上一个也是空格
			continue
		} else if s[i] == ' ' {
			// 上一组放入栈中，且新建队列
			stack = append(stack, queue)
			queue = make([]rune, 0)
			lastSpace = true
		} else if s[i] != ' ' {
			queue = append(queue, rune(s[i]))
			lastSpace = false
			if i == len(s)-1 {
				stack = append(stack, queue)
			}
		}

	}

	newWordList := make([]string, 0)
	for len(stack) != 0 {
		var lastWord []rune
		lastWord = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		newWord := ""
		for len(lastWord) != 0 {
			newWordStr := lastWord[len(lastWord)-1]
			lastWord = lastWord[:len(lastWord)-1]
			newWord = string(newWordStr) + newWord
		}
		newWordList = append(newWordList, newWord)
	}

	res := strings.Join(newWordList, " ")
	return res
}
