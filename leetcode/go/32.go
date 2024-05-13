package main

import "fmt"

func main() {
	s := "))()())"
	res := longestValidParentheses(s)
	fmt.Println(res)
}

func longestValidParentheses(s string) int {
	// leetcode
	// 有效括号 dp[i] 表示 到s[i]最长有效括号的长度
	// if s[i]='(' dp[i]不表
	// if s[i]=')': if s[i-1]='(' dp[i]=dp[i-2]+2
	// 				if s[i-1]=')' && if s[i-dp[i-1]-1]='(' dp[i] = dp[i-1]+2+ dp[i-dp[i-1]-2]

	// 初始化
	if len(s) < 2 {
		return 0
	}
	dp := make([]int, len(s))
	maxLen := 0

	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			dp[i] = 0
			continue
		}
		if i < 2 && s[i-1] == '(' {
			dp[i] = 2
		} else if s[i-1] == '(' {
			dp[i] = dp[i-2] + 2
		} else if i-dp[i-1]-2 >= 0 && s[i-dp[i-1]-1] == '(' {
			dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
		} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
			dp[i] = dp[i-1] + 2
		}
		maxLen = max(maxLen, dp[i])
	}
	return maxLen
}
