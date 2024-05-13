package main

import "fmt"

func main() {
	res := longestPalindrome("aaaaa")
	fmt.Println(res)
}

func longestPalindrome(s string) string {
	// dp[i][j]表示是否是回文字符串
	// 如果dp[i][j]==true && s[i-1]=s[j+1] 则dp[i-1][j+1] = true

	rows := make([][]bool, 0)

	for i := 0; i < len(s); i++ {
		cols := make([]bool, len(s))
		rows = append(rows, cols)
	}

	res := string(s[0])
	for i := 0; i < len(rows); i++ {
		// 单个字符必是回文
		rows[i][i] = true
		// 两个连续字符如果相同也是回文
		if i > 0 && s[i-1] == s[i] {
			rows[i-1][i] = true
			res = s[i-1 : i+1]
		}

	}

	// 动态规划
	count := 0
	for i, row := range rows {
		for j, val := range row {
			if val {
				tempI, tempJ := i, j
				for tempI > 0 && tempJ < len(s)-1 {
					if s[tempI-1] == s[tempJ+1] {
						tempI -= 1
						tempJ += 1
						rows[tempI][tempJ] = true
						if tempJ-tempI > count {
							res = s[tempI : tempJ+1]
							count = tempJ - tempI
						}
					} else {
						break
					}
				}
			}
		}
	}

	return res
}
